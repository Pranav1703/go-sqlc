-- name: GetAllProducts :many
SELECT * from product;

-- name: GetProduct :one
SELECT * from product WHERE id=$1 LIMIT 1;

-- name: CreateProduct :one
INSERT INTO product (
    name, price, available
) VALUES (
    $1, $2, $3
) RETURNING *;