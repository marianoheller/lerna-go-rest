package infrastructure

import (
	"errors"
	"fmt"
	"math/rand"

	d "learn/internal/domain"
)

// Article fixture data
var articles = []*d.Article{
	{ID: "1", UserID: 100, Title: "Hi", Slug: "hi"},
	{ID: "2", UserID: 200, Title: "sup", Slug: "sup"},
	{ID: "3", UserID: 300, Title: "alo", Slug: "alo"},
	{ID: "4", UserID: 400, Title: "bonjour", Slug: "bonjour"},
	{ID: "5", UserID: 500, Title: "whats up", Slug: "whats-up"},
}

// User fixture data
var users = []*d.User{
	{ID: 100, Name: "Peter"},
	{ID: 200, Name: "Julia"},
}

type RawRepository struct{}

func NewRawRepository() *RawRepository {
	return &RawRepository{}
}

func (r *RawRepository) CreateArticle(articleReq *d.CreateArticleRequest) (*d.Article, error) {
	article := d.Article{
		ID:     fmt.Sprintf("%d", rand.Intn(100)+10),
		UserID: articleReq.UserID,
		Title:  articleReq.Title,
		Slug:   articleReq.Slug,
	}
	articles = append(articles, &article)
	return &article, nil
}

func (r *RawRepository) GetArticle(id string) (*d.Article, error) {
	for _, a := range articles {
		if a.ID == id {
			return a, nil
		}
	}
	return nil, errors.New("article not found.")
}

func (r *RawRepository) GetArticles() ([]*d.Article, error) {
	return articles, nil
}

func (r *RawRepository) UpdateArticle(id string, articleReq *d.UpdateArticleRequest) (*d.Article, error) {
	for i, a := range articles {
		if a.ID == id {
			articles[i].Title = articleReq.Title
			articles[i].Slug = articleReq.Slug
			return articles[i], nil
		}
	}
	return nil, errors.New("article not found.")
}

func (r *RawRepository) RemoveArticle(id string) (*d.Article, error) {
	for i, a := range articles {
		if a.ID == id {
			articles = append((articles)[:i], (articles)[i+1:]...)
			return a, nil
		}
	}
	return nil, errors.New("article not found.")
}

func (r *RawRepository) GetUser(id int64) (*d.User, error) {
	for _, u := range users {
		if u.ID == id {
			return u, nil
		}
	}
	return nil, errors.New("user not found.")
}
