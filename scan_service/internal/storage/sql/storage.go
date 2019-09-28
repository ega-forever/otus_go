package db

import (
	"context"
	"fmt"
	"github.com/ega-forever/otus_go/scan_service/internal/domain/models"
	"github.com/jackc/pgx/v4"
	"log"
)

type Storage struct {
	ctx context.Context
	db  *pgx.Conn
}

func New(uri string) *Storage {

	ctx := context.Background()
	db, err := pgx.Connect(ctx, uri) // *sql.DB
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}

	return &Storage{ctx: ctx, db: db}
}

func (storage *Storage) Migrate() error {

	query := `create table if not exists events (
    id SERIAL,
    text text,
    timestamp double precision
)`

	_, err := storage.db.Exec(storage.ctx, query) // sql.Result
	if err != nil {
		return err
	}

	fmt.Println("migrated tables")
	return nil
}

func (storage *Storage) SaveEvent(event *models.Event) (*models.Event, error) {

	query := `insert into events(text, timestamp) values($1, $2) RETURNING id`

	var id int64
	err := storage.db.QueryRow(storage.ctx, query, event.Text, event.Timestamp).Scan(&id)

	if err != nil {
		return nil, err
	}

	ev := models.Event{Id: id, Text: event.Text, Timestamp: event.Timestamp}

	return &ev, nil
}

func (storage *Storage) FindEventsAfterTimestamp(timestamp int64) ([]*models.Event, error) {

	query := `select * from events where timestamp > $1`
	rows, err := storage.db.Query(storage.ctx, query, timestamp)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	envs := make([]*models.Event, 0)

	for rows.Next() {
		ev := models.Event{}
		err := rows.Scan(&ev.Id, &ev.Text, &ev.Timestamp)

		if err != nil {
			return nil, err
		}

		envs = append(envs, &ev)
	}

	return envs, nil
}
