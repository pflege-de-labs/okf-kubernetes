# OKF bundle: Kubernetes documentation

[Open Knowledge Format v0.2](https://github.com/GoogleCloudPlatform/knowledge-catalog/blob/main/okf/SPEC.md)
bundle of the Kubernetes documentation sections
[Concepts](https://kubernetes.io/docs/concepts/), [Tasks](https://kubernetes.io/docs/tasks/)
and [Reference](https://kubernetes.io/docs/reference/), pinned to one
`kubernetes/website` commit (currently Kubernetes v1.37, `2fb3b0d636e8`).

## Layout

| Path | Content | Maintained by |
|------|---------|---------------|
| `catalog/` | Bundle root (`okf_version: "0.2"` in `catalog/index.md`) | |
| `catalog/requirements/` | Curated constraints: name formats and lengths, label/annotation rules, namespaces, size limits, ports. Every claim is footnoted to a commit-pinned source line range | hand |
| `catalog/concepts/`, `catalog/tasks/`, `catalog/reference/` | One concept per docs page, mirroring the site structure | `cmd/okfgen` |
| `catalog/log.md` | Update history | hand |
| `.claude-plugin/`, `plugins/kubernetes-okf/` | Claude Code marketplace and skill that tells Claude when to use this bundle: [docs/claude-plugin.md](docs/claude-plugin.md) | hand |

Concept ids follow the page path, e.g. `kubernetes/concepts/overview/working-with-objects/names`.
Start with `kubernetes/requirements/object-names` for naming rules.

Glossary terms (`reference/glossary/*`) are the main relationship hubs: every
glossary tooltip on the site became a link to them.

## Claude Code

```sh
/plugin install kubernetes-okf --marketplace pflege-de-labs/okf-kubernetes
```

The skill makes Claude use this bundle through the pflege.de OKF MCP server; details in [docs/claude-plugin.md](docs/claude-plugin.md).

## Regenerate

```sh
git clone --depth 1 --filter=blob:none --sparse https://github.com/kubernetes/website.git /tmp/website
git -C /tmp/website sparse-checkout set content/en/docs/concepts content/en/docs/tasks \
  content/en/docs/reference content/en/examples content/en/includes i18n/en data/releases
go run ./cmd/okfgen -site /tmp/website -out catalog -at "$(date -u +%Y-%m-%dT%H:%M:%SZ)"
```

`-at` sets `generated.at` in every document (default: now).

The generator replaces `catalog/{concepts,tasks,reference}` and the generated
`index.md` files; `catalog/requirements/` and `catalog/log.md` are left alone.
When the source commit changes, re-check the line ranges cited in
`catalog/requirements/` and add a `log.md` entry. Mapping rules:
[docs/generator.md](docs/generator.md).

## Validate

```sh
go test ./...
uvx --from okf-ingest==0.13.1 okf validate catalog/
```

`okf validate` reports same-page `#anchor` links as `broken_link` warnings; these are expected.

To check it as the pflege.de MCP server loads it (from a checkout of `pflege-de/mcp`):

```sh
go run ./cmd/okf-mcp serve --transport http --packs okf --bundle kubernetes=<this repo>/catalog
```

Expect `loaded OKF catalog bundles=1` with no warnings.

## License

Bundle content is derived from the Kubernetes documentation, licensed under
[CC BY 4.0](https://creativecommons.org/licenses/by/4.0/); see [NOTICE](NOTICE).
