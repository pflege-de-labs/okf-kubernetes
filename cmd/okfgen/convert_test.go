package main

import (
	"strings"
	"testing"
)

func testGen() (*gen, *page) {
	term := &page{src: "reference/glossary/pod.md", out: "reference/glossary/pod.md", term: true,
		fm: map[string]any{"title": "Pod"}, body: "The smallest unit.\n<!--more-->\nMore detail."}
	names := &page{src: "concepts/overview/working-with-objects/names.md", out: "concepts/overview/working-with-objects/names.md",
		url: "/docs/concepts/overview/working-with-objects/names/", fm: map[string]any{}}
	labels := &page{src: "concepts/overview/working-with-objects/labels.md", out: "concepts/overview/working-with-objects/labels.md",
		url: "/docs/concepts/overview/working-with-objects/labels/", fm: map[string]any{}}
	g := &gen{major: "1", minor: 37, patch: "1.37.1", commit: "abc", headings: map[string]string{"whatsnext": "What's next"},
		terms: map[string]*page{"pod": term}, unknown: map[string]int{},
		byURL: map[string]*page{names.url: names, labels.url: labels},
		bySrc: map[string]*page{names.out: names, labels.out: labels},
		pages: []*page{term, names, labels},
		gates: []gate{{name: "Foo", stages: []map[string]any{{"stage": "beta", "defaultValue": true, "fromVersion": "1.36"}}}},
	}
	return g, names
}

func TestConvert(t *testing.T) {
	g, p := testGen()
	in := `<!-- overview -->
Run {{< glossary_tooltip text="pods" term_id="pod" >}} and {{< glossary_tooltip term_id="pod" >}}.
{{< glossary_definition term_id="pod" length="short" >}}
{{< note >}}
Careful.
{{< /note >}}
{{< comment >}}hidden {{< note >}}x{{< /note >}}{{< /comment >}}
{{< feature-state feature_gate_name="Foo" >}}
See [labels](
/docs/concepts/overview/working-with-objects/labels/#syntax), [sibling](../labels/), [setup](/docs/setup/) and {{< ref "labels#x" >}}.
Version {{< skew nextMinorVersion >}} / {{< skew currentPatchVersion >}}; escaped {{</* note */>}}.
## {{% heading "whatsnext" %}}
### DNS {#dns}
{{< tabs name="t" >}}{{< tab name="Linux" codelang="bash" >}}
ls
{{< /tab >}}{{< /tabs >}}
Rate {{< highlight promql >}}sum(x){{< /highlight >}} now.
{{< mystery a="b" >}}`
	got := g.convert(p, in, 0)
	for _, want := range []string{
		"Run [pods](/reference/glossary/pod.md) and [Pod](/reference/glossary/pod.md).",
		"The smallest unit.",
		"**Note:**\nCareful.",
		"**Feature state:** Kubernetes v1.36 [beta] (feature gate `Foo`, enabled by default)",
		"[labels](/concepts/overview/working-with-objects/labels.md#syntax)",
		"[sibling](/concepts/overview/working-with-objects/labels.md)",
		"[setup](https://kubernetes.io/docs/setup/)",
		"/concepts/overview/working-with-objects/labels.md#x",
		"Version 1.38 / 1.37.1; escaped {{< note >}}.",
		"## What's next",
		"### DNS {#dns}\n",
		"**Linux:**\n\n```bash\nls\n```",
		"Rate `sum(x)` now.",
	} {
		if !strings.Contains(got, want) {
			t.Errorf("missing %q in:\n%s", want, got)
		}
	}
	for _, bad := range []string{"hidden", "More detail", "<!--", "{{< glossary", "mystery"} {
		if strings.Contains(got, bad) {
			t.Errorf("unexpected %q in:\n%s", bad, got)
		}
	}
	if g.unknown["mystery"] != 1 {
		t.Errorf("unknown shortcode not counted: %v", g.unknown)
	}
}

func TestDescriptionFallback(t *testing.T) {
	g, p := testGen()
	body := "**Feature state:** x\n\nEach [object](/a.md) has\na name. Second sentence.\n"
	if got := g.description(p, body); got != "Each object has a name." {
		t.Errorf("got %q", got)
	}
}
