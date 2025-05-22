-- name: CreateNews :one
INSERT INTO "news" (
  slug,
  image_url,
  title,
  description,
  author,
  content
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetNews :one
SELECT * FROM "news"
WHERE news_id = $1 LIMIT 1;

-- name: GetAllNews :many
SELECT * FROM "news";

-- name: ListNews :many
SELECT * FROM "news"
ORDER BY news_id
LIMIT $1
OFFSET $2;