-- name: CreateNews :one
INSERT INTO "news" (
  title,
  description,
  content,
  image,
  author_id
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetNews :one
SELECT * FROM "news"
WHERE id = $1 LIMIT 1;

-- name: GetAllNewsWithAuthor :many
SELECT *
FROM "news" n
JOIN "users" u ON n.author_id = u.user_id
ORDER BY n.created_at DESC;

-- name: ListNews :many
SELECT * FROM "news"
ORDER BY created_at DESC
LIMIT $1
OFFSET $2;

-- name: RemoveAuthorFromBlogPosts :exec
UPDATE "news"
SET author_id = NULL,
    updated_at = NOW()
WHERE author_id = $1;

-- name: DeleteNews :execrows
DELETE FROM "news"
WHERE id = $1;