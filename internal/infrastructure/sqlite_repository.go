package infrastructure

import (
	"context"
	"errors"

	"database/sql"
	d "learn/internal/domain"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
	lo "github.com/samber/lo"

	g "learn/generated"
)

type SQLiteRepository struct {
	queries *g.Queries
	ctx     *context.Context
}

func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	ctx := context.Background()

	return &SQLiteRepository{
		queries: g.New(db),
		ctx:     &ctx,
	}
}

func (r *SQLiteRepository) CreateArticle(articleReq *d.CreateArticleRequest) (*d.Article, error) {
	created, err := r.queries.CreateArticle(*r.ctx, g.CreateArticleParams{
		UserID: articleReq.UserID,
		Title:  articleReq.Title,
		Slug:   articleReq.Slug,
	})
	if err != nil {
		return nil, err
	}

	return toDomain(created), nil
}

func (r *SQLiteRepository) GetArticle(id string) (*d.Article, error) {
	parsedId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}

	article, err := r.queries.GetArticle(*r.ctx, parsedId)

	return toDomain(article), errors.New("article not found.")
}

func (r *SQLiteRepository) GetArticles() ([]*d.Article, error) {
	foundArticles, err := r.queries.GetArticles(*r.ctx)

	result := lo.Map(foundArticles, func(article g.Article, index int) *d.Article {
		return toDomain(article)
	})

	return result, err
}

func (r *SQLiteRepository) UpdateArticle(id string, articleReq *d.UpdateArticleRequest) (*d.Article, error) {
	parsedId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}

	updated, err := r.queries.UpdateArticle(*r.ctx, g.UpdateArticleParams{
		ID:    parsedId,
		Title: articleReq.Title,
		Slug:  articleReq.Slug,
	})

	if err != nil {
		return nil, err
	}

	return toDomain(updated), nil
}

func (r *SQLiteRepository) RemoveArticle(id string) (*d.Article, error) {
	parsedId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}

	deleted, err := r.queries.DeleteArticle(*r.ctx, parsedId)
	if err != nil {
		return nil, err
	}

	return toDomain(deleted), nil
}

func (r *SQLiteRepository) GetUser(id int64) (*d.User, error) {
	foundUser, err := r.queries.GetUser(*r.ctx, id)

	if err != nil {
		return nil, err
	}

	mappedUser := d.User{
		ID:   foundUser.ID,
		Name: foundUser.Name,
	}

	return &mappedUser, nil
}

func toDomain(record g.Article) *d.Article {
	return &d.Article{
		ID:     strconv.Itoa(int(record.ID)),
		UserID: record.UserID,
		Title:  record.Title,
		Slug:   record.Slug,
	}
}
