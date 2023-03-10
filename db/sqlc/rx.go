package db

import (
	"context"
	"database/sql"
	"time"
)

// RX provides all functions to execute queries and actions on the database.

type RX struct {
	// The underlying database connection.
	*Queries // embedded
	db       *sql.DB
}

// this func is called in main.go and is passed the database connection. It returns a pointer to RX.

func NewRX(db *sql.DB) *RX {
	return &RX{
		db:      db,
		Queries: New(db),
	}
}

// this func is used to execute the transaction. It takes a context and a function
// that takes a pointer to Queries as a parameter. It returns an error.
// It begins a transaction, creates a new pointer to Queries, and passes it to the function.
// If the function returns an error, the transaction is rolled back and the error is returned.
// If the function does not return an error, the transaction is committed and nil is returned.

func (q *RX) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := q.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	qx := New(tx)
	err = fn(qx)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return rbErr
		}
		return err
	}
	return tx.Commit()
}

type CreatePatientRoutineParams struct {
	FirstName    string    `json:"firstName" binding:"required"`
	LastName     string    `json:"lastName" binding:"required"`
	RxName       string    `json:"rxname" binding:"required"`
	Administered time.Time `json:"administered" binding:"required"`
}

type CreatePatientRoutineResults struct {
	FirstName    Patient    `json:"firstName"`
	LastName     Patient    `json:"lastName"`
	RxName       Medication `json:"rxname"`
	Administered time.Time  `json:"administered"`
}

func (q *RX) CreatePatientRoutine(ctx context.Context, arg CreatePatientRoutineParams) (CreatePatientRoutineResults, error) {
	var result CreatePatientRoutineResults

	err := q.execTx(context.Background(), func(q *Queries) error {
		var err error
		result.FirstName, err = q.CreatePatient(ctx, CreatePatientParams{
			FirstName: arg.FirstName,
			LastName:  arg.LastName,
		})

		if err != nil {
			return err
		}

		result.RxName, err = q.CreateMedication(ctx, CreateMedicationParams{
			RxName:       arg.RxName,
			Administered: arg.Administered,
		})

		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}
