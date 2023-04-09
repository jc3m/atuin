-- name: ListEvents :many
SELECT * FROM tubul.events;

-- name: CreateEvent :one
INSERT INTO tubul.events (
  title, event_date
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdateEvent :one
UPDATE tubul.events
  SET title = $2, event_date = $3
WHERE id = $1
RETURNING *;

-- name: DeleteEvent :exec
DELETE FROM tubul.events
WHERE id = $1;
