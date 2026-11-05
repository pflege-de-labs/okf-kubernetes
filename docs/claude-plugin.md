# Kubernetes knowledge with Claude

This repository is a Claude Code plugin marketplace named `okf-kubernetes` with one plugin, `kubernetes-okf`. Its skill tells Claude to answer upstream Kubernetes questions (name limits, label rules, size limits, ports, API fields, feature gates, kubectl, concepts, tasks) from this bundle via the pflege.de OKF MCP server instead of general knowledge, starting with the curated `kubernetes/requirements/*` concepts for validation questions.

The plugin lives here rather than in `pflege-de/mcp` because this bundle is ours; the Crossplane plugin is in `mcp` because its bundle is external.

## 1. Add the MCP server

The pflege.de server prefixes concept ids with the bundle name it was given (`--bundle kubernetes=...`, chart value `config.bundles.kubernetes`). The skill expects `kubernetes`, the same way the Crossplane skill in `mcp` expects `crossplane`; the bundle itself has no prefix (upstream okf-crossplane-v2 also uses bare ids such as `composition-developer-starter`). See [docs/crossplane-v2-okf.md](https://github.com/pflege-de/mcp/blob/main/docs/crossplane-v2-okf.md) in `pflege-de/mcp` for `claude mcp add`; for a local server, see "Validate" in the [README](../README.md).

## 2. Install the plugin

```sh
claude plugin marketplace add pflege-de-labs/okf-kubernetes   # or a local path to this checkout
claude plugin install kubernetes-okf@okf-kubernetes
```

Restart running sessions or run `/reload-plugins`. The skill is used automatically for Kubernetes questions, or explicitly with `/kubernetes-okf:kubernetes-okf`.

Updates: `claude plugin marketplace update okf-kubernetes && claude plugin update kubernetes-okf@okf-kubernetes`.

Team repository `.claude/settings.json`:

```json
{
  "extraKnownMarketplaces": {
    "okf-kubernetes": {
      "source": { "source": "github", "repo": "pflege-de-labs/okf-kubernetes" }
    }
  },
  "enabledPlugins": {
    "kubernetes-okf@okf-kubernetes": true
  }
}
```

The plugin contains only the skill; it does not register the MCP server.
