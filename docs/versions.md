# Versions and release lifecycle

The bundle documents one Kubernetes release: the one in the website checkout's `hugo.toml`
(`latest`). It should be the newest release, not the one our clusters run, so the agent can
point out features that an upgrade would enable.

## How the agent learns which version applies

- `about` names the bundle release and source commit. It is a concept rather than frontmatter
  because the MCP server's `okf_get` does not return extra frontmatter keys such as `kubernetes_version`.
- Feature gates: `reference/command-line-tools-reference/feature-gates` lists stage, default and
  since/until for every stage of every gate; `feature-gates-removed` covers removed gates.
- `reference/using-api/deprecation-guide` lists API removals per release.
- `releases/lifecycle` lists release, maintenance-mode and end-of-life dates of every minor release
  from `data/releases/{schedule,eol}.yaml`. It contains dates and the status rule, not a status:
  a stored status would go stale between regenerations, and the upstream data lags (a branch can stay
  in `schedule.yaml` after its end of life).
- `releases/version-skew-policy` and `releases/patch-releases` (support period) are converted pages.

The skill ([SKILL.md](../plugins/kubernetes-okf/skills/kubernetes-okf/SKILL.md)) tells the agent to
determine the target cluster version itself (user, project or `kubectl version`). The deployed version
is not stored in this public, upstream-only bundle.

## Not covered

- Unreleased versions: planned features of the next minor can slip; the lifecycle page only links the
  `kubernetes/sig-release` schedule.
- API field changes that are neither behind a feature gate nor in the deprecation guide.
