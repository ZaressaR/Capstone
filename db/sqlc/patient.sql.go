// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: patient.sql

package db

import (
	"context"
)

const createPatient = `-- name: CreatePatient :one
INSERT INTO patient(
    first_name,
    last_name
)VALUES(
    $1, $2
) RETURNING patient_id, first_name, last_name
`

type CreatePatientParams struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (q *Queries) CreatePatient(ctx context.Context, arg CreatePatientParams) (Patient, error) {
	row := q.db.QueryRowContext(ctx, createPatient, arg.FirstName, arg.LastName)
	var i Patient
	err := row.Scan(&i.PatientID, &i.FirstName, &i.LastName)
	return i, err
}

const deletePatient = `-- name: DeletePatient :exec
DELETE FROM patient WHERE first_name = $1
`

func (q *Queries) DeletePatient(ctx context.Context, firstName string) error {
	_, err := q.db.ExecContext(ctx, deletePatient, firstName)
	return err
}

const getPatient = `-- name: GetPatient :one
SELECT patient_id, first_name, last_name FROM patient WHERE first_name = $1
`

func (q *Queries) GetPatient(ctx context.Context, firstName string) (Patient, error) {
	row := q.db.QueryRowContext(ctx, getPatient, firstName)
	var i Patient
	err := row.Scan(&i.PatientID, &i.FirstName, &i.LastName)
	return i, err
}
