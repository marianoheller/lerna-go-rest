
-- name: GetArticles :many
SELECT * FROM articles
ORDER BY title;

-- name: GetArticle :one
SELECT * FROM articles
WHERE id = ?;

-- name: CreateArticle :one
INSERT INTO articles (user_id, title, slug)
VALUES (?, ?, ? )
RETURNING *;

-- name: UpdateArticle :one
UPDATE articles
set
    title = ?,
    slug = ?
WHERE id = ?
RETURNING *;

-- name: DeleteArticle :one
DELETE FROM articles
WHERE id = ?
RETURNING *;