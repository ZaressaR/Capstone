package db

import ()

const (
	dbDriver = "postgres"
	dbSource = "postgresql:///patient_profile?sslmode=disable"
	dbName   = "patient_profile"
)

var ctx = context.Background()
var testQueries *Queries

func TestMain(m *testing.M) {
	testdb, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	testQueries = New(testdb)
	os.Exit(m.Run())

}
