-- name: CreatePatient :one
INSERT INTO patient(
    first_name,
    last_name
)VALUES(
    $1, $2
) RETURNING *;


-- name: DeletePatient :exec
DELETE FROM patient WHERE first_name = $1;

-- name: GetPatient :one
SELECT * FROM patient WHERE first_name = $1;