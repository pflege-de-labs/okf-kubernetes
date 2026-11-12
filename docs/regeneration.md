# Weekly regeneration

[`.github/workflows/regenerate.yml`](../.github/workflows/regenerate.yml) runs every Monday (and on
demand via *Run workflow*). It runs [`scripts/regenerate.sh`](../scripts/regenerate.sh) and keeps one
PR from the branch `chore/regenerate-catalog` open with the result.

## What the script does

1. Sparse-clones kubernetes/website `main` (`WEBSITE_REF` overrides it). It stops if `catalog/about.md`
   already names that commit.
2. Checks whether the source files cited in `catalog/requirements/` changed since their pinned commits.
3. Runs `cmd/okfgen`, adds a `catalog/log.md` entry, and runs `go test` and `okf validate`. It fails on
   links to missing concepts (anchor-only warnings are expected) and lists shortcodes `okfgen` dropped in the PR.
4. No cited file changed: moves all citation pins to the new commit. Otherwise it leaves them and lists
   the changed files in the PR for a human to check the cited line ranges.
5. Writes the PR title and body to `$OUT_DIR`. A new Kubernetes minor version gets a `feat(catalog)`
   title and a notice to check the concept ids named in the plugin skill.

Run it locally: `SITE=<existing website checkout> OUT_DIR=/tmp/regen bash scripts/regenerate.sh`.

## The PR

- Commits and the PR come from the agent GitHub App: the workflow trades its OIDC token for an agent
  token at the token broker (Actions variable `AGENT_TOKEN_BROKER_ACTIONS_URL`, optional runner
  `AGENT_RUNNER`), as described in token-broker's `docs/actions.md`.
- Each run rebuilds the branch from `main` and force-pushes it. If the branch has commits the job did not
  make itself (committer other than the bot, which includes `agent-commit` and web edits), for example
  fixed requirement citations, the run skips the push and warns instead.
- After merging, bump the bundle `ref` of the MCP server deployment to the merge commit.
