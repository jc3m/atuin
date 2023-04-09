// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: queries.sql

package tubul

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createEvent = `-- name: CreateEvent :one
INSERT INTO tubul.events (
  title, event_date
) VALUES (
  $1, $2
)
RETURNING id, title, event_date
`

type CreateEventParams struct {
	Title     string    `json:"title"`
	EventDate time.Time `json:"event_date"`
}

func (q *Queries) CreateEvent(ctx context.Context, arg CreateEventParams) (TubulEvent, error) {
	row := q.db.QueryRowContext(ctx, createEvent, arg.Title, arg.EventDate)
	var i TubulEvent
	err := row.Scan(&i.ID, &i.Title, &i.EventDate)
	return i, err
}

const deleteEvent = `-- name: DeleteEvent :exec
DELETE FROM tubul.events
WHERE id = $1
`

func (q *Queries) DeleteEvent(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteEvent, id)
	return err
}

const listEvents = `-- name: ListEvents :many
SELECT id, title, event_date FROM tubul.events
`

func (q *Queries) ListEvents(ctx context.Context) ([]TubulEvent, error) {
	rows, err := q.db.QueryContext(ctx, listEvents)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TubulEvent
	for rows.Next() {
		var i TubulEvent
		if err := rows.Scan(&i.ID, &i.Title, &i.EventDate); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateEvent = `-- name: UpdateEvent :one
UPDATE tubul.events
  SET title = $2, event_date = $3
WHERE id = $1
RETURNING id, title, event_date
`

type UpdateEventParams struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	EventDate time.Time `json:"event_date"`
}

func (q *Queries) UpdateEvent(ctx context.Context, arg UpdateEventParams) (TubulEvent, error) {
	row := q.db.QueryRowContext(ctx, updateEvent, arg.ID, arg.Title, arg.EventDate)
	var i TubulEvent
	err := row.Scan(&i.ID, &i.Title, &i.EventDate)
	return i, err
}
