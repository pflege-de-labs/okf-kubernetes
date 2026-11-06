# okfgen mapping rules

`cmd/okfgen` converts `content/en/docs/{concepts,tasks,reference}` of a
`kubernetes/website` checkout into OKF concepts.

## Files

- `foo/bar.md` becomes `catalog/foo/bar.md`.
- Section pages `foo/_index.md` and page bundles `foo/index.md` become `catalog/foo.md`, since `index.md` is reserved in OKF. A generated `catalog/foo/index.md` lists the directory.
- Skipped: `headless` pages, pages with `_build.render: never|false`, and the ~490 feature-gate stubs, which are rendered into the feature-gate tables on `reference/command-line-tools-reference/feature-gates.md` and `feature-gates-removed.md`.

## Frontmatter

| OKF field | Source |
|-----------|--------|
| `type` | `content_type`: `concept`, `task`, `tutorial`, `reference`, `api-reference`, `tool-reference`; `glossary-term` for glossary entries; `section` for `_index.md` without a type |
| `title` | `title` |
| `description` | `description`, else glossary `short_description`, else the first sentence of the body |
| `resource` | Published kubernetes.io URL |
| `sources` | Commit-pinned GitHub source file and published URL |
| `generated` | `process:okf-kubernetes-okfgen` and the `-at` time |
| `kubernetes_version`, `source_commit`, `source_path` | Release from `hugo.toml`, checkout commit, file |
| `api_metadata`, `min_kubernetes_server_version` | Copied when present |

## Body

- Hugo shortcodes are rendered to Markdown: glossary tooltips become links to `/reference/glossary/<term>.md`; `glossary_definition`, `include`, tab includes and `code_sample` are inlined; `feature-state` becomes a "Feature state" line (resolving feature gates to their current stage); notes, tabs, headings, `skew` and `param` are rendered as text. Unknown shortcodes are dropped and reported.
- Site links (`/docs/...`, relative links, `ref`/`relref`, aliases) become bundle links when the target page is in the bundle, otherwise absolute `https://kubernetes.io` URLs.
- HTML comments are removed. Custom heading ids (`{#id}`) are kept so in-page anchors such as `names.md#dns-label-names` keep working in renderers that support them.
