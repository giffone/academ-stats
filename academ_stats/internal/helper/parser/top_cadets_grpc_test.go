package parser_test

import (
	"academ_stats/internal/helper/parser"
	"academ_stats/internal/helper/test"
	"academ_stats/internal/repository/pb/excel_table"
	"encoding/json"
	"testing"
)

func TestTopCadetsGrpc(t *testing.T) {
	src, err := test.CheckTopCadetsGraphql(
		"../test/test_files/src/top_cadets_graphql.json",
		"../test/test_files/want/top_cadets_graphql.json",
		"../test/test_files/src/hours_grpc.json",
	)
	if err != nil {
		t.Fatalf("test top cadets grpc: %s", err)
	}

	tcGot := parser.TopCadetsGrpc(src)

	tcWant, err := test.ReadFile[excel_table.TopCadetsRequest]("../test/test_files/want/top_cadets_grpc.json")
	if err != nil {
		t.Fatalf("read want response: %s", err)
	}

	got, err := json.Marshal(tcGot)
	if err != nil {
		t.Fatalf("marshal top cadets: got: %s", err)
	}
	want, err := json.Marshal(tcWant)
	if err != nil {
		t.Fatalf("marshal top cadets: want: %s", err)
	}

	if string(got) != string(want) {
		t.Fatalf("json is not equal")
	}
}
