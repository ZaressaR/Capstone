-- name: CreateMedication :one
INSERT INTO medication (
    patient_id,
    rx_name,
    administered
) VALUES (
    $1, $2, $3
)  RETURNING *;

