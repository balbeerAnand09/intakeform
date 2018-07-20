package stores_test

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
	"go.appointy.com/google/intakeform/internal/stores"
)

func TestPostgresIntakeFormStore(t *testing.T) {
	const conn = "user=postgres password=qwertyuiop dbname=google sslmode=disable"
	db, err := sql.Open("postgres", conn)

	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}

	s := stores.NewPostgresIntakeFormStore(db)

	RunIntakeFormStoreTests(s, t)

	// clearing the database
	/*_, err = db.Query(`DELETE FROM "intake_form"."intake_forms"`)
	if err != nil {
		log.Fatalln("Not Deleted", err)
	}*/

}

/*
DataBase: "Google"
Schema: "intake_form"
Table: "intake_forms"
Columns: {
	program_id character varying 100,
	location_id character varying 100,
	id character varying 100,
	data byte array,
	created_by character varying 100,
	created_on character varying 100,
	updated_by character varying 100,
	updated_on character varying 100,
	deleted_by character varying 100,
	deleted_on character varying 100,
	is_deleted bool,
}
*/
