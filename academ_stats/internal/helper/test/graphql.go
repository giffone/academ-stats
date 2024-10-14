package test

import (
	"academ_stats/internal/domain"
	"academ_stats/internal/domain/request"
	"academ_stats/internal/helper/parser"
	"academ_stats/internal/repository/pb/session_manager"

	"encoding/json"
	"fmt"
)

func CheckTopCadetsGraphql(tc_req_dir, tc_res_dir, hours_dir string) (*domain.TopCadets, error) {
	srcTC, err := ReadFile[request.TopCadets](tc_req_dir)
	if err != nil {
		return nil, fmt.Errorf("read requst: %s", err)
	}

	srcHours, err := ReadFile[session_manager.CadetsTimeResponse](hours_dir)
	if err != nil {
		return nil, fmt.Errorf("read hours: %s", err)
	}

	tcGot := parser.TopCadetsGraphql(srcTC, Adm, parser.HoursGrpc(srcHours), &Period)

	tcWant, err := ReadFile[domain.TopCadets](tc_res_dir)
	if err != nil {
		return nil, fmt.Errorf("read want response: %s", err)
	}

	got, _ := json.Marshal(tcGot)
	want, _ := json.Marshal(tcWant)

	fmt.Println(string(got))

	if string(got) != string(want) {
		return nil, fmt.Errorf("json is not equal")
	}

	return tcGot, nil
}
