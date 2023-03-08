package db

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func createPatientProfile(t *testing.T) Patient {
	arg := CreatePatientParams{
		FirstName: "",
		LastName:  "",
	}
	patient, err := testQueries.CreatePatient(ctx, arg)
	require.NoError(t, err)
	require.NotEmpty(t, patient)

	require.Equal(t, arg.FirstName, patient.FirstName)
	require.Equal(t, arg.LastName, patient.LastName)

	return patient

}

// func TestCreatePatient(t *testing.T) {
// 	createPatientProfile(t)
// }

// func TestGetPatient(t *testing.T) {
// 	patient := createPatientProfile(t)

// 	patient.PatientID = 1

// 	patient1, err := testQueries.GetPatient(ctx, patient.PatientID)

// 	require.NoError(t, err)
// 	require.NotEmpty(t, patient)

// }

func TestCreatePatientProfile(t *testing.T) {
	createPatientProfile(t)
}

func TestDeletePatient(t *testing.T) {
	arg := DeletePatientParams{
		FirstName: "",
		LastName:  "",
	}

	err := testQueries.DeletePatient(ctx, arg)

	require.NoError(t, err)

}
