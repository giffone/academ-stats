package excel_test

import (
	"excel_table/internal/domain"
	"excel_table/internal/helper/test"
	"excel_table/internal/repository/excel"
	"fmt"
	"log"
	"testing"
)

func TestTopCadetsExcelFile(t *testing.T) {
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

	file, err := excel.ExcelFile([]*domain.Sheet{got})
	if err != nil {
		t.Fatalf("excel file build: %v", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Println(err)
		}
	}()

	rows, err := file.Rows(got.Name)
	if err != nil {
		t.Fatalf("excel file row iterator: %v", err)
	}
	y := 0
	for rows.Next() {
		if y >= len(topCadetsData) {
			t.Fatalf("rows want %d but got more", len(topCadetsData))
		}
		row, err := rows.Columns()
		if err != nil {
			t.Fatalf("excel file col iterator: next: %v", err)
		}
		for x, colCell := range row {
			if x >= len(topCadetsData[y]) {
				t.Fatalf("row length want %d but got more", len(topCadetsData[y]))
			}
			if topCadetsData[y][x] != colCell {
				t.Fatalf("on cell[%d][%d] want %s but got %s", y, x, topCadetsData[y][x], colCell)
			}
		}
		y++
	}
	if err = rows.Close(); err != nil {
		fmt.Println(err)
	}

	// cols, err := file.Cols(got.Name)
	// if err != nil {
	// 	t.Fatalf("excel file col iterator: %v", err)
	// }

	// for cols.Next() {
	// 	col, err := cols.Rows()
	// 	if err != nil {
	// 		t.Fatalf("excel file col iterator: next: %v", err)
	// 	}
	// 	for _, rowCell := range col {
	// 		fmt.Print(rowCell, "\t")
	// 	}
	// 	fmt.Println()
	// }

	// Save spreadsheet by the given path.
	// tables, err := file.GetTables(got.Name)
	// if err != nil {
	// 	t.Fatalf("excel file get tables: %v", err)
	// }

	// if err := file.SaveAs("Book1.xlsx"); err != nil {
	// 	t.Fatal(err)
	// }
}

var topCadetsData = [][]string{
	{"Personal", "", "", "", "", "", "", "", "Hours", "Module", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "Checkpoint"},
	{"", "", "", "", "", "", "", "", "", "Go", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "C", "other", "", "Go"},
	{"", "", "", "", "", "", "", "", "", "5000 xp", "", "6125 xp", "", "", "", "", "", "9200 xp", "", "", "", "12250 xp", "", "", "", "", "", "18375 xp", "19100 xp", "24500 xp", "34375 xp", "76250 xp", "75250 xp", "5000 xp", "10000 xp", "#76", "#77", "#78", "#79", "#80", "#81", "#82"},
	{"ID", "Login", "Full name", "Age", "Gender", "Admission", "Level", "XP total", "Jan24 - May24", "go-reloaded", "tetris-optimizer", "ascii-art", "ascii-art-fs", "ascii-art-output", "ascii-art-justify", "ascii-art-reverse", "ascii-art-color", "ascii-art-web", "ascii-art-web-stylize", "ascii-art-web-dockerize", "ascii-art-web-export-file", "my-ls-1", "net-cat", "groupie-tracker-filters", "groupie-tracker-geolocalization", "groupie-tracker-visualizations", "groupie-tracker-search-bar", "push-swap", "forum-authentication", "groupie-tracker", "lem-in", "forum", "atm-management-system", "guess-it-1", "math-skills", "8 Feb", "15 Feb", "29 Feb", "7 Mar", "14 Mar", "28 Mar", "4 Apr"},
	{"870", "aabdikha", "Aibek Abdikhalyk", "23", "Male", "piscine23sep", "21", "421,750", "100", "done", "done", "done", "done", "done", "done", "done", "done", "done", "done", "done", "done", "", "done", "done", "done", "done", "done", "done", "done", "done", "done", "done", "done", "done", "done", "7", "1", "7", "", "7", "7", "7"},
	{"634", "aolzhass", "Akhmediyar  Olzhassov", "20", "Male", "piscine23sep", "21", "376,900", "200", "done", "done", "done", "done", "done", "done", "done", "done", "done", "done", "done", "done", "", "done", "", "done", "done", "done", "done", "", "done", "done", "done", "done", "", "done", "7", "7"},
	{"1,254", "famanbaye", "Farhad AMANBAYEV", "34", "Male", "piscine23nov", "17", "223,725", "", "done", "done", "done", "done", "done", "done", "done", "done", "done", "done", "done", "done", "done", "done", "done", "done", "done", "done", "done", "", "done", "", "", "", "", "done", "7", "7", "7", "7", "7"},
	{"25", "Akhmet_bayan", "Bayan Akhmet", "23", "Male", "piscine23sep", "16", "185,100", "", "done", "done", "done", "done", "done", "done", "done", "done", "done", "done", "done", "done", "", "done", "done", "done", "done", "done", "", "", "done", "", "", "", "", "done", "4", "2", "4", "7", "2", "3"},
	{"1,543", "aabdrama", "Alikhan Abdraman", "19", "Male", "piscine24jan", "0", "200", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "1", "1"},
	{"1,181", "tyergazy", "Temirlan Yergazyuly", "23", "Male", "piscine23nov / piscine24jan", "0", "200", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "1", "1"},
	{"341", "aabdykal", "Alnur Abdykalykov", "21", "Male", "piscine23sep", "0", "0"},
	{"600", "ukabdoll", "Uristem  Kabdolla", "22", "Male", "piscine23sep", "0", "0"},
}
