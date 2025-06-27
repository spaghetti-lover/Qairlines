-- name: CreateNews :one
INSERT INTO "news" (
    title,
    description,
    content,
    image,
    author_id
  )
VALUES ($1, $2, $3, $4, $5)
RETURNING *;
-- name: GetNews :one
SELECT *
FROM "news"
WHERE id = $1
LIMIT 1;
-- name: ListNews :many
SELECT *
FROM "news"
ORDER BY id
LIMIT $1 OFFSET $2;
-- name: RemoveAuthorFromBlogPosts :exec
UPDATE "news"
SET author_id = NULL,
  updated_at = NOW()
WHERE author_id = $1;
-- name: DeleteNews :execrows
DELETE FROM "news"
WHERE id = $1;
-- name: UpdateNews :one
UPDATE "news"
SET title = $1,
  description = $2,
  content = $3,
  image = $4,
  author_id = $5,
  updated_at = $6
WHERE id = $7
RETURNING id,
  title,
  description,
  content,
  image,
  author_id,
  created_at,
  updated_at;