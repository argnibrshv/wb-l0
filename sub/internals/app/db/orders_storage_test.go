package db

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetOrderByIDSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("mock fail", err)
	}
	defer db.Close()

	storage := NewOrdersStorage(db)

	rows := sqlmock.NewRows([]string{"id", "order_details"})

	expect := struct {
		ID            string
		order_details string
	}{
		ID:            "b563feb7b2b84b6test",
		order_details: `{"order_uid":"b563feb7b2b84b6test","testjson":"testjson"}`,
	}

	rows.AddRow(expect.ID, expect.order_details)

	mock.
		ExpectQuery("SELECT id, order_details FROM orders WHERE id = ?").
		WithArgs(expect.ID).
		WillReturnRows(rows)

	item, err := storage.GetOrderByID(expect.ID)
	if err != nil {
		t.Errorf("fail to get data: #{err}")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations were not meet: #{err}")
	}

	if item != expect.order_details {
		t.Errorf("result not match, want %v, got %v", expect.order_details, item)
	}
}

func TestGetOrderByIDFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("mock fail", err)
	}
	defer db.Close()

	storage := NewOrdersStorage(db)

	rows := sqlmock.NewRows([]string{"id", "order_details"})

	expect := struct {
		ID            string
		order_details string
	}{
		ID:            "b563feb7b2b84b6test",
		order_details: `{"order_uid":"b563feb7b2b84b6test","testjson":"testjson"}`,
	}

	rows.AddRow(expect.ID, expect.order_details)

	mock.
		ExpectQuery("SELECT id, order_details FROM orders WHERE id = ?").
		WithArgs("b563feb7b2b84b6test0").
		WillReturnError(errors.New("sql: no rows in result set"))

	_, err = storage.GetOrderByID("b563feb7b2b84b6test0")
	if err == nil {
		t.Errorf("must return error: sql: no rows in result set")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations were not meet: %v", err)
	}
}
