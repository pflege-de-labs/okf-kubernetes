package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	releasesSection = "releases"
	lifecycleOut    = "releases/lifecycle.md"
	aboutOut        = "about.md"
)

// releasePages are the content/en/releases pages worth an agent's attention; the rest are contacts and downloads.
var releasePages = map[string]bool{
	"_index.md": true, "patch-releases.md": true, "version-skew-policy.md": true, "release.md": true, "notes.md": true,
}

// releaseDescriptions replaces descriptions the fallback gets wrong (release.md starts with an auto-generation notice).
var releaseDescriptions = map[string]string{
	"release.md": "How a Kubernetes minor release is run: development, code freeze and post-release phases, and how enhancements, issues and PRs target a release milestone.",
}

type patchRelease struct {
	Release            string `yaml:"release"`
	TargetDate         string `yaml:"targetDate"`
	CherryPickDeadline string `yaml:"cherryPickDeadline"`
}

// releaseData mirrors kubernetes/website data/releases/{schedule,eol}.yaml.
type releaseData struct {
	Schedules []struct {
		Release                  string         `yaml:"release"`
		ReleaseDate              string         `yaml:"releaseDate"`
		MaintenanceModeStartDate string         `yaml:"maintenanceModeStartDate"`
		EndOfLifeDate            string         `yaml:"endOfLifeDate"`
		Next                     patchRelease   `yaml:"next"`
		PreviousPatches          []patchRelease `yaml:"previousPatches"`
	} `yaml:"schedules"`
	Upcoming []patchRelease `yaml:"upcoming_releases"`
	Branches []struct {
		Release           string `yaml:"release"`
		FinalPatchRelease string `yaml:"finalPatchRelease"`
		EndOfLifeDate     string `yaml:"endOfLifeDate"`
		Note              string `yaml:"note"`
	} `yaml:"branches"`
}

func (g *gen) loadReleaseData() error {
	for _, f := range []string{"schedule.yaml", "eol.yaml"} {
		b, err := os.ReadFile(filepath.Join(g.site, "data/releases", f))
		if err != nil {
			return err
		}
		if err := yaml.Unmarshal(b, &g.releases); err != nil {
			return fmt.Errorf("data/releases/%s: %w", f, err)
		}
	}
	if len(g.releases.Schedules) == 0 {
		return fmt.Errorf("data/releases/schedule.yaml: no schedules")
	}
	return nil
}

// latestPatch is the newest released patch of minor v, e.g. 1.37.1.
func (g *gen) latestPatch(v string) string {
	for _, s := range g.releases.Schedules {
		if s.Release == v && len(s.PreviousPatches) > 0 {
			return s.PreviousPatches[0].Release
		}
	}
	return v + ".0"
}

func dash(s string) string {
	if s == "" {
		return "-"
	}
	return s
}

// lifecycleBody lists dates only: a baked-in status would go stale between regenerations.
func (g *gen) lifecycleBody() string {
	var b strings.Builder
	fmt.Fprintf(&b, `The Kubernetes project maintains release branches for the three most recent minor releases.
A new minor release ships roughly every four months. Each minor release gets about 14 months of
patch support: 12 months of standard support, then 2 months of maintenance mode, then end of life
([support period](/releases/patch-releases.md#support-period)).

The tables are a snapshot of kubernetes/website@%s as of %s. They do not state a status.
Derive the status of a release from today's date:

| Status | Condition | Fixes shipped |
|---|---|---|
| supported (standard support) | today < maintenance mode start | regular monthly patch releases |
| maintenance mode (security and critical fixes only) | maintenance mode start <= today < end of life | CVE fixes, dependency and base-image updates, critical core component fixes |
| end of life (unsupported) | today >= end of life | none; upgrade to a supported release |

A release listed under maintained releases whose end-of-life date has passed is end of life;
the upstream data is updated after the fact.

## Maintained releases

| Release | Released | Latest patch | Next patch (target date) | Maintenance mode start | End of life |
|---|---|---|---|---|---|
`, g.commit[:12], g.commitDate)
	for _, s := range g.releases.Schedules {
		next := "-"
		if s.Next.Release != "" {
			next = fmt.Sprintf("%s (%s)", s.Next.Release, s.Next.TargetDate)
		}
		fmt.Fprintf(&b, "| %s | %s | %s | %s | %s | %s |\n", s.Release, dash(s.ReleaseDate), g.latestPatch(s.Release),
			next, dash(s.MaintenanceModeStartDate), dash(s.EndOfLifeDate))
	}
	b.WriteString("\n## End-of-life releases\n\n| Release | Final patch | End of life | Note |\n|---|---|---|---|\n")
	for _, e := range g.releases.Branches {
		fmt.Fprintf(&b, "| %s | %s | %s | %s |\n", e.Release, e.FinalPatchRelease, e.EndOfLifeDate, oneLine(e.Note))
	}
	if len(g.releases.Upcoming) > 0 {
		b.WriteString("\n## Upcoming patch release dates\n\n| Target date | Cherry-pick deadline |\n|---|---|\n")
		for _, u := range g.releases.Upcoming {
			fmt.Fprintf(&b, "| %s | %s |\n", u.TargetDate, u.CherryPickDeadline)
		}
	}
	next := g.major + "." + fmt.Sprint(g.minor+1)
	fmt.Fprintf(&b, `
## Next minor release

At the snapshot date, v%[2]s was the newest release and v%[1]s was in development; neither v%[1]s
nor later releases are part of this bundle. Minor releases ship about every four months, so if
today is more than four months after the newest release date above, assume a newer minor exists.
The v%[1]s schedule is in
[kubernetes/sig-release](https://github.com/kubernetes/sig-release/tree/master/releases/release-%[1]s).
Features planned for an unreleased version can still slip.

## Feature availability by release

The pages of this bundle describe Kubernetes v%[2]s. To check whether a feature is available
in another release, use the per-stage history (stage, default, since, until) in
[Feature Gates](/reference/command-line-tools-reference/feature-gates.md) and
[removed feature gates](/reference/command-line-tools-reference/feature-gates-removed.md),
API removals in the [Deprecated API Migration Guide](/reference/using-api/deprecation-guide.md),
and supported component version combinations in the [Version Skew Policy](/releases/version-skew-policy.md).
`, next, g.version())
	return b.String()
}

func (g *gen) aboutBody() string {
	return fmt.Sprintf(`This bundle documents **Kubernetes v%[1]s** (latest patch release %[2]s per the upstream schedule data).
It was generated from [kubernetes/website@%[3]s](https://github.com/kubernetes/website/tree/%[4]s), as of %[5]s.

Every page describes the behavior of v%[1]s. Before applying it to a cluster, compare v%[1]s with
the cluster's version (`+"`kubectl version`"+`):

* Cluster older than v%[1]s: features may be missing, still alpha or beta, or disabled by default.
  Check the stage history in [Feature Gates](/reference/command-line-tools-reference/feature-gates.md)
  and suggest the upgrade that makes a needed feature available.
* Cluster newer than v%[1]s: features may have graduated, been removed, or changed. Check
  [removed feature gates](/reference/command-line-tools-reference/feature-gates-removed.md) and the
  [Deprecated API Migration Guide](/reference/using-api/deprecation-guide.md), and say that the bundle is older.
* Support status (supported, maintenance mode, end of life) and dates of every release:
  [Release Lifecycle](/releases/lifecycle.md).
* Supported version differences between components, and upgrade order:
  [Version Skew Policy](/releases/version-skew-policy.md).
* Curated naming rules and limits: [Requirements and limits](/requirements/index.md).
`, g.version(), g.patch, g.commit[:12], g.commit, g.commitDate)
}

// releasePagesGenerated adds the lifecycle and about concepts, which have no website page.
func (g *gen) releasePagesGenerated() {
	lifecycle := &page{
		root: "", src: "data/releases/schedule.yaml", out: lifecycleOut, url: "/releases/",
		fm: map[string]any{
			"title":        "Release Lifecycle",
			"description":  "Release, maintenance-mode and end-of-life dates of every Kubernetes minor release, and how to derive whether a release is supported, in maintenance mode (security fixes only), or end of life.",
			"content_type": "reference",
		},
		body:   g.lifecycleBody(),
		weight: -1,
		extra:  []source{{ID: "k8s-eol-data", Resource: g.sourceRaw + "data/releases/eol.yaml", Title: "Kubernetes website source: data/releases/eol.yaml"}},
	}
	about := &page{
		root: "", src: "hugo.toml", out: aboutOut, // no published page
		fm: map[string]any{
			"title":        "About this bundle: Kubernetes v" + g.version(),
			"description":  "This bundle documents Kubernetes v" + g.version() + " (kubernetes/website@" + g.commit[:12] + "); how to apply it to clusters running other versions.",
			"content_type": "reference",
		},
		body: g.aboutBody(),
	}
	for _, p := range []*page{lifecycle, about} {
		g.bySrc[p.out] = p
		g.pages = append(g.pages, p)
	}
}
