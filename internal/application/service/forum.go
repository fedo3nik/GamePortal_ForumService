package service

import (
	"context"
	"log"
	"strconv"

	service "github.com/fedo3nik/GamePortal_ForumService/internal/application"

	"github.com/fedo3nik/GamePortal_ForumService/internal/domain/entities"
	"github.com/fedo3nik/GamePortal_ForumService/internal/infrastructure/database/postgres"
	e "github.com/fedo3nik/GamePortal_ForumService/internal/util/error"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

type Forum interface {
	AddForum(ctx context.Context, title string, topic string, text string, token string) (*entities.Forum, error)
	GetForum(ctx context.Context, ID string) (*entities.Forum, error)
}

type ForumService struct {
	Pool       *pgxpool.Pool
	AccessKey  string
	RefreshKey string
}

func (f ForumService) AddForum(ctx context.Context, title, topic, text, token string) (*entities.Forum, error) {
	userID, err := service.ValidateAccessToken(token, f.AccessKey)
	if err != nil {
		return nil, err
	}

	var forum entities.Forum

	forum.Title = title
	forum.Topic = topic
	forum.Text = text
	forum.UserID = userID

	id, err := postgres.InsertForum(ctx, f.Pool, &forum)
	if err != nil {
		log.Printf("DB: %v", err)
		return nil, errors.Wrap(e.ErrDB, "insert")
	}

	forum.ID = id

	return &forum, nil
}

func (f ForumService) GetForum(ctx context.Context, strID string) (*entities.Forum, error) {
	id, err := strconv.Atoi(strID)
	if err != nil {
		log.Printf("Convert error: %v", err)
		return nil, err
	}

	forum, err := postgres.SelectForum(ctx, f.Pool, id)
	if err != nil {
		return nil, errors.Wrap(e.ErrDB, "select")
	}

	return forum, nil
}

func NewForumService(pool *pgxpool.Pool, accessKey, refreshKey string) *ForumService {
	return &ForumService{Pool: pool, AccessKey: accessKey, RefreshKey: refreshKey}
}
