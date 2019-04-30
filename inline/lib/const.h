#pragma once

#define CallArraySize 512
#define ONEMB 1048576LL
#define F_LIM (ONEMB << 5)
#define M_LIM (ONEMB << 7)
#define S_LIM (ONEMB << 7)
#define C_LIM 3.0
#define DEBUG 0

char* UserExecutableFilePath;
char* InDataFilePath;
char* OutDataFilePath;
char* UserExecErrorFilePath;

//errorcode
static const int NormalExited = 0;
static const int FileIsNotExist = 1;
static const int ForkFailed = 2;
static const int PtraceFailed = 3;
static const int ExecFailed = 4;
static const int PtraceSetOptionsFailed = 5;
static const int RunTimeError = 6;
static const int TimeLimitError = 7;
static const int OutLimitError = 8;
static const int PipFailedError = 9;
static const int IllegalSyscallError = 10;
static char nothing[] = "";
static char FileIsNotExistMessage[] = "file is not exist";
static char ForkFailedMessage[] = "fork failed!";
static char PtraceFailedMessage[] = "ptrace failed!";
static char ExecFailedMessage[] = "exec failed";
static char PtraceSetOptionsFailedMessage[] = "ptrace set options failed";
static char RunTimeErrorMessage[] = "runtime error!";
static char TimeLimitErrorMessage[] = "time limit!";
static char OutLimitErrorMessage[] = "outut limit!";
static char PipFailedErrorMessage[] = "pipe open failed";
static char IllegalSyscallErrorMessage[] = "illegal syscall!";
static char* ErrorMessage[] = {nothing, FileIsNotExistMessage, ForkFailedMessage, PtraceFailedMessage, ExecFailedMessage, PtraceSetOptionsFailedMessage, RunTimeErrorMessage, TimeLimitErrorMessage, OutLimitErrorMessage, PipFailedErrorMessage, IllegalSyscallErrorMessage};
//syscall limits
