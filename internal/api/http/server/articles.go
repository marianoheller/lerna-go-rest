package server

import (
	"net/http"

	d "learn/internal/domain"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type ArticlesApi struct {
	service *d.DomainService
}

func NewArticlesApi(service *d.DomainService, r *chi.Mux) {
	api := ArticlesApi{service: service}

	r.Route("/articles", func(r chi.Router) {
		r.With(paginate).Get("/", api.ListArticles)
		r.Post("/", api.CreateArticle) // POST /articles
		r.Route("/{articleID}", func(r chi.Router) {
			r.Get("/", api.GetArticle)       // GET /articles/123
			r.Delete("/", api.DeleteArticle) // DELETE /articles/123
		})
	})
}

func (api ArticlesApi) CreateArticle(w http.ResponseWriter, r *http.Request) {
	data := &ArticleRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	articleReq := data.CreateArticleRequest
	createdArticle, err := api.service.NewArticle(articleReq)

	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, api.NewArticleResponse(createdArticle))
}

func (api ArticlesApi) NewArticleResponse(article *d.Article) *ArticleResponse {
	resp := &ArticleResponse{Article: article}

	if resp.User == nil {
		if user, _ := api.service.GetUser(resp.UserID); user != nil {
			resp.User = NewUserPayloadResponse(user)
		}
	}

	return resp
}

// paginate is a stub, but very possible to implement middleware logic
// to handle the request params for handling a paginated request.
func paginate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// just a stub.. some ideas are to look at URL query params for something like
		// the page number, or the limit, and send a query cursor down the chain
		next.ServeHTTP(w, r)
	})
}

func (api ArticlesApi) ListArticles(w http.ResponseWriter, r *http.Request) {
	articles, err := api.service.GetArticles()
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
	}

	if err := render.RenderList(w, r, api.NewArticleListResponse(articles)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func (api ArticlesApi) GetArticle(w http.ResponseWriter, r *http.Request) {
	articleID := chi.URLParam(r, "articleID")
	article, err := api.service.GetArticle(articleID)

	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	render.Render(w, r, api.NewArticleResponse(article))
}

// TODO: update article

func (api ArticlesApi) DeleteArticle(w http.ResponseWriter, r *http.Request) {
	var err error

	// Assume if we've reach this far, we can access the article
	// context because this handler is a child of the ArticleCtx
	// middleware. The worst case, the recoverer middleware will save us.
	article := r.Context().Value("article").(*d.Article)

	article, err = api.service.RemoveArticle(article.ID)
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	render.Render(w, r, api.NewArticleResponse(article))
}

func (api ArticlesApi) NewArticleListResponse(articles []*d.Article) []render.Renderer {
	list := []render.Renderer{}
	for _, article := range articles {
		list = append(list, api.NewArticleResponse(article))
	}
	return list
}

//--
// Error response payloads & renderers
//--

// ErrResponse renderer type for handling all sorts of errors.
//
// In the best case scenario, the excellent github.com/pkg/errors package
// helps reveal information on the error, setting it on Err, and in the Render()
// method, using it to set the application-specific error code in AppCode.
type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}

func ErrRender(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 422,
		StatusText:     "Error rendering response.",
		ErrorText:      err.Error(),
	}
}

var ErrNotFound = &ErrResponse{HTTPStatusCode: 404, StatusText: "Resource not found."}
