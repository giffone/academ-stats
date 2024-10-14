package postgres_test

import (
	"context"
	"excel_table/internal/domain"
	"excel_table/internal/repository/postgres"
	"fmt"
	"testing"

	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/assert"
)

// value     |value_title|queue|
// ----------+-----------+-----+
// go        |Go         |    0|
// golang    |Go         |    0|
// GO        |Go         |    0|
// Go        |Go         |    0|
// Golang    |Go         |    0|
// js        |Js         |    1|
// Js        |Js         |    1|
// javascript|Js         |    1|
// JavaScript|Js         |    1|
//           |other      | 1000|
// rust      |Rust       |    2|

func TestGetQueueOk(t *testing.T) {
	// create pool
	pool, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("failed to create mock pool: %v", err)
	}
	defer pool.Close()

	// make fake storage with nil pool, we need only method - MockConn()
	storage := postgres.NewStorage(pool)

	// make fake data
	rows := pool.NewRows([]string{"value", "value_title", "queue"}).
		AddRow("go", "Go", 0).
		AddRow("golang", "Go", 0).
		AddRow("", "other", 1000).
		AddRow("js", "Js", 1).
		AddRow("rust", "Rust", 2)

	// set query and rows that need to return
	q := fmt.Sprintf(postgres.QueryGetQueue, domain.SqlFieldLanguage)
	pool.ExpectQuery(q).WillReturnRows(rows)

	// testing function
	res, err := storage.GetQueue(context.Background())
	if err != nil {
		t.Fatalf("tested function: %v", err)
	}

	want := map[string]domain.TableQueue{
		"go":     {Title: "Go", Queue: 0},
		"golang": {Title: "Go", Queue: 0},
		"":       {Title: "other", Queue: 1000},
		"js":     {Title: "Js", Queue: 1},
		"rust":   {Title: "Rust", Queue: 2},
	}

	// check results
	equal := assert.Equal(t, want, res)
	if !equal {
		t.Fatal("not equal got and want")
	}

	if err := pool.ExpectationsWereMet(); err != nil {
		t.Fatalf("expectations were met: %v", err)
	}
}

func TestGetQueueEmpty(t *testing.T) {
	// create pool
	pool, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("failed to create mock pool: %v", err)
	}
	defer pool.Close()

	// make fake storage with nil pool, we need only method - MockConn()
	storage := postgres.NewStorage(pool)

	// make fake data
	rows := pool.NewRows([]string{"value", "value_title", "queue"})

	// set query and rows that need to return
	q := fmt.Sprintf(postgres.QueryGetQueue, domain.SqlFieldLanguage)
	pool.ExpectQuery(q).WillReturnRows(rows)

	// testing function
	res, err := storage.GetQueue(context.Background())
	if err != nil {
		t.Fatalf("tested function: %v", err)
	}

	want := map[string]domain.TableQueue{}

	// check results
	if !assert.Equal(t, want, res) {
		t.Fatal("not equal got and want")
	}

	if err := pool.ExpectationsWereMet(); err != nil {
		t.Fatalf("expectations were met: %v", err)
	}
}
