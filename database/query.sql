-- name: GetAllProducts :many
SELECT * from product;

-- name: GetProduct :one
SELECT * from product WHERE id=$1 LIMIT 1;