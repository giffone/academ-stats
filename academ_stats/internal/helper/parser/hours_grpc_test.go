package parser_test

import (
	"academ_stats/internal/domain"
	"academ_stats/internal/helper/parser"
	"academ_stats/internal/helper/test"
	"academ_stats/internal/repository/pb/session_manager"
	"testing"
)

func TestHoursGrpc(t *testing.T) {
	if res := parser.HoursGrpc(nil); res != nil {
		t.Fatal("must return nil response on nil request")
	}

	src, err := test.ReadFile[session_manager.CadetsTimeResponse]("../test/test_files/src/hours_grpc.json")
	if err != nil {
		t.Fatalf("read hours: req: %s", err)
	}

	got := parser.HoursGrpc(src)

	want, err := test.ReadFile[map[int]domain.HoursDTO]("../test/test_files/want/hours_grpc.json")
	if err != nil {
		t.Fatalf("read hours: res: %s", err)
	}

	if len(got) != len(*want) {
		t.Fatal("length want/got data not equal")
	}

	for userID, wantUser := range *want {
		gotUser, ok := got[userID]
		if !ok {
			t.Fatalf("cant not find user \"%d\"", userID)
		}

		if len(wantUser.Month) != len(gotUser.Month) {
			t.Fatal("length want/got months not equal")
		}

		for j := 0; j < len(wantUser.Month); j++ {
			if wantUser.Month[j].Year != gotUser.Month[j].Year {
				t.Fatalf("user id \"%d\" year want \"%s\" but got \"%s\"", userID, wantUser.Month[j].Year, gotUser.Month[j].Year)
			} else if wantUser.Month[j].Month != gotUser.Month[j].Month {
				t.Fatalf("user id \"%d\" month want \"%s\" but got \"%s\"", userID, wantUser.Month[j].Month, gotUser.Month[j].Month)
			} else if wantUser.Month[j].Hours != gotUser.Month[j].Hours {
				t.Fatalf("user id \"%d\" month want \"%d\" but got \"%d\"", userID, wantUser.Month[j].Hours, gotUser.Month[j].Hours)
			}
		}
	}
}
