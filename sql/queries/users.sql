-- name: CreateUsers :one
INSERT INTO users (
  id,
  username,
  hash_pass,
  created_at,
  updated_at,
  role,
  section,
  email
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
