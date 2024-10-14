package parser_test

import (
	"encoding/json"
	"excel_table/internal/domain/request"
	"excel_table/internal/helper/parser"
	"excel_table/internal/helper/test"
	"excel_table/internal/repository/pb/excel_table"
	"testing"
)

func TestTopCadets(t *testing.T) {
	gotSrc, err := test.ReadFile[excel_table.TopCadetsRequest]("../test/test_files/src/top_cadets_grpc.json")
	if err != nil {
		t.Fatalf("can not read got_src file: %s", err)
	}

	tc, err := parser.TopCadets(gotSrc)
	if err != nil {
		t.Fatalf("parser request: %s", err)
	}

	wantSrc, err := test.ReadFile[request.TopCadets]("../test/test_files/want/top_cadets.json")
	if err != nil {
		t.Fatalf("can not read want_src file: %s", err)
	}

	got, err := json.Marshal(tc)
	if err != nil {
		t.Fatalf("marshal top cadets: got: %s", err)
	}
	want, err := json.Marshal(wantSrc)
	if err != nil {
		t.Fatalf("marshal top cadets: want: %s", err)
	}

	if string(got) != string(want) {
		t.Fatalf("json is not equal")
	}
}
