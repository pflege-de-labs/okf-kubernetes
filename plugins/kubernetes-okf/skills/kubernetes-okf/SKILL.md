---
name: kubernetes-okf
description: Use the pflege.de OKF MCP server as the primary knowledge source for upstream Kubernetes facts - object name formats and length limits, label/annotation rules, size limits, ports, API resource fields, feature gates and feature states, kubectl and component flags, concepts, tasks, and the glossary. Use when writing, reviewing, or validating Kubernetes manifests, or answering questions about Kubernetes behavior, limits, or APIs.
---

# Kubernetes OKF

For every upstream Kubernetes question, use the OKF MCP tools before answering or writing manifests. The server is usually registered as `okf` (tools appear as `mcp__okf__<tool>`). The Kubernetes bundle is generated from the Kubernetes documentation (kubernetes.io/docs: Concepts, Tasks, Reference) for one release; each concept states it in `kubernetes_version`.

## Workflow

1. Call `okf_search` with the key terms of the question. Results are concept ids such as `${user_config.bundle_name}/concepts/workloads/pods`: the server's bundle name (plugin setting `bundle_name`), then the path inside the bundle. If the results use a different bundle name, use it in place of `${user_config.bundle_name}/` for every id below.
2. For validation questions (names, lengths, character sets, sizes, ports, reserved values), go to the curated requirements first. They are short and cite exact source lines:

   | Topic | Concept id |
   |-------|-----------|
   | Name formats and per-kind length limits (e.g. CronJob 52, Service 63) | `${user_config.bundle_name}/requirements/object-names` |
   | Label and annotation keys, values, 256 KiB annotation total | `${user_config.bundle_name}/requirements/labels-and-annotations` |
   | Namespace names, initial and reserved namespaces | `${user_config.bundle_name}/requirements/namespaces` |
   | ConfigMap/Secret 1 MiB and other size limits | `${user_config.bundle_name}/requirements/size-limits` |
   | Port numbers and names, NodePort range | `${user_config.bundle_name}/requirements/ports-and-ranges` |

3. Resolve links in a concept body to ids: `${user_config.bundle_name}/` plus the link target resolved against the current concept's directory, without `.md` and without `#anchor`. Links starting with `/` are relative to the bundle root. `okf_relations` lists them for you (`links_to`) together with concepts that link back (`linked_by`). Links to `https://kubernetes.io/...` point to pages outside the bundle. Links ending in `index.md` are section listings, not concepts: use the section concept `<dir>` (e.g. `${user_config.bundle_name}/concepts/workloads`) or `okf_search` instead.
4. Use `okf_get` for the full text. Cite the concept's `sources` (`k8s-docs` is the published page, `k8s-docs-source` the pinned source file).

Other useful starting points:

| Topic | Concept id |
|-------|-----------|
| API resource fields (one page per kind) | `${user_config.bundle_name}/reference/kubernetes-api/<group>/<kind>-<version>`, e.g. `.../core/pod-v1` |
| Feature gates (default, stage, since, until) | `${user_config.bundle_name}/reference/command-line-tools-reference/feature-gates` |
| Removed feature gates | `${user_config.bundle_name}/reference/command-line-tools-reference/feature-gates-removed` |
| Well-known labels, annotations, taints | `${user_config.bundle_name}/reference/labels-annotations-taints` |
| Component flags | `${user_config.bundle_name}/reference/command-line-tools-reference/<component>`, e.g. `kubelet`, `kube-apiserver` |
| kubectl commands | `${user_config.bundle_name}/reference/kubectl/generated/kubectl_<command>` |
| Glossary terms | `${user_config.bundle_name}/reference/glossary/<term>`, e.g. `pod` |
| Deprecation policy | `${user_config.bundle_name}/reference/deprecation-policy` |

## Versions and feature states

- The bundle documents a single Kubernetes release (`kubernetes_version`). If the user's cluster runs a different version, say which version the answer is based on and point out version-dependent behavior instead of assuming it applies.
- Preserve feature-state lines such as `**Feature state:** Kubernetes v1.37 [beta] (feature gate X, disabled by default)`. Do not present an alpha or beta feature as generally available, and mention when a feature gate must be enabled.

## Evidence boundaries

Treat the OKF concepts as the primary source for upstream Kubernetes facts. Local project files and user-supplied material remain in scope.

Not covered by this bundle: managed offerings (EKS, GKE, AKS), distributions, Helm, operators, and third-party controllers. Use their own documentation and do not infer their behavior from upstream defaults. For Crossplane, use the Crossplane OKF skill.

If the bundle lacks the information, say so and name the gap instead of guessing a limit, field, or default. If the MCP server is unavailable, say so and fall back to general knowledge only with that caveat.
