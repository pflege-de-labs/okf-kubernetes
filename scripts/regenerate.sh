#!/usr/bin/env bash
# Regenerates catalog/ from kubernetes/website and writes a PR title and body to $OUT_DIR.
# Exits 0 without changes when the catalog is already current. Details: docs/regeneration.md
set -euo pipefail

out=${OUT_DIR:-.regen}
site=${SITE:-$(mktemp -d)/website}
ref=${WEBSITE_REF:-main}
mkdir -p "$out"

field() { sed -n "s/^$1: //p" catalog/about.md; }
old=$(field source_commit)
old_version=$(field kubernetes_version)

if [ ! -d "$site/.git" ]; then
  git clone -q --filter=blob:none --sparse --branch "$ref" https://github.com/kubernetes/website.git "$site"
  git -C "$site" sparse-checkout set content/en/docs/concepts content/en/docs/tasks content/en/docs/reference \
    content/en/examples content/en/includes content/en/releases i18n/en data/releases
fi
new=$(git -C "$site" rev-parse HEAD)
if [ "$old" = "$new" ]; then
  echo "catalog is current at kubernetes/website@$new"
  exit 0
fi

# Curated requirements cite line ranges at pinned commits; a range stays valid only if its file is unchanged.
cites=$(grep -ho 'kubernetes/website/blob/[0-9a-f]\{40\}/[^#)" ]*' catalog/requirements/*.md | sed 's#.*/blob/##' | sort -u)
pins=$(cut -d/ -f1 <<<"$cites" | sort -u)
for c in $pins; do
  git -C "$site" cat-file -e "$c^{commit}" 2>/dev/null || git -C "$site" fetch -q origin "$c"
done
changed=$(while IFS=/ read -r c p; do
  git -C "$site" diff --quiet "$c" "$new" -- "$p" || echo "$p"
done <<<"$cites" | sort -u)
cited=$(cut -d/ -f2- <<<"$cites" | sort -u)

go run ./cmd/okfgen -site "$site" -out catalog 2> >(tee "$out/okfgen.log" >&2)
new_version=$(field kubernetes_version)
if [ -z "$changed" ]; then
  for c in $pins; do
    perl -pi -e "s/\Q$c\E/$new/g" catalog/requirements/*.md
  done
fi

entry="* **Update**: Regenerated concepts, tasks, reference and releases from kubernetes/website@${new:0:12} (Kubernetes $new_version) with \`cmd/okfgen\`."
# Joins an existing section for today instead of adding a second one.
TODAY=$(date -u +%F) ENTRY=$entry perl -0pi -e \
  's/^(# Catalog Update Log\n\n)(## \Q$ENV{TODAY}\E\n)?/$1 . "## $ENV{TODAY}\n$ENV{ENTRY}\n" . ($2 ? "" : "\n")/e' catalog/log.md

go test ./...
if ! uvx --from okf-ingest==0.13.1 okf validate catalog/ >"$out/validate.txt" 2>&1; then
  grep -E 'conformant|\[error' "$out/validate.txt" | head -20
  exit 1
fi
summary=$(grep -m1 conformant "$out/validate.txt")
# Anchor-only warnings are expected; a link to a missing page means upstream removed or moved it.
if grep 'unresolved link: [^#]' "$out/validate.txt"; then
  echo "links to missing concepts, see above"
  exit 1
fi
dropped=$(grep -o 'dropped unknown shortcode.*' "$out/okfgen.log" || true)

if [ "$old_version" != "$new_version" ]; then
  echo "feat(catalog): regenerate for Kubernetes $new_version" >"$out/title"
else
  echo "chore(catalog): regenerate from kubernetes/website@${new:0:12}" >"$out/title"
fi
{
  echo "Regenerates \`catalog/\` from kubernetes/website [\`${old:0:12}...${new:0:12}\`](https://github.com/kubernetes/website/compare/$old...$new) (Kubernetes $old_version → $new_version)."
  echo
  if [ "$old_version" != "$new_version" ]; then
    echo "**New minor release.** Check that the concept ids named in the plugin skill still exist and review removed or renamed pages."
    echo
  fi
  if [ -z "$changed" ]; then
    echo "Curated requirements: none of the $(wc -w <<<"$cited" | tr -d ' ') cited source files changed, so their citations now point to \`${new:0:12}\`."
  else
    echo "**Curated requirements need review.** These cited source files changed; all citations keep their pinned commits until the line ranges in \`catalog/requirements/\` are checked and moved by hand:"
    echo
    while read -r f; do
      echo "- \`$f\`"
    done <<<"$changed"
  fi
  if [ -n "$dropped" ]; then
    echo
    echo "**Content dropped by okfgen.** Unknown shortcodes; add handlers in \`cmd/okfgen/convert.go\` if they carry content:"
    echo
    while read -r d; do echo "- $d"; done <<<"$dropped"
  fi
  echo
  echo "Checks: \`go test\` passed; \`okf validate\`: $summary (same-page anchor warnings are expected)."
  echo
  echo "After merge, bump the bundle \`ref\` of the MCP server deployment to the merge commit."
} >"$out/body.md"
cat "$out/title"
