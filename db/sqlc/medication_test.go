package db

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createMedicationProfile(t *testing.T) Medication {
	arg := CreateMedicationParams{
		RxName:       "",
		Administered: time.Weekday(0),
	}
	medication, err := testQueries.CreateMedication(ctx, arg)
	require.NoError(t, err)
	require.NotEmpty(t, medication)

	require.Equal(t, arg.RxName, medication.RxName)
	require.Equal(t, arg.Administered, medication.Administered)

	return medication

}

func TestCreateMedicationProfile(t *testing.T) {
	createMedicationProfile(t)

}
