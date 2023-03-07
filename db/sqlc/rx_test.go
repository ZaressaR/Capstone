package db

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreatePatientRoutine(t *testing.T) {
	rx := NewRX(testDB)

	patient1 := createPatientProfile(t)

	errs := make(chan error)
	results := make(chan CreatePatientRoutineResults)

	for i := 0; i < 2; i++ { // 2 concurrent requests to create a patient routine
		go func() {
			result, err := rx.CreatePatientRoutine(ctx, CreatePatientRoutineParams{
				FirstName:    patient1.FirstName,
				LastName:     patient1.LastName,
				RxName:       "",
				Administered: time.Time{},
			})
			errs <- err
			results <- result
		}()
	}
	// Check that we get the same patient back
	for i := 0; i < 2; i++ {
		err := <-errs
		require.Error(t, err)

		result := <-results
		require.NotEmpty(t, result)

		// check that the patient has the correct fields
		require.Equal(t, patient1.FirstName, result.FirstName)
	}
}
