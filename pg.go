package stores

import (
	"context"
	"database/sql"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"

	"go.appointy.com/google/pb/intake_form"
	"go.appointy.com/google/pb/locations"
)

// pg is the IntakeFormStore Implementation using the Postgres as its Backend
type pg struct {
	*sql.DB
}

// NewPostgresIntakeFormStore creates a new IntakeForm Store with Postgres as its backend
func NewPostgresIntakeFormStore(db *sql.DB) IntakeFormStore {
	return &pg{db}
}

// AddIntakeForms adds a intake forms to the store
func (db *pg) AddIntakeForms(ctx context.Context, in *intake_form.IntakeForm) (*intake_form.IntakeFormIdentifier, error) {
	//user_email = ctx.Value("email").(string)
	userEmail := "lav@appointy.com"
	val := &intake_form.IntakeForm{
		Data: in.Data,
	}
	p, err := proto.Marshal(val)
	if err != nil {
		return nil, err
	}

	got := &intake_form.IntakeFormIdentifier{
		Base: &locations.LocationRoot{},
	}

	// Generate a new ID
	in.Id = uuid.New().String()
	// Insertion Query for IntakeForm
	const insrt = `INSERT INTO "intake_form"."intake_forms" ("program_id", "location_id", "id", "data", "created_by", "created_on", "is_deleted") VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING "program_id", "location_id", "id"`
	err = db.QueryRowContext(ctx, insrt, in.Base.ProgramId, in.Base.LocationId, in.Id, p, userEmail, time.Now(), false).Scan(&got.Base.ProgramId, &got.Base.LocationId, &got.Id)
	if err != nil {
		return nil, err
	}
	return got, err
}

// UpdateIntakeForm updates intake form
func (db *pg) UpdateIntakeForm(ctx context.Context, in *intake_form.IntakeForm) error {
	//user_email = ctx.Value("email").(string)
	userEmail := "lav@appointy.com"
	val := &intake_form.IntakeForm{
		Data: in.Data,
	}
	p, err := proto.Marshal(val)
	if err != nil {
		return err
	}

	const up = `UPDATE "intake_form"."intake_forms" SET "data" = $1, "updated_by" = $5, "updated_on" = $6  WHERE "id" = $2 AND "program_id" = $3 AND "location_id" = $4`
	_, err = db.ExecContext(ctx, up, p, in.Id, in.Base.ProgramId, in.Base.LocationId, userEmail, time.Now())
	if err != nil {
		return err
	}
	return err
}

// DeleteIntakeForm deletes intake form
func (db *pg) DeleteIntakeForm(ctx context.Context, ID string) error {
	//user_email = ctx.Value("email").(string)
	userEmail := "lav@appointy.com"
	sqlStatement := `UPDATE "intake_form"."intake_forms" SET "deleted_by" = $2, "deleted_on" = $3, "is_deleted" = $4  WHERE "id" = $1`
	_, err := db.ExecContext(ctx, sqlStatement, ID, userEmail, time.Now(), true)
	if err != nil {
		return err
	}
	return err
}

// GetIntakeForms get intake form by location
func (db *pg) GetIntakeForms(ctx context.Context, loc string) (*intake_form.IntakeFormList, error) {
	// Queries for selecting
	const stmt = `SELECT "program_id", "location_id", "id", "data" FROM "intake_form"."intake_forms" WHERE "location_id" = $1 AND "is_deleted" = $2`
	lst := &intake_form.IntakeFormList{
		Forms: []*intake_form.IntakeForm{},
	}

	rows, err := db.QueryContext(ctx, stmt, loc, false)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	p := make([]byte, 0)

	for rows.Next() {
		mp := intake_form.IntakeForm{
			Base: &locations.LocationRoot{},
		}
		mp2 := intake_form.IntakeForm{
			Base: &locations.LocationRoot{},
		}
		if err = rows.Scan(&mp.Base.ProgramId, &mp.Base.LocationId, &mp.Id, &p); err != nil {
			return nil, err
		}
		err = proto.Unmarshal(p, &mp2)
		mp.Data = mp2.Data
		if err != nil {
			return nil, err
		}
		lst.Forms = append(lst.Forms, &mp)
	}
	return lst, nil
}

// GetIntakeFormById get intake form by id
func (db *pg) GetIntakeFormById(ctx context.Context, id string) (*intake_form.IntakeForm, error) {
	// Queries for selecting
	const stmt = `SELECT "program_id", "location_id", "id", "data" FROM "intake_form"."intake_forms" WHERE "id" = $1 AND "is_deleted" = $2`
	v := &intake_form.IntakeForm{
		Base: &locations.LocationRoot{},
	}

	p := make([]byte, 0)

	err := db.QueryRowContext(ctx, stmt, id, false).Scan(&v.Base.ProgramId, &v.Base.LocationId, &v.Id, &p)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}

	v2 := &intake_form.IntakeForm{
		Base: &locations.LocationRoot{},
	}
	err = proto.Unmarshal(p, v2)
	if err != nil {
		return nil, err
	}
	v.Data = v2.Data

	return v, nil
}
