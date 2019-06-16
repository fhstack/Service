package method

import (
	"context"
	"fmt"
	"github.com/l-f-h/Service/comparer"
	"github.com/l-f-h/Service/compiler"
	"github.com/l-f-h/Service/conf"
	"github.com/l-f-h/Service/generated"
	"github.com/l-f-h/Service/utils"
	"github.com/l-f-h/common"
	"log"
)

func (server *Server) Judge(c context.Context, req *pb_gen.JudgeRequest) (*pb_gen.JudgeResponse, error) {
	//construct a compiler
	cmpl := ctorOneCompiler(req)
	resp := &pb_gen.JudgeResponse{}
	callBackUrl := req.GetCallBackUrl()
	if res := cmpl.Compile(); res == false {
		resp.ResCode = common.CE
		resp.CompileInfo = cmpl.CompileInfo
		log.Printf("compile self happened an error, submitId = %s", req.SubmitId)
		return resp, nil
	}
	//construct
	UserExecutableFilePath := utils.GetUserExecutableFilePath(req.SubmitId)
	InDataFilePath := utils.GetInDataFilePath(req.ProblemId)
	OutDataFilePath := utils.GetOutDataFilePath(req.SubmitId)
	command := fmt.Sprintf("%s %s %s %s", conf.RunPath, UserExecutableFilePath, InDataFilePath, OutDataFilePath)
	log.Println(command)
	//err is useless
	stdout, stderr, _ := utils.Exec(command)
	log.Println("Exec UserExecutableFile program is ok")
	log.Printf("***********stderr**********\n%s", stderr)
	log.Printf("***********stdout**********\n%s", stdout)
	ErrType := utils.GetError(stderr)
	if ErrType != common.AC {
		switch callBackUrl == "" {
		case true:
			resp.ResCode = ErrType
			return resp, nil
		case false:
			utils.Callback(callBackUrl, resp)
			return resp, nil
		}
	}
	log.Println("GetError from stderr is ok")
	if stdout == "" {
		//shouldn't happen this, just for safe
		switch callBackUrl == "" {
		case true:
			resp.ResCode = common.ServerInternalError
			return resp, nil
		case false:
			utils.Callback(callBackUrl, resp)
			return resp, nil
		}
	}
	memory, time, err := utils.GetMemoryAndTime(stdout)

	if err != nil {
		log.Printf("get run result failed: %v", err)
		switch req.CallBackUrl == "" {
		case true:
			resp.ResCode = common.ServerInternalError
			return resp, nil
		case false:
			utils.Callback(callBackUrl, resp)
			return resp, nil
		}
		return resp, nil
	}
	resp.Memory = float32(memory >> 20) //MB
	resp.Time = time                    //ms
	log.Println("GetTimeAndMemoryByProblemId is ok")
	if resp.Memory > req.MemoryLimit {
		switch req.CallBackUrl == "" {
		case true:
			resp.ResCode = common.ML
			return resp, nil
		case false:
			utils.Callback(callBackUrl, resp)
			return resp, nil
		}
	}
	if resp.Time > req.TimeLimit {
		switch req.CallBackUrl == "" {
		case true:
			resp.ResCode = common.TL
			return resp, nil
		case false:
			//callback
			utils.Callback(callBackUrl, resp)
			return resp, nil
		}
	}
	cmper := &comparer.Comparer{
		SubmitId:  req.SubmitId,
		ProblemId: req.ProblemId,
	}
	result := cmper.Diff()
	//TODO support AC ration
	if result == true {
		resp.ResCode = common.AC
	} else {
		resp.ResCode = common.EA
	}
	switch req.CallBackUrl == "" {
	case true:
		return resp, nil
	case false:
		utils.Callback(callBackUrl, resp)
		return resp, nil
	}
	return resp, nil
}

func ctorOneCompiler(req *pb_gen.JudgeRequest) *compiler.Compiler {
	Compiler := compiler.Compiler{}
	Compiler.SubmitId = req.SubmitId
	Compiler.ProblemId = req.ProblemId
	Compiler.Language = req.Language
	Compiler.UserCode = req.SourceCode
	log.Printf("construct one compiler %v", Compiler)
	return &Compiler
}
