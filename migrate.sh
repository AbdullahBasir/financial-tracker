#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ENV_FILE="$SCRIPT_DIR/.env"

if [[ -f "$ENV_FILE" ]]; then
  set -a
  source "$ENV_FILE"
  set +a
else
  echo "Error: .env file not found at $ENV_FILE"
  exit 1
fi

if [[ -z "${DB_URL:-}" ]]; then
  echo "Error: DB_URL is not set. Add it to your .env file."
  exit 1
fi

if [[ $# -lt 1 ]]; then
  echo "Usage: $0 {up|down}"
  exit 1
fi

case "$1" in
  up)
    goose --dir sql/schema postgres "$DB_URL" up
    ;;
  down)
    goose --dir sql/schema postgres "$DB_URL" down
    ;;
  *)
    echo "Unknown command: $1"
    echo "Usage: $0 {up|down}"
    exit 1
    ;;
esac