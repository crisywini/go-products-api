-- name: CreateProduct :one
INSERT INTO products (name, description, price, stock, user_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetProduct :one
SELECT * FROM products
WHERE id = $1;

-- name: ListProducts :many
SELECT * FROM products
ORDER BY created_at DESC;

-- name: UpdateProduct :one
UPDATE products
SET name = $2, description = $3, price = $4, stock = $5, updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = $1;