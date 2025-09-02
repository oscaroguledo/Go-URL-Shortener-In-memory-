#!/usr/bin/env bash
set -e

# === Load .env ===
if [ -f .env ]; then
  set -a  # Automatically export all variables
  # Read only valid lines: non-empty and not starting with #
  grep -v '^\s*#' .env | grep -v '^\s*$' | while IFS= read -r line; do
    # shellcheck disable=SC1090
    eval "export $line"
  done
  set +a

else
  echo ".env file not found."
  exit 1
fi

# === Required vars ===
DB_URL="${DATABASE_URL}"
DB_TYPE="${DATABASE_TYPE}"
MIGRATIONS_PATH="./migrations"

# === Usage ===
usage() {
  echo "Usage: $0 <command> [n|version]"
  echo ""
  echo "Commands:"
  echo "  up [n]      – Apply all or next n migrations"
  echo "  down [n]    – Rollback last or last n migrations"
  echo "  goto <ver>  – Migrate to a specific version"
  echo "  version     – Print current migration version"
  echo "  drop        – Drop everything (CAREFUL!)"
  echo ""
  echo "Example:"
  echo "  $0 up"
  echo "  $0 down 1"
  echo "  $0 goto 5"
  exit 1
}

# === Parse input ===
CMD="$1"
ARG="$2"

if [ -z "$CMD" ]; then
  usage
fi

# === Run migration commands ===
case "$CMD" in
    new)
        if [ -z "$ARG" ]; then
        echo "Please provide a migration name (e.g., create_users_table)."
        usage
        fi
        migrate create -ext sql -dir "$MIGRATIONS_PATH" "$ARG"
        ;;

    up)
        migrate -database "$DB_URL" -path "$MIGRATIONS_PATH" up ${ARG:-}
        ;;
    down)
        migrate -database "$DB_URL" -path "$MIGRATIONS_PATH" down ${ARG:-}
        ;;
    goto)
        if [ -z "$ARG" ]; then
        echo "Please specify a version."
        usage
        fi
        migrate -database "$DB_URL" -path "$MIGRATIONS_PATH" goto "$ARG"
        ;;
    version)
        migrate -database "$DB_URL" -path "$MIGRATIONS_PATH" version
        ;;
    drop)
        migrate -database "$DB_URL" -path "$MIGRATIONS_PATH" drop -f
        ;;
    *)
        usage
        ;;
esac
