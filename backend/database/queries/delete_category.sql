-- name: ArchiveCategory :exec
UPDATE categories
SET archived_at = CURRENT_TIMESTAMP
WHERE id = $1 AND user_id = $2 AND archived_at IS NULL;