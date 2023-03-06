-- name: CreatePatient :one
INSERT INTO patient(
    first_name,
    last_name
)VALUES(
    $1, $2
) RETURNING *;