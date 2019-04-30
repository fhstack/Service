package utils

import (
	"fmt"
	"github.com/l-f-h/common"

	"log"
)

func GetMemoryAndTime(stdout string) (memory int32, time float32, err error) {
	_, err = fmt.Sscanln(stdout, &memory, &time)
	if err != nil {
		log.Fatalf("Scanln memory and time failed: %v", err)
		return 0, 0, err
	}
	return memory, time, nil
}

//from inline/lib/const.h
/*const int NormalExited = 0;
const int FileIsNotExist = 1;
const int ForkFailed = 2;
const int PtraceFailed = 3;
const int ExecFailed = 4;
const int PtraceSetOptionsFailed = 5;
const int RunTimeError = 6;
const int TimeLimitError = 7;
const int OutLimitError = 8;
const int PipFailedError = 9;
const int IllegalSyscallError = 10;*/

func GetError(stderr string) (res int32) {
	if stderr == "" {
		res = common.AC
		return
	}
	var ErrCode int
	_, err := fmt.Sscan(stderr, &ErrCode)
	if err != nil {
		log.Fatalf("Scanf ErrCode failed: %v", err)
		res = common.ServerInternalError
		return
	}
	if ErrCode == 6 || ErrCode == 10 {
		res = common.RE
	} else if ErrCode == 7 {
		res = common.TL
	} else if ErrCode == 8 {
		res = common.OL
	} else {
		res = common.ServerInternalError
	}
	return
}
