package utils

import (
	"github.com/l-f-h/Service/conf"
)

func GetUserExecutableFilePath(submitId string) string {
	UserExecutableFilePath := conf.UserSubmitHome + conf.SubmitPrefix + submitId + "/" + conf.GeneratedProgramName
	return UserExecutableFilePath
}

func GetInDataFilePath(problemId string) string {
	InDataFilePath := conf.ProblemDataHome + problemId + "/" + "in.data"
	return InDataFilePath
}

func GetOutDataFilePath(submitId string) string {
	OutDataFilePath := conf.UserSubmitHome + conf.SubmitPrefix + submitId + "/" + "out.data"
	return OutDataFilePath
}

func GetStdDataFilePath(problemId string) string {
	StdDataFilePath := conf.ProblemDataHome + problemId + "/" + "std.data"
	return StdDataFilePath
}

func GetSubmitWorkPath(submitId string) string {
	SubmitWorkPath := conf.UserSubmitHome + "/" + conf.SubmitPrefix + submitId
	return SubmitWorkPath
}

func GetCompleteSourceCodeName(submitId string, language string) string {
	CompleteSourceCodeName := conf.SourceCodeName + "_" + submitId + "." + language
	return CompleteSourceCodeName
}
