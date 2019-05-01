#include <unistd.h>
#include <stdlib.h>
#include <string.h>
#include <memory.h>
#include <stdio.h>
#include <sys/reg.h>
#include <sys/user.h>
#include <sys/types.h>
#include <sys/wait.h>
#include <sys/resource.h>
#include <sys/ptrace.h>
#include "const.h"
#include "okcalls.h"
#include "ncalls.h"
#include "utils.h"
#include "parent.h"

int CallCounter[CallArraySize];

void InitSyscallLimits() {
    int i;
    memset(CallCounter, 0, CallArraySize);
    //now c and cpp
    for(i = 0;i == 0 || LANG_CV[i]; i++) {
        CallCounter[LANG_CV[i]] = ncalls[LANG_CV[i]];
    }
    
    if(DEBUG)
        fprintf(stderr,"InitSyscallLimits is successful\n");
}

void WatchChild(pid_t c_pid) {
    float UsedTime = 0;
    int Memory = 0;
    int status, sig;
    struct user_regs_struct reg;
    struct rusage resource_use;
    int first_set = 0;
	//get the useage
	//check the exitcode
	//check the if exit by signal
	//check the syscall
	while(1) {
		/*The  wait3() and wait4() system calls are similar to waitpid(2), but
        additionally return resource usage information about  the  child  in
        the structure pointed to by rusage.*/
        /*__WALL wait for all children, regardless of type ("clone" or "non-clone").*/
        if(wait4(c_pid, &status, __WALL, &resource_use) == -1) {
            fprintf(stderr, "%d\n%s", RunTimeError, ErrorMessage[RunTimeError]);
            return;
        }
        if(first_set) {
            if(ptrace(PTRACE_SETOPTIONS, c_pid, NULL, PTRACE_O_TRACEEXIT | PTRACE_O_TRACESYSGOOD) != 0) {
                fprintf(stderr, "%d\n%s", PtraceSetOptionsFailed, ErrorMessage[PtraceSetOptionsFailed]);
                return;
            }/* can't find PTRACE_O_EXITKILL? */
            first_set = 0;
        }
        
        int curMemory = GetMemory(c_pid);
        if(curMemory > Memory) {
            Memory = curMemory;    
        }
        
        if(WIFEXITED(status)) {
            if(WEXITSTATUS(status) != 0) {
                fprintf(stderr, "%d\n%s", RunTimeError, ErrorMessage[RunTimeError]);
                return;
            }
            if(DEBUG)
                printf("user program normal exit\n");
            break;
        }
        //check some STOPPED signal
        if(WIFSTOPPED(status)) {
            sig = WSTOPSIG(status);
            //if(DEBUG)
            //    printf("child proc is stopped by: %d!\n", sig);
            switch(sig) {
            case SIGTRAP:break; //normal
            case SIGFPE:        //div 0
                fprintf(stderr, "%d\n%s div zero", RunTimeError, ErrorMessage[RunTimeError]);
                return;
            case SIGALRM:
                fprintf(stderr, "%d\n%s", TimeLimitError, ErrorMessage[TimeLimitError]);
                return;
            case SIGXCPU:
                fprintf(stderr, "%d\n%s", TimeLimitError, ErrorMessage[TimeLimitError]);
                return;
            case SIGXFSZ:
                fprintf(stderr, "%d\n%s", OutLimitError, ErrorMessage[OutLimitError]);
                return;
            case SIGSEGV:
                fprintf(stderr, "%d\n%s segment fault!", RunTimeError, ErrorMessage[RunTimeError]);
            case SIGSYS:
                fprintf(stderr, "%d\n%s", IllegalSyscallError, ErrorMessage[IllegalSyscallError]);
                return;
            default:
                fprintf(stderr, "stopped by unknown signal: %d", sig);
                return;
            }
        }
        if(GetFileSize(UserExecErrorFilePath) != 0) {
            ptrace(PTRACE_KILL, c_pid, NULL, NULL);
            fprintf(stderr, "%d\n%s", RunTimeError, ErrorMessage[RunTimeError]);
            break;
        }
        //exit by signal
        if(WIFSIGNALED(status)) {
            sig = WTERMSIG(status);
            if(DEBUG)
                printf("user exit by signal: %d", sig);
            switch(sig) {
            case SIGCHLD:
            case SIGXCPU:
            case SIGKILL:
            case SIGALRM:
            case SIGSYS:
                fprintf(stderr, "%d\n%s", TimeLimitError, ErrorMessage[TimeLimitError]);
                return;
            case SIGXFSZ:
                fprintf(stderr, "%d\n%s", OutLimitError, ErrorMessage[OutLimitError]);
                return;
            default:
                fprintf(stderr, "%d\n%s", RunTimeError, ErrorMessage[RunTimeError]);
                return;
            }
        }
        //check syscalls
        ptrace(PTRACE_GETREGS, c_pid, NULL, &reg);
        long callId = reg.orig_rax;
        //illegal syscalls give runtime error    
        //if(DEBUG)
        //    printf("cur syscall: %ld\n", callId);
        if((CallCounter[callId]) <= 0) {
            ptrace(PTRACE_KILL, c_pid, NULL, NULL);
            fprintf(stderr,"%d\n%s%ld", RunTimeError, "user illegal syscall: ",callId);    //special ErrorMessage
            return;
        }
        ptrace(PTRACE_SYSCALL, c_pid, NULL, NULL);
    }
    //get time
    UsedTime += resource_use.ru_stime.tv_sec*1000 + resource_use.ru_stime.tv_usec / 1000.0;
    UsedTime += resource_use.ru_utime.tv_sec*1000 + resource_use.ru_utime.tv_usec / 1000.0;
    //user normal run
    fprintf(stdout,"%d %f", Memory, UsedTime);
    if(DEBUG)
        printf("over\n");
}
