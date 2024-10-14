package top_cadets_test

import (
	"excel_table/internal/domain"
	"excel_table/internal/helper/test"
	"testing"
)

func TestMatrix(t *testing.T) {
	got, err := test.Build("../../helper/test/test_files/src/top_cadets_grpc.json", "../../helper/test/test_files/src/table_queue.json")
	if err != nil {
		t.Fatal(err)
	}
	if got == nil {
		t.Fatal("getting data is nil")
	}

	if test.Period == "" {
		t.Fatal("check the period")
	}

	wantHead := buildHead()
	if len(wantHead) == 0 {
		t.Fatal("head: not created head")
	}

	if len(wantHead) != len(got.Head) {
		t.Fatalf("head: length want \"%d\" but got \"%d\"", len(wantHead), len(got.Head))
	}

	// check head
	for i := 0; i < len(wantHead); i++ {
		for j := 0; j < len(wantHead[i]); j++ {
			if len(wantHead[i]) != len(got.Head[i]) {
				t.Fatalf("head: length on index [%d] want \"%d\" but got \"%d\"", i, len(wantHead[i]), len(got.Head[i]))
			}
			if wantHead[i][j].ID != got.Head[i][j].ID {
				t.Fatalf("head: id on index [%d][%d] want \"%d\" but got \"%d\"", i, j, wantHead[i][j].ID, got.Head[i][j].ID)
			}
			if wantHead[i][j].Title != got.Head[i][j].Title {
				t.Fatalf("head: title on index [%d][%d] want \"%s\" but got \"%s\"", i, j, wantHead[i][j].Title, got.Head[i][j].Title)
			}
			if wantHead[i][j].RaiseX != got.Head[i][j].RaiseX {
				t.Fatalf("head: raiseX on index [%d][%d] want \"%d\" but got \"%d\"", i, j, wantHead[i][j].RaiseX, got.Head[i][j].RaiseX)
			}
		}
	}

	// check data
	wantData := buildData()
	if len(wantData) == 0 {
		t.Fatal("data: not created data")
	}

	if len(wantData) != len(got.Data) {
		t.Fatalf("data: length want \"%d\" but got \"%d\"", len(wantHead), len(got.Head))
	}

	for i := 0; i < len(wantData); i++ {
		for j := 0; j < len(wantData[i]); j++ {
			if j >= len(got.Data[i]) {
				break
			}
			if wantData[i][j].ID != got.Data[i][j].ID {
				if j == 8 && got.Data[i][j].ID == 0 {
					continue // ignore users who do not have hours
				}
				t.Fatalf("data: id on index [%d][%d] want \"%d\" but got \"%d\"", i, j, wantData[i][j].ID, got.Data[i][j].ID)
			}
		}
	}
}

func buildHead() [][]domain.Head {
	return [][]domain.Head{
		{
			// head_1
			{ID: domain.IndexPersonal, Title: domain.NamePersonal, RaiseX: 7},
			{ID: domain.IndexHours, Title: domain.NameHours, RaiseX: 0},
			{ID: domain.IndexModule, Title: domain.NameModule, RaiseX: 25},
			{ID: domain.IndexCheckpoint, Title: domain.NameCheckpoint, RaiseX: 6},
		},
		{
			// head_2
			{RaiseX: 7},
			{RaiseX: 0},
			{ID: 0, Title: "Go", RaiseX: 22},
			{ID: 3, Title: "C"},
			{ID: 1000, Title: domain.NameLangOther, RaiseX: 1},
			{ID: 0, Title: "Go", RaiseX: 6},
		},
		{
			// head_3
			{RaiseX: 7},
			{RaiseX: 0},
			{ID: 5000, Title: "5000 xp", RaiseX: 1},
			{ID: 6125, Title: "6125 xp", RaiseX: 5},
			{ID: 9200, Title: "9200 xp", RaiseX: 3},
			{ID: 12250, Title: "12250 xp", RaiseX: 5},
			{ID: 18375, Title: "18375 xp", RaiseX: 0},
			{ID: 19100, Title: "19100 xp", RaiseX: 0},
			{ID: 24500, Title: "24500 xp", RaiseX: 0},
			{ID: 34375, Title: "34375 xp", RaiseX: 0},
			{ID: 76250, Title: "76250 xp", RaiseX: 0},
			{ID: 75250, Title: "75250 xp", RaiseX: 0},
			{ID: 5000, Title: "5000 xp", RaiseX: 0},
			{ID: 10000, Title: "10000 xp", RaiseX: 0},
			{ID: 76, Title: "#76", RaiseX: 0},
			{ID: 77, Title: "#77", RaiseX: 0},
			{ID: 78, Title: "#78", RaiseX: 0},
			{ID: 79, Title: "#79", RaiseX: 0},
			{ID: 80, Title: "#80", RaiseX: 0},
			{ID: 81, Title: "#81", RaiseX: 0},
			{ID: 82, Title: "#82", RaiseX: 0},
		},
		{
			// head_4
			{ID: domain.IndexID, Title: domain.NameID},
			{ID: domain.IndexLogin, Title: domain.NameLogin},
			{ID: domain.IndexFullName, Title: domain.NameFullName},
			{ID: domain.IndexAge, Title: domain.NameAge},
			{ID: domain.IndexGender, Title: domain.NameGender},
			{ID: domain.IndexAdmission, Title: domain.NameAdmission},
			{ID: domain.IndexLevel, Title: domain.NameLevel},
			{ID: domain.IndexXPTotal, Title: domain.NameXPTotal},
			{ID: domain.IndexHours, Title: test.Period},
			{ID: 101759, Title: "go-reloaded"},
			{ID: 101956, Title: "tetris-optimizer"},
			{ID: 101753, Title: "ascii-art"},
			{ID: 101767, Title: "ascii-art-fs"},
			{ID: 101777, Title: "ascii-art-output"},
			{ID: 101778, Title: "ascii-art-justify"},
			{ID: 101779, Title: "ascii-art-reverse"},
			{ID: 101808, Title: "ascii-art-color"},
			{ID: 101755, Title: "ascii-art-web"},
			{ID: 101756, Title: "ascii-art-web-stylize"},
			{ID: 101780, Title: "ascii-art-web-dockerize"},
			{ID: 101781, Title: "ascii-art-web-export-file"},
			{ID: 101757, Title: "my-ls-1"},
			{ID: 101758, Title: "net-cat"},
			{ID: 101782, Title: "groupie-tracker-filters"},
			{ID: 101783, Title: "groupie-tracker-geolocalization"},
			{ID: 101784, Title: "groupie-tracker-visualizations"},
			{ID: 101785, Title: "groupie-tracker-search-bar"},
			{ID: 101754, Title: "push-swap"},
			{ID: 101948, Title: "forum-authentication"},
			{ID: 101809, Title: "groupie-tracker"},
			{ID: 101810, Title: "lem-in"},
			{ID: 101941, Title: "forum"},
			{ID: 101955, Title: "atm-management-system"},
			{ID: 101943, Title: "guess-it-1"},
			{ID: 101942, Title: "math-skills"},
			{ID: 76, Title: "8 Feb"},
			{ID: 77, Title: "15 Feb"},
			{ID: 78, Title: "29 Feb"},
			{ID: 79, Title: "7 Mar"},
			{ID: 80, Title: "14 Mar"},
			{ID: 81, Title: "28 Mar"},
			{ID: 82, Title: "4 Apr"},
		},
	}
}

func buildData() [][]domain.Cell {
	c := [][]domain.Cell{}
	i := 8
	userData := []domain.Cell{
		{ID: domain.IndexID},
		{ID: domain.IndexLogin},
		{ID: domain.IndexFullName},
		{ID: domain.IndexAge},
		{ID: domain.IndexGender},
		{ID: domain.IndexAdmission},
		{ID: domain.IndexLevel},
		{ID: domain.IndexXPTotal},
		{ID: domain.IndexHours},
	}

	for i > 0 {
		c = append(c, userData)
		i--
	}

	return c
}
