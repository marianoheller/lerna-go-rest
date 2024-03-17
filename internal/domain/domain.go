package domain

type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Article struct {
	ID     string `json:"id"`
	UserID int64  `json:"user_id"`
	Title  string `json:"title"`
	Slug   string `json:"slug"`
}

type CreateArticleRequest struct {
	UserID int64  `json:"user_id"`
	Title  string `json:"title"`
	Slug   string `json:"slug"`
}

type UpdateArticleRequest struct {
	Title string `json:"title"`
	Slug  string `json:"slug"`
}

type Repository interface {
	CreateArticle(articleReq *CreateArticleRequest) (*Article, error)

	GetArticle(id string) (*Article, error)

	GetArticles() ([]*Article, error)

	UpdateArticle(id string, articleReq *UpdateArticleRequest) (*Article, error)

	RemoveArticle(id string) (*Article, error)

	GetUser(id int64) (*User, error)
}

type DomainService struct {
	repository Repository
}

func NewDomainService(repository Repository) *DomainService {
	return &DomainService{
		repository: repository,
	}
}

func (s *DomainService) NewArticle(articleReq *CreateArticleRequest) (*Article, error) {
	return s.repository.CreateArticle(articleReq)
}

func (s *DomainService) GetArticles() ([]*Article, error) {
	return s.repository.GetArticles()
}

func (s *DomainService) GetArticle(id string) (*Article, error) {
	return s.repository.GetArticle(id)
}

func (s *DomainService) RemoveArticle(id string) (*Article, error) {
	return s.repository.RemoveArticle(id)
}

func (s *DomainService) GetUser(id int64) (*User, error) {
	return s.repository.GetUser(id)
}
