package parser_test

import (
	"academ_stats/internal/helper/test"
	"testing"
)

func TestTopCadetsGraphql(t *testing.T) {
	if _, err := test.CheckTopCadetsGraphql(
		"../test/test_files/src/top_cadets_graphql.json",
		"../test/test_files/want/top_cadets_graphql.json",
		"../test/test_files/src/hours_grpc.json",
	); err != nil {
		t.Fatalf("test top cadets graphql: %s", err)
	}
}
