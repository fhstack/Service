package compiler

import (
	"github.com/l-f-h/Service/utils"
	"github.com/l-f-h/judge/conf"
	"log"
	"os"
	"time"
)

type Compiler struct {
	SubmitId    string
	ProblemId   string
	Language    string
	UserCode    string
	CompileInfo string
}

//compile args
//TODO add flags could limit some user illegal behavior
var LanguageToCompileArgs = map[string]string{
	"c":   "gcc -o",
	"cpp": "g++ -o",
}

func (compiler *Compiler) createSourceCodeFileAndOutDataFile(CompleteSourceCodeName string) error {
	//user sourcecode and executable program dir
	SubmitWorkPath := utils.GetSubmitWorkPath(compiler.SubmitId)
	err := os.MkdirAll(SubmitWorkPath, os.FileMode(0755))
	if err != nil {
		log.Printf("create SubmitWorkPath failed: %v", err)
		return err
	}
	err = os.Chdir(SubmitWorkPath)
	if err != nil {
		log.Printf("chdir failed: %v", err)
		return err
	}
	code, err := os.Create(CompleteSourceCodeName)
	defer code.Close()
	if err != nil {
		log.Printf("create SourceCode failed: %v", err)
		return err
	}
	_, err = code.Write([]byte(compiler.UserCode))
	if err != nil {
		log.Printf("write SourceCode failed: %v", err)
		return err
	}
	outdata, err := os.Create("out.data")
	if err != nil {
		log.Printf("create out.data failed: %v", err)
		return err
	}
	err = outdata.Chmod(os.FileMode(0666))
	if err != nil {
		log.Printf("chmod out.data failed: %v", err)
		return err
	}
	return nil
}

func (compiler *Compiler) Compile() (successful bool) {
	//编译型语言
	CompileArg := LanguageToCompileArgs[compiler.Language]
	CompleteSourceCodeName := utils.GetCompleteSourceCodeName(compiler.SubmitId, compiler.Language)
	CompleteCompileArg := CompileArg + " " + conf.GeneratedProgramName + " " + CompleteSourceCodeName
	err := compiler.createSourceCodeFileAndOutDataFile(CompleteSourceCodeName)
	if err != nil {
		log.Printf("CreateSourceCodeFile happend error: %v", err)
		return false
	}
	if compiler.Language == "c" || compiler.Language == "cpp" || compiler.Language == "java" {
		ch := make(chan bool, 1)
		var stderr string
		go func() {
			_, stderr, err = utils.Exec(CompleteCompileArg)
			ch <- true
			close(ch)
		}()
		//allowed compile time max =  10 s
		timer := time.NewTimer(time.Second * 10)
		defer timer.Stop()
		select {
		case <-ch:
			if err != nil {
				log.Printf("compile submit submitId=%s failed error: %v", compiler.SubmitId, err)
				successful = false
			} else {
				log.Printf("compile submit submitId=%s successful", compiler.SubmitId)
				successful = true
			}
			compiler.CompileInfo = stderr
			return successful
		case <-timer.C: //compile out of time limit
			compiler.CompileInfo = "compile out of time limit"
			log.Printf("detect the attack behavior, the submit_id = %s", compiler.SubmitId)
			return true
		}
	} else {
		//chmod: python
		return true
	}
}
