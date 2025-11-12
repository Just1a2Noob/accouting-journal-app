-- name: CreateUsers :one
INSERT INTO users (
  id,
  username,
  created_at,
  updated_at,
  email,
  section,
  role
) VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6,
  $7,
  $8
)
RETURNING *;
