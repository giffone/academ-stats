package parser_test

import (
	"academ_stats/internal/domain/request"
	"academ_stats/internal/helper/parser"
	"academ_stats/internal/helper/test"
	"testing"
)

// TestPersonalGraphql checks parsing all fields in struct Personal
func TestPersonalGraphql(t *testing.T) {
	src, err := test.ReadFile[request.TopCadets]("../test/test_files/src/top_cadets_graphql.json")
	if err != nil {
		t.Fatalf("read requst: %s", err)
	}

	if len(src.Cadets) == 0 {
		t.Fatal("no data to parse")
	}

	personal := parser.PersonalGraphql(&src.Cadets[0], test.Adm)

	// check fields
	if personal.ID != src.Cadets[0].UserID {
		t.Fatalf("id want \"%d\" but got \"%d\"", src.Cadets[0].UserID, personal.ID)
	} else if personal.Login != src.Cadets[0].Login {
		t.Fatalf("login want \"%s\" but got \"%s\"", src.Cadets[0].Login, personal.Login)
	} else if personal.FullName != src.Cadets[0].User.Attrs.FirstName+" "+src.Cadets[0].User.Attrs.LastName {
		t.Fatalf("full name want \"%s\" but got \"%s\"", src.Cadets[0].User.Attrs.FirstName+" "+src.Cadets[0].User.Attrs.LastName, personal.FullName)
	} else if personal.Gender != src.Cadets[0].User.Attrs.Gender {
		t.Fatalf("gender want \"%s\" but got \"%s\"", src.Cadets[0].User.Attrs.Gender, personal.Gender)
	} else if personal.Level != src.Cadets[0].Level {
		t.Fatalf("gender want \"%d\" but got \"%d\"", src.Cadets[0].Level, personal.Level)
	}

	if personal.Age <= 16 || personal.Age >= 100 {
		t.Logf("check age %s", src.Cadets[0].User.Attrs.DateOfBirth)
	}
}
