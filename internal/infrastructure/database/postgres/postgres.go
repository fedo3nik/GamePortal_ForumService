package postgres

import (
	"context"
	"log"

	"github.com/fedo3nik/GamePortal_ForumService/internal/domain/entities"
	"github.com/jackc/pgx/v4/pgxpool"
)

func InsertForum(ctx context.Context, p *pgxpool.Pool, f *entities.Forum) (int, error) {
	var id int

	conn, err := p.Acquire(ctx)
	if err != nil {
		log.Printf("Create connection from pool error: %v", err)
		return 0, err
	}

	defer conn.Release()

	row := conn.QueryRow(ctx, "INSERT INTO forums (userId, title, topic, forum_text) VALUES($1, $2, $3, $4) RETURNING id",
		&f.UserID, &f.Title, &f.Topic, &f.Text)

	err = row.Scan(&id)
	if err != nil {
		log.Printf("Scan error: %v", err)
		return 0, err
	}

	return id, nil
}

func SelectForum(ctx context.Context, p *pgxpool.Pool, id int) (*entities.Forum, error) {
	var f entities.Forum

	conn, err := p.Acquire(ctx)
	if err != nil {
		log.Printf("Create connection from pool error: %v", err)
		return nil, err
	}

	defer conn.Release()

	err = conn.QueryRow(ctx, "SELECT id, userId, title, topic, forum_text FROM forums WHERE id=$1", id).
		Scan(&f.ID, &f.UserID, &f.Title, &f.Topic, &f.Text)
	if err != nil {
		log.Printf("Select error: %v", err)
		return nil, err
	}

	return &f, nil
}
