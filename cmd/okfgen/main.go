// Command okfgen converts the Kubernetes website documentation (concepts,
// tasks, reference) into an Open Knowledge Format (OKF) v0.2 bundle.
//
// See docs/generator.md for the mapping rules.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

var sections = []string{"concepts", "tasks", "reference"}

const (
	gatesDir    = "reference/command-line-tools-reference/feature-gates"
	glossaryDir = "reference/glossary"
	curatedDir  = "requirements"
	generatedBy = "process:okf-kubernetes-okfgen"
)

type page struct {
	src    string // relative to content/en/docs
	out    string // relative to the bundle root
	url    string // published site path, e.g. /docs/concepts/
	fm     map[string]any
	body   string
	weight int
	term   bool // glossary term
	desc   string
}

type gate struct {
	name    string
	stages  []map[string]any
	removed bool
	body    string
}

type gen struct {
	site      string // kubernetes/website checkout
	docs      string // <site>/content/en/docs
	out       string
	commit    string
	minor     int // minor of the latest release, e.g. 37 for v1.37
	major     string
	patch     string
	at        string
	pages     []*page
	byURL     map[string]*page
	bySrc     map[string]*page
	terms     map[string]*page
	gates     []gate
	headings  map[string]string
	unknown   map[string]int
	warnings  []string
	sourceRaw string // GitHub blob prefix
}

func main() {
	site := flag.String("site", "", "path to a kubernetes/website checkout")
	out := flag.String("out", "catalog", "bundle root to write")
	at := flag.String("at", time.Now().UTC().Format(time.RFC3339), "generated.at timestamp")
	flag.Parse()
	if *site == "" {
		log.Fatal("-site is required")
	}
	g := &gen{
		site: *site, docs: filepath.Join(*site, "content/en/docs"), out: *out, at: *at,
		byURL: map[string]*page{}, bySrc: map[string]*page{}, terms: map[string]*page{},
		unknown: map[string]int{},
	}
	if err := g.run(); err != nil {
		log.Fatal(err)
	}
}

func (g *gen) run() error {
	if err := g.loadSite(); err != nil {
		return err
	}
	for _, s := range sections {
		if err := g.load(s); err != nil {
			return err
		}
	}
	for _, s := range sections {
		if err := os.RemoveAll(filepath.Join(g.out, s)); err != nil {
			return err
		}
		_ = os.Remove(filepath.Join(g.out, s+".md"))
	}
	written := []*page{}
	for _, p := range g.pages {
		body := g.convert(p, p.body, 0)
		if strings.TrimSpace(body) == "" { // layout-only section pages, e.g. the glossary listing
			body = "See the [contents](" + path.Base(strings.TrimSuffix(p.out, ".md")) + "/index.md) of this section."
		}
		if err := g.write(p, body); err != nil {
			return err
		}
		written = append(written, p)
	}
	if err := g.writeIndexes(written); err != nil {
		return err
	}
	for name, n := range g.unknown {
		g.warnings = append(g.warnings, fmt.Sprintf("dropped unknown shortcode %q (%d)", name, n))
	}
	sort.Strings(g.warnings)
	for _, w := range g.warnings {
		log.Print("warning: ", w)
	}
	log.Printf("wrote %d concepts from kubernetes/website@%s (v%s.%d)", len(written), g.commit[:12], g.major, g.minor)
	return nil
}

var (
	reLatest  = regexp.MustCompile(`(?m)^latest\s*=\s*"v(\d+)\.(\d+)"`)
	reHeading = regexp.MustCompile(`(?m)^\[(\w+)_heading\]\s*\nother\s*=\s*"([^"]*)"`)
)

func (g *gen) loadSite() error {
	out, err := exec.Command("git", "-C", g.site, "rev-parse", "HEAD").Output()
	if err != nil {
		return fmt.Errorf("resolve source commit: %w", err)
	}
	g.commit = strings.TrimSpace(string(out))
	g.sourceRaw = "https://github.com/kubernetes/website/blob/" + g.commit + "/content/en/docs/"

	hugo, err := os.ReadFile(filepath.Join(g.site, "hugo.toml"))
	if err != nil {
		return err
	}
	m := reLatest.FindSubmatch(hugo)
	if m == nil {
		return fmt.Errorf("no latest version in hugo.toml")
	}
	g.major = string(m[1])
	g.minor, _ = strconv.Atoi(string(m[2]))

	i18n, err := os.ReadFile(filepath.Join(g.site, "i18n/en/en.toml"))
	if err != nil {
		return err
	}
	g.headings = map[string]string{}
	for _, h := range reHeading.FindAllSubmatch(i18n, -1) {
		g.headings[string(h[1])] = string(h[2])
	}

	// Patch version is only used by the skew shortcode; missing data is not fatal.
	var sched struct {
		Schedules []struct {
			Release         string
			PreviousPatches []struct{ Release string } `yaml:"previousPatches"`
		}
	}
	if b, err := os.ReadFile(filepath.Join(g.site, "data/releases/schedule.yaml")); err == nil && yaml.Unmarshal(b, &sched) == nil {
		for _, s := range sched.Schedules {
			if s.Release == g.version() && len(s.PreviousPatches) > 0 {
				g.patch = s.PreviousPatches[0].Release
			}
		}
	}
	if g.patch == "" {
		g.patch = g.version() + ".0"
	}
	return nil
}

func (g *gen) version() string { return fmt.Sprintf("%s.%d", g.major, g.minor) }

func (g *gen) load(section string) error {
	return filepath.WalkDir(filepath.Join(g.docs, section), func(p string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() || !strings.HasSuffix(p, ".md") {
			return err
		}
		rel, _ := filepath.Rel(g.docs, p)
		rel = filepath.ToSlash(rel)
		raw, err := os.ReadFile(p)
		if err != nil {
			return err
		}
		fm, body, err := splitFrontmatter(raw)
		if err != nil {
			g.warnings = append(g.warnings, fmt.Sprintf("%s: %v", rel, err))
			return nil
		}
		dir, base := path.Split(rel)
		dir = strings.TrimSuffix(dir, "/")
		if dir == gatesDir && base != "index.md" {
			g.gates = append(g.gates, gate{name: str(fm["title"]), stages: maps(fm["stages"]), removed: fm["removed"] == true, body: body})
			return nil
		}
		if fm["headless"] == true || neverRendered(fm) {
			return nil
		}
		pg := &page{src: rel, fm: fm, body: body, weight: toInt(fm["weight"])}
		switch {
		case base == "_index.md" || base == "index.md":
			pg.out = dir + ".md"
			pg.url = "/docs/" + dir + "/"
		default:
			pg.out = rel
			pg.url = "/docs/" + strings.TrimSuffix(rel, ".md") + "/"
		}
		if dir == glossaryDir && base != "index.md" {
			pg.term = true
			// Hugo resolves term_id against the file name, not the id key.
			g.terms[strings.ToLower(strings.TrimSuffix(base, ".md"))] = pg
		}
		if prev := g.bySrc[pg.out]; prev != nil {
			return fmt.Errorf("output collision: %s and %s both map to %s", prev.src, rel, pg.out)
		}
		g.bySrc[pg.out] = pg
		g.pages = append(g.pages, pg)
		g.byURL[strings.ToLower(pg.url)] = pg
		for _, a := range list(fm["aliases"]) {
			g.byURL[normURL(a)] = pg
		}
		return nil
	})
}

func neverRendered(fm map[string]any) bool {
	b, ok := fm["_build"].(map[string]any)
	if !ok {
		return false
	}
	r := fmt.Sprint(b["render"])
	return r == "false" || r == "never"
}

func splitFrontmatter(raw []byte) (map[string]any, string, error) {
	s := strings.ReplaceAll(string(raw), "\r\n", "\n")
	if !strings.HasPrefix(s, "---\n") {
		return map[string]any{}, s, nil
	}
	end := strings.Index(s[4:], "\n---")
	if end < 0 {
		return nil, "", fmt.Errorf("unterminated frontmatter")
	}
	fm := map[string]any{}
	if err := yaml.Unmarshal([]byte(s[4:4+end]), &fm); err != nil {
		return nil, "", fmt.Errorf("frontmatter: %w", err)
	}
	body := s[4+end+4:]
	body = strings.TrimPrefix(body, "\n")
	return fm, body, nil
}

type actor struct {
	By string `yaml:"by"`
	At string `yaml:"at"`
}

type source struct {
	ID       string `yaml:"id"`
	Resource string `yaml:"resource"`
	Title    string `yaml:"title"`
}

type frontmatter struct {
	Type              string         `yaml:"type"`
	Title             string         `yaml:"title,omitempty"`
	Description       string         `yaml:"description,omitempty"`
	Resource          string         `yaml:"resource,omitempty"`
	Tags              []string       `yaml:"tags,omitempty,flow"`
	Generated         actor          `yaml:"generated,flow"`
	Sources           []source       `yaml:"sources"`
	KubernetesVersion string         `yaml:"kubernetes_version"`
	SourceCommit      string         `yaml:"source_commit"`
	SourcePath        string         `yaml:"source_path"`
	MinServerVersion  string         `yaml:"min_kubernetes_server_version,omitempty"`
	APIMetadata       map[string]any `yaml:"api_metadata,omitempty"`
}

func (g *gen) write(p *page, body string) error {
	title := g.title(p)
	p.desc = g.description(p, body)
	published := "https://kubernetes.io" + p.url
	if p.term {
		published = "https://kubernetes.io/docs/reference/glossary/?all=true#term-" + termID(p)
	}
	fm := frontmatter{
		Type:              conceptType(p),
		Title:             title,
		Description:       p.desc,
		Resource:          published,
		Tags:              g.tags(p),
		Generated:         actor{By: generatedBy, At: g.at},
		KubernetesVersion: "v" + g.version(),
		SourceCommit:      g.commit,
		SourcePath:        "content/en/docs/" + p.src,
		MinServerVersion:  str(p.fm["min-kubernetes-server-version"]),
	}
	if m, ok := p.fm["api_metadata"].(map[string]any); ok {
		fm.APIMetadata = m
	}
	fm.Sources = []source{
		{ID: "k8s-docs-source", Resource: g.sourceRaw + p.src, Title: "Kubernetes website source: " + p.src},
		{ID: "k8s-docs", Resource: published, Title: title + " | Kubernetes documentation"},
	}
	var buf bytes.Buffer
	enc := yaml.NewEncoder(&buf)
	enc.SetIndent(2)
	if err := enc.Encode(fm); err != nil {
		return err
	}
	content := "---\n" + buf.String() + "---\n\n# " + title + "\n\n" + strings.TrimSpace(body) + "\n"
	dst := filepath.Join(g.out, filepath.FromSlash(p.out))
	if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
		return err
	}
	return os.WriteFile(dst, []byte(content), 0o644)
}

func termID(p *page) string {
	return strings.TrimSuffix(path.Base(p.src), ".md")
}

func (g *gen) title(p *page) string {
	if t := strings.TrimSpace(str(p.fm["title"])); t != "" {
		return t
	}
	return strings.TrimSuffix(path.Base(p.out), ".md")
}

func conceptType(p *page) string {
	if p.term {
		return "glossary-term"
	}
	switch strings.TrimSpace(str(p.fm["content_type"])) {
	case "concept":
		return "concept"
	case "task":
		return "task"
	case "tutorial":
		return "tutorial"
	case "api_reference":
		return "api-reference"
	case "tool-reference":
		return "tool-reference"
	case "reference":
		return "reference"
	}
	if strings.HasSuffix(p.src, "_index.md") {
		return "section"
	}
	return strings.SplitN(p.src, "/", 2)[0]
}

func (g *gen) tags(p *page) []string {
	parts := strings.Split(p.out, "/")
	tags := []string{"kubernetes", parts[0]}
	if len(parts) > 2 {
		tags = append(tags, parts[1])
	}
	if p.term {
		tags = append(tags, list(p.fm["tags"])...)
	}
	return tags
}

var (
	reMDLink  = regexp.MustCompile(`\[([^\]]*)\]\([^)]*\)`)
	reSpaces  = regexp.MustCompile(`\s+`)
	reSentEnd = regexp.MustCompile(`[.!?](\s|$)`)
)

// description prefers the curated frontmatter and falls back to the first sentence of the body.
func (g *gen) description(p *page, body string) string {
	for _, k := range []string{"description", "short_description"} {
		if d := oneLine(reMDLink.ReplaceAllString(g.shortcodes(p, str(p.fm[k]), 1), "$1")); d != "" {
			return d
		}
	}
	inFence := false
	lines := strings.Split(body, "\n")
	for i, line := range lines {
		t := strings.TrimSpace(line)
		if strings.HasPrefix(t, "```") {
			inFence = !inFence
			continue
		}
		if inFence || t == "" || strings.ContainsAny(t[:1], "#|<>*-!`[") || strings.HasPrefix(t, "1.") {
			continue
		}
		for _, next := range lines[i+1:] { // the sentence may wrap onto following lines
			if strings.TrimSpace(next) == "" || strings.HasPrefix(strings.TrimSpace(next), "```") {
				break
			}
			t += " " + strings.TrimSpace(next)
		}
		t = oneLine(reMDLink.ReplaceAllString(t, "$1"))
		if loc := reSentEnd.FindStringIndex(t); loc != nil {
			t = t[:loc[0]+1]
		}
		if len(t) > 300 {
			t = t[:297] + "..."
		}
		return t
	}
	return g.title(p)
}

func oneLine(s string) string { return strings.TrimSpace(reSpaces.ReplaceAllString(s, " ")) }

// writeIndexes writes a reserved index.md for every generated directory plus the bundle root.
func (g *gen) writeIndexes(pages []*page) error {
	children := map[string][]*page{}
	dirs := map[string]bool{}
	for _, p := range pages {
		d := path.Dir(p.out)
		children[d] = append(children[d], p)
		for ; d != "."; d = path.Dir(d) {
			dirs[d] = true
		}
	}
	for d := range dirs {
		var b strings.Builder
		title := path.Base(d)
		if sp := g.bySrc[d+".md"]; sp != nil {
			title = g.title(sp)
		}
		fmt.Fprintf(&b, "# %s\n\n", title)
		entries := children[d]
		sort.Slice(entries, func(i, j int) bool { return less(g, entries[i], entries[j]) })
		listed := map[string]bool{}
		for _, p := range entries {
			name := path.Base(p.out)
			fmt.Fprintf(&b, "* [%s](%s) - %s", g.title(p), name, p.desc)
			if sub := strings.TrimSuffix(p.out, ".md"); dirs[sub] {
				fmt.Fprintf(&b, " ([contents](%s/index.md))", path.Base(sub))
				listed[sub] = true
			}
			b.WriteString("\n")
		}
		subs := []string{}
		for s := range dirs {
			if path.Dir(s) == d && !listed[s] {
				subs = append(subs, s)
			}
		}
		sort.Strings(subs)
		for _, s := range subs {
			fmt.Fprintf(&b, "* [%s](%s/index.md)\n", path.Base(s), path.Base(s))
		}
		if err := os.WriteFile(filepath.Join(g.out, d, "index.md"), []byte(b.String()), 0o644); err != nil {
			return err
		}
	}
	return g.writeRootIndex()
}

func less(g *gen, a, b *page) bool {
	if a.weight != b.weight {
		return a.weight < b.weight
	}
	return g.title(a) < g.title(b)
}

func (g *gen) writeRootIndex() error {
	var b strings.Builder
	b.WriteString("---\nokf_version: \"0.2\"\n---\n\n# Kubernetes Knowledge Catalog\n\n")
	fmt.Fprintf(&b, "OKF bundle of the Kubernetes v%s documentation (kubernetes/website@%s).\n\n", g.version(), g.commit[:12])
	b.WriteString("# Start here\n\n")
	b.WriteString("* [Requirements and limits](requirements/index.md) - Curated, source-cited naming rules, size limits, reserved values, and ranges.\n\n")
	for _, s := range sections {
		p := g.bySrc[s+".md"]
		fmt.Fprintf(&b, "# %s\n\n* [%s](%s.md) - %s ([contents](%s/index.md))\n\n", g.title(p), g.title(p), s, p.desc, s)
	}
	return os.WriteFile(filepath.Join(g.out, "index.md"), []byte(b.String()), 0o644)
}

func str(v any) string {
	if v == nil {
		return ""
	}
	if s, ok := v.(string); ok {
		return s
	}
	return fmt.Sprint(v)
}

func toInt(v any) int {
	i, _ := v.(int)
	return i
}

func list(v any) []string {
	var out []string
	if l, ok := v.([]any); ok {
		for _, x := range l {
			if s := str(x); s != "" {
				out = append(out, s)
			}
		}
	}
	return out
}

func maps(v any) []map[string]any {
	var out []map[string]any
	if l, ok := v.([]any); ok {
		for _, x := range l {
			if m, ok := x.(map[string]any); ok {
				out = append(out, m)
			}
		}
	}
	return out
}
