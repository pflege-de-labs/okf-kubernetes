package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const maxDepth = 5 // nested include/glossary_definition expansion

var (
	reShortcode = regexp.MustCompile(`(?s)\{\{([<%])\s*(/?)\s*([A-Za-z][\w-]*)(.*?)(/?)\s*[>%]\}\}`)
	reArg       = regexp.MustCompile(`(?:([\w-]+)=)?(?:"([^"]*)"|(\S+))`)
	reComment   = regexp.MustCompile(`(?s)<!--.*?-->`)
	reBlank     = regexp.MustCompile(`\n{3,}`)
	reMDHref    = regexp.MustCompile(`\]\(\s*([^)\s]+)((?:\s+"[^"]*")?)\)`)
	reHTMLHref  = regexp.MustCompile(`(href|src)="([^"]+)"`)
	reRefDef    = regexp.MustCompile(`(?m)^(\s*\[[^\]]+\]:\s*)(\S+)`)
)

type args struct {
	named map[string]string
	pos   []string
}

func parseArgs(s string) args {
	a := args{named: map[string]string{}}
	for _, m := range reArg.FindAllStringSubmatch(s, -1) {
		v := m[2]
		if m[3] != "" {
			v = m[3]
		}
		if m[1] != "" {
			a.named[m[1]] = v
		} else {
			a.pos = append(a.pos, v)
		}
	}
	return a
}

type frame struct {
	name string
	args args
}

// convert renders Hugo shortcodes to plain Markdown and rewrites site links to bundle links.
func (g *gen) convert(p *page, body string, depth int) string {
	body = g.shortcodes(p, body, depth)
	if depth > 0 {
		return body
	}
	body = reComment.ReplaceAllString(body, "")
	body = strings.NewReplacer("{{</*", "{{<", "*/>}}", ">}}", "{{%/*", "{{%", "*/%}}", "%}}").Replace(body)
	body = g.rewriteLinks(p, body)
	return reBlank.ReplaceAllString(body, "\n\n")
}

func (g *gen) shortcodes(p *page, body string, depth int) string {
	var b strings.Builder
	var stack []frame
	skip := 0 // inside {{< comment >}}
	last := 0
	for _, m := range reShortcode.FindAllStringSubmatchIndex(body, -1) {
		if skip == 0 {
			b.WriteString(body[last:m[0]])
		}
		last = m[1]
		closing := body[m[4]:m[5]] == "/"
		name := body[m[6]:m[7]]
		a := parseArgs(body[m[8]:m[9]])
		selfClosing := body[m[10]:m[11]] == "/"

		if name == "comment" {
			if closing {
				skip--
			} else {
				skip++
			}
			continue
		}
		if skip > 0 {
			continue
		}
		if closing {
			var open frame
			for i := len(stack) - 1; i >= 0; i-- {
				if stack[i].name == name {
					open, stack = stack[i], stack[:i]
					break
				}
			}
			b.WriteString(g.closeTag(name, open))
			continue
		}
		if name == "highlight" && b.Len() > 0 && !strings.HasSuffix(b.String(), "\n") {
			a.named["_inline"] = "true" // used mid-sentence: render as inline code, not a fence
		}
		if paired[name] && !selfClosing {
			stack = append(stack, frame{name, a})
		}
		b.WriteString(g.openTag(p, name, a, selfClosing, depth))
	}
	if skip == 0 {
		b.WriteString(body[last:])
	}
	return b.String()
}

var paired = map[string]bool{
	"note": true, "caution": true, "warning": true, "tabs": true, "tab": true, "table": true,
	"highlight": true, "mermaid": true, "alert": true, "details": true, "example": true,
}

func (g *gen) openTag(p *page, name string, a args, selfClosing bool, depth int) string {
	switch name {
	case "note", "caution", "warning":
		return "**" + g.label(name) + "**"
	case "tabs":
		return ""
	case "tab":
		out := "**" + a.named["name"] + ":**\n\n"
		if inc := a.named["include"]; inc != "" {
			return out + g.includeFile(p, path.Join(path.Dir(p.src), inc), depth)
		}
		if lang := a.named["codelang"]; lang != "" && !selfClosing {
			return out + "```" + lang
		}
		return out
	case "table":
		if c := a.named["caption"]; c != "" {
			return "*" + c + "*\n"
		}
		return ""
	case "highlight":
		if a.named["_inline"] != "" {
			return "`"
		}
		return "```" + first(a.pos)
	case "mermaid":
		return "```mermaid"
	case "alert":
		if t := a.named["title"]; t != "" {
			return "**" + t + "**"
		}
		return ""
	case "details":
		return "**" + a.named["summary"] + "**"
	case "example":
		return "["
	case "heading":
		if h, ok := g.headings[first(a.pos)]; ok {
			return h
		}
		return first(a.pos)
	case "glossary_tooltip":
		return g.tooltip(a)
	case "glossary_definition":
		return g.definition(a, depth)
	case "feature-state":
		return g.featureState(a)
	case "feature-gate-table":
		return g.gateTable(p, a, depth)
	case "code_sample", "codenew", "code":
		return g.codeSample(a)
	case "include":
		name := first(a.pos)
		if b, err := os.ReadFile(filepath.Join(g.site, "content/en/includes", name)); err == nil {
			return g.expand(p, b, depth)
		}
		return g.includeFile(p, path.Join(path.Dir(p.src), name), depth)
	case "skew":
		return g.skew(a)
	case "param":
		return g.param(p, first(a.pos))
	case "version-check":
		if v := str(p.fm["min-kubernetes-server-version"]); v != "" {
			return fmt.Sprintf("Your Kubernetes server must be at or later than version %s. To check the version, enter `kubectl version`.", v)
		}
		return "To check the version, enter `kubectl version`."
	case "ref", "relref":
		return g.ref(p, first(a.pos))
	case "figure":
		alt := a.named["alt"]
		if alt == "" {
			alt = a.named["title"]
		}
		out := fmt.Sprintf("![%s](%s)", alt, a.named["src"])
		if c := a.named["caption"]; c != "" {
			out += "\n\n*" + c + "*"
		}
		return out
	case "api-reference":
		return g.apiReference(a)
	case "thirdparty-content":
		return "> Items on this page refer to third party products or projects that are not part of Kubernetes itself."
	case "feature-gate-list":
		return g.gateList(p, a, depth)
	case "link":
		return "[" + a.named["text"] + "](" + a.named["url"] + ")"
	case "page-api-reference":
		return g.kindLink(a.named["kind"])
	case "pageinfo", "legacy-repos-deprecation", "toc", "cve-feed":
		return ""
	}
	g.unknown[name]++
	return ""
}

func (g *gen) closeTag(name string, open frame) string {
	switch name {
	case "highlight":
		if open.args.named["_inline"] != "" {
			return "`"
		}
		return "```"
	case "mermaid":
		return "```"
	case "tab":
		if open.args.named["codelang"] != "" {
			return "```"
		}
	case "example":
		return "](https://github.com/kubernetes/website/blob/" + g.commit + "/content/en/examples/" + open.args.named["file"] + ")"
	}
	return ""
}

func (g *gen) label(name string) string {
	return map[string]string{"note": "Note:", "caution": "Caution:", "warning": "Warning:"}[name]
}

func first(s []string) string {
	if len(s) == 0 {
		return ""
	}
	return s[0]
}

func (g *gen) tooltip(a args) string {
	id := a.named["term_id"]
	text := a.named["text"]
	t := g.terms[strings.ToLower(id)]
	if t == nil {
		g.warnings = append(g.warnings, "unknown glossary term "+id)
		return text
	}
	if text == "" {
		text = g.title(t)
	}
	return "[" + text + "](/" + t.out + ")"
}

func (g *gen) definition(a args, depth int) string {
	t := g.terms[strings.ToLower(a.named["term_id"])]
	if t == nil || depth >= maxDepth {
		return ""
	}
	def := t.body
	if a.named["length"] != "all" {
		def, _, _ = strings.Cut(def, "<!--more-->")
	}
	def = strings.TrimSpace(strings.ReplaceAll(def, "<!--more-->", ""))
	if pre := a.named["prepend"]; pre != "" && def != "" {
		def = pre + " " + strings.ToLower(def[:1]) + def[1:]
	}
	return g.shortcodes(t, def, depth+1)
}

func (g *gen) expand(p *page, raw []byte, depth int) string {
	if depth >= maxDepth {
		return ""
	}
	_, body, err := splitFrontmatter(raw)
	if err != nil {
		g.warnings = append(g.warnings, p.src+": include: "+err.Error())
		return ""
	}
	return g.shortcodes(p, strings.TrimSpace(body), depth+1)
}

// includeFile reads a docs file relative to content/en/docs.
func (g *gen) includeFile(p *page, rel string, depth int) string {
	b, err := os.ReadFile(filepath.Join(g.docs, filepath.FromSlash(rel)))
	if err != nil {
		g.warnings = append(g.warnings, fmt.Sprintf("%s: missing include %s", p.src, rel))
		return ""
	}
	return g.expand(p, b, depth)
}

func (g *gen) codeSample(a args) string {
	file := a.named["file"]
	b, err := os.ReadFile(filepath.Join(g.site, "content/en/examples", filepath.FromSlash(file)))
	if err != nil {
		g.warnings = append(g.warnings, "missing code sample "+file)
		return "`" + file + "`"
	}
	lang := a.named["language"]
	if lang == "" {
		lang = strings.TrimPrefix(path.Ext(file), ".")
	}
	return fmt.Sprintf("[`%s`](https://github.com/kubernetes/website/blob/%s/content/en/examples/%s)\n\n```%s\n%s\n```",
		file, g.commit, file, lang, strings.TrimRight(string(b), "\n"))
}

func (g *gen) currentStage(gt gate) map[string]any {
	if len(gt.stages) == 0 {
		return nil
	}
	return gt.stages[len(gt.stages)-1]
}

func (g *gen) featureState(a args) string {
	if name := a.named["feature_gate_name"]; name != "" {
		for _, gt := range g.gates {
			if gt.name != name {
				continue
			}
			s := g.currentStage(gt)
			if s == nil {
				break
			}
			def := "disabled"
			if s["defaultValue"] == true {
				def = "enabled"
			}
			return fmt.Sprintf("**Feature state:** Kubernetes v%s [%s] (feature gate `%s`, %s by default)", str(s["fromVersion"]), str(s["stage"]), name, def)
		}
		g.warnings = append(g.warnings, "unknown feature gate "+name)
		return fmt.Sprintf("**Feature state:** feature gate `%s`", name)
	}
	return fmt.Sprintf("**Feature state:** Kubernetes %s [%s]", a.named["for_k8s_version"], a.named["state"])
}

// gateTable mirrors the site's feature-gate-table shortcode, with descriptions inlined.
func (g *gen) gateTable(p *page, a args, depth int) string {
	removed := a.named["show-removed"] != ""
	alphaBeta := strings.Contains(a.named["include"], "alpha") || strings.Contains(a.named["include"], "beta")
	var b strings.Builder
	b.WriteString("| Feature | Default | Stage | Since | Until |\n|---|---|---|---|---|\n")
	for _, gt := range sortedGates(g.gates) {
		if gt.removed != removed {
			continue
		}
		if !removed {
			graduated := false
			for _, s := range gt.stages {
				if st := str(s["stage"]); st == "stable" || st == "deprecated" {
					graduated = true
				}
			}
			if alphaBeta == graduated {
				continue
			}
		}
		for _, s := range gt.stages {
			fmt.Fprintf(&b, "| `%s` | `%v` | %s | %s | %s |\n", gt.name, s["defaultValue"], str(s["stage"]), str(s["fromVersion"]), str(s["toVersion"]))
		}
	}
	return b.String()
}

// gateList mirrors the site's feature-gate-list shortcode: one description per gate.
func (g *gen) gateList(p *page, a args, depth int) string {
	removed := a.named["show-removed"] != ""
	var b strings.Builder
	for _, gt := range sortedGates(g.gates) {
		if gt.removed != removed {
			continue
		}
		desc := oneLine(reComment.ReplaceAllString(g.shortcodes(p, gt.body, depth+1), ""))
		fmt.Fprintf(&b, "* **`%s`**: %s\n", gt.name, desc)
	}
	return b.String()
}

func (g *gen) kindLink(kind string) string {
	for _, q := range g.pages {
		if m, ok := q.fm["api_metadata"].(map[string]any); ok && str(m["kind"]) == kind {
			return "[" + kind + "](/" + q.out + ")"
		}
	}
	g.warnings = append(g.warnings, "unknown api kind "+kind)
	return kind
}

func sortedGates(gs []gate) []gate {
	out := append([]gate(nil), gs...)
	sort.Slice(out, func(i, j int) bool { return out[i].name < out[j].name })
	return out
}

func (g *gen) skew(a args) string {
	add := func(n int) string {
		sep := "."
		if len(a.pos) > 2 && a.pos[2] != "" {
			sep = a.pos[2]
		}
		return g.major + sep + strconv.Itoa(g.minor+n)
	}
	n := 0
	if len(a.pos) > 1 {
		n, _ = strconv.Atoi(a.pos[1])
	}
	switch first(a.pos) {
	case "latestVersion", "currentVersion":
		return g.version()
	case "nextMinorVersion":
		return add(1)
	case "prevMinorVersion":
		return add(-1)
	case "oldestMinorVersion":
		return add(-2)
	case "latestVersionAddMinor", "currentVersionAddMinor":
		return add(n)
	case "currentPatchVersion":
		return g.patch
	}
	g.unknown["skew "+first(a.pos)]++
	return ""
}

func (g *gen) param(p *page, name string) string {
	if v := str(p.fm[name]); v != "" {
		return v
	}
	switch name {
	case "version":
		return "v" + g.version()
	case "fullversion":
		return "v" + g.patch
	}
	g.unknown["param "+name]++
	return ""
}

// ref resolves a Hugo ref relative to the page's source directory.
func (g *gen) ref(p *page, target string) string {
	target, anchor, _ := strings.Cut(target, "#")
	if anchor != "" {
		anchor = "#" + anchor
	}
	var base string
	switch {
	case strings.HasPrefix(target, "/docs/"), strings.HasPrefix(target, "docs/"):
		base = strings.TrimPrefix(strings.TrimPrefix(target, "/"), "docs/")
	default:
		base = path.Join(path.Dir(p.src), target)
	}
	base = strings.TrimSuffix(base, ".md")
	for _, c := range []string{base + ".md", base + "/_index.md", base + "/index.md"} {
		for _, q := range g.pages {
			if q.src == c {
				return "/" + q.out + anchor
			}
		}
	}
	g.warnings = append(g.warnings, fmt.Sprintf("%s: unresolved ref %s", p.src, target))
	return "https://kubernetes.io/docs/" + base + "/" + anchor
}

func (g *gen) apiReference(a args) string {
	src := "reference/kubernetes-api/" + a.named["page"] + ".md"
	q := g.bySrc[src]
	text := a.named["text"]
	if text == "" {
		text = a.named["anchor"]
	}
	if q == nil {
		g.warnings = append(g.warnings, "unknown api-reference page "+a.named["page"])
		return text
	}
	if text == "" {
		if m, ok := q.fm["api_metadata"].(map[string]any); ok {
			text = str(m["kind"])
		}
	}
	link := "/" + q.out
	if an := a.named["anchor"]; an != "" {
		link += "#" + an
	}
	return "[" + text + "](" + link + ")"
}

func (g *gen) rewriteLinks(p *page, body string) string {
	body = reMDHref.ReplaceAllStringFunc(body, func(m string) string {
		s := reMDHref.FindStringSubmatch(m)
		return "](" + g.resolve(p, s[1]) + s[2] + ")"
	})
	body = reHTMLHref.ReplaceAllStringFunc(body, func(m string) string {
		s := reHTMLHref.FindStringSubmatch(m)
		return s[1] + `="` + g.resolve(p, s[2]) + `"`
	})
	return reRefDef.ReplaceAllStringFunc(body, func(m string) string {
		s := reRefDef.FindStringSubmatch(m)
		return s[1] + g.resolve(p, s[2])
	})
}

// resolve maps a site URL to a bundle path when the target page is in the bundle,
// otherwise to an absolute kubernetes.io URL.
func (g *gen) resolve(p *page, href string) string {
	const site = "https://kubernetes.io"
	if strings.HasPrefix(href, site+"/") {
		href = strings.TrimPrefix(href, site)
	}
	if href == "" || strings.HasPrefix(href, "#") || strings.Contains(href, "://") ||
		strings.HasPrefix(href, "mailto:") || strings.HasPrefix(href, "{{") {
		return href
	}
	target, anchor, hasAnchor := strings.Cut(href, "#")
	if hasAnchor {
		anchor = "#" + anchor
	}
	target, query, hasQuery := strings.Cut(target, "?")
	if hasQuery {
		query = "?" + query
	}
	if strings.HasSuffix(target, ".md") && !strings.HasPrefix(target, "/docs/") {
		return href // already a bundle link
	}
	if !strings.HasPrefix(target, "/") {
		trail := strings.HasSuffix(target, "/")
		target = path.Join(p.url, target)
		if trail {
			target += "/"
		}
	}
	key := normURL(target)
	if key == "/docs/reference/glossary/" && strings.HasPrefix(anchor, "#term-") {
		if t := g.terms[strings.TrimPrefix(anchor, "#term-")]; t != nil {
			return "/" + t.out
		}
	}
	if q := g.byURL[key]; q != nil {
		return "/" + q.out + anchor
	}
	return site + target + query + anchor
}

func normURL(u string) string {
	u = strings.ToLower(strings.TrimSuffix(u, "index.html"))
	if !strings.HasPrefix(u, "/") {
		u = "/" + u
	}
	if !strings.HasSuffix(u, "/") && path.Ext(u) == "" {
		u += "/"
	}
	return u
}
