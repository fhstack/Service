package comparer

import (
	"fmt"
	"github.com/l-f-h/Service/conf"
	"github.com/l-f-h/Service/utils"
	"log"
)

func (comparer *Comparer) Diff() bool {
	OutDataFilePath := utils.GetOutDataFilePath(comparer.SubmitId)
	StdDataFilePath := utils.GetStdDataFilePath(comparer.ProblemId)
	command := fmt.Sprintf("%s %s %s", conf.ComparePath, OutDataFilePath, StdDataFilePath)
	stdout, stderr, err := utils.Exec(command)
	if err != nil {
		log.Printf("Exec Compare failed!")
		return false
	}
	if stderr != "" {
		log.Fatalf("Compare: %s", stderr)
		return false
	}
	var result int32
	_, err = fmt.Sscan(stdout, &result)
	if err != nil {
		log.Fatalf("Sscan stdout Compare result failde: %v", err)
		return false
	}
	log.Printf("Compare result: %d", result)
	if result == 0 {
		return false
	} else {
		return true
	}
}
