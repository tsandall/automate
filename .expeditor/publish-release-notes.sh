#!/bin/bash

set -eou pipefail

source_channel=$EXPEDITOR_PROMOTABLE

# Download the manifest
aws s3 cp "s3://chef-automate-artifacts/${source_channel}/latest/automate/manifest.json" manifest.json --profile chef-cd

build_version=$(jq -r -c ".build"  manifest.json)

git clone "https://x-access-token:${GITHUB_TOKEN}@github.com/chef/automate.wiki.git"

pushd ./automate.wiki
  # Publish release notes to S3
  aws s3 cp Pending-Release-Notes.md "s3://chef-automate-artifacts/release-notes/automate/${build_version}.md" --acl public-read --content-type "text/plain" --profile chef-cd
  aws s3 cp Pending-Release-Notes.md "s3://chef-automate-artifacts/${EXPEDITOR_TARGET_CHANNEL}/latest/automate/release-notes.md" --acl public-read --content-type "text/plain" --profile chef-cd

  # Reset "Pending Release Notes" wiki page
  cat >./Pending-Release-Notes.md <<EOH
## New Features
-

## Improvements
-

## Bug Fixes
-

## Backward Incompatibilities
-
EOH

  # Push changes back up to GitHub
  git add .
  git commit -m "Release Notes for promoted build $build_version"
  git push origin master
popd

rm -rf automate.wiki
