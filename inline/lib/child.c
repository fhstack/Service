#include <unistd.h>
#include <sys/types.h>
#include <sys/ptrace.h>
#include <stdlib.h>
#include <time.h>
#include <sys/resource.h>
#include <stdio.h>
#include "child.h"
#include "const.h"
void SetResourceLimit() {
    struct rlimit lim;
    //CPU
    lim.rlim_cur = C_LIM;
    lim.rlim_max = C_LIM + 1;
    setrlimit(RLIMIT_CPU, &lim);
    alarm(0);
    alarm(C_LIM * 3);
    //file
    lim.rlim_cur = F_LIM;
    lim.rlim_max = F_LIM + ONEMB;
    setrlimit(RLIMIT_FSIZE, &lim);
    //proc for c&cpp one process is enough
    lim.rlim_cur = lim.rlim_max = 1;
    setrlimit(RLIMIT_NPROC, &lim);
    //stack
    lim.rlim_cur = lim.rlim_max = S_LIM;
    setrlimit(RLIMIT_STACK, &lim);
    //memory
    lim.rlim_cur = M_LIM;
    lim.rlim_max = M_LIM + ONEMB;
    
    if(DEBUG)
        fprintf(stderr,"####SetResourceLimit is successful!\n");
}

int SetChildRunEnv() {
    //child low priority
    nice(19);
    //reopne the child in and out
    //if(freopen(UserExecErrorFilePath, "a+", stderr) == NULL)    
    //   return 1;
    if(freopen(InDataFilePath, "r", stdin) == NULL)
        return 1;
    if(freopen(OutDataFilePath, "w", stdout) == NULL)
        return 1;
    if(DEBUG)
        fprintf(stderr, "####child call PTRACEME successful!\n");
    SetResourceLimit();
    alarm(5);
    return 0;
}


void Child(int rpip) {
    if(SetChildRunEnv() != 0) //TODO tell parent by pip
        return;
    if(DEBUG)
        fprintf(stderr,"####SetChildRunEnv is successful\nalready to execl\n");
    
    if(ptrace(PTRACE_TRACEME, 0, NULL, NULL) != 0) {
        return;
    }
    write(rpip, "ok", 3);
    usleep(10000); //gurantee parent is already to watch 10ms
    if(execl(UserExecutableFilePath, UserExecutableFilePath,(char*) NULL) == -1) {
        fprintf(stderr, "%d\n%s", ExecFailed, ErrorMessage[ExecFailed]);
        return;
    }
    fflush(stderr);
    exit(0);
}
