package db

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func createPatientProfile(t *testing.T) (Patient, error) {
	arg := CreatePatientParams{
		FirstName: "",
		LastName:  "",
	}
	patient, err := testQueries.CreatePatient(ctx, arg)
	require.NoError(t, err)
	require.NotEmpty(t, patient)

	require.Equal(t, arg.FirstName, patient.FirstName)
	require.Equal(t, arg.LastName, patient.LastName)

	return patient, err

}

// func TestCreatePatient(t *testing.T) {
// 	createPatientProfile(t)
// }

func TestGetPatient(t *testing.T) {
	patient, err := createPatientProfile(t)
	require.NoError(t, err)
	require.NotEmpty(t, patient)

}
