#!/bin/bash
set -e -o pipefail

TARGET_HEAD=$(git rev-parse refs/remotes/origin/$1)
CURRENT_HEAD=$(git rev-parse HEAD)

if [[ "$TARGET_HEAD" == "$CURRENT_HEAD" ]]; then
  echo "Current commit is head of $1."
  exit 0
fi

echo "Current commit is NOT head of $1."
exit 1
