#include <unistd.h>
#include <stdio.h>
#include "lib/okcalls.h"
#include "lib/const.h"
#include "lib/ncalls.h"
#include "lib/utils.h"
#include "lib/parent.h"
#include "lib/child.h"

void Begin() {
    //pip transform nothing just synchronize
    int pip[2];
    if(pipe(pip) < 0) {
        fprintf(stderr, "%d\n%s", PipFailedError, ErrorMessage[PipFailedError]);
        return;
    }
    pid_t pid = fork();
    if(pid < 0) {
        fprintf(stderr, "%d\n%s", ForkFailed, ErrorMessage[ForkFailed]);
    } else if(pid == 0) {
        close(pip[0]);
        Child(pip[1]);
    } else {
        //parent
        char buf[10] = {0};
        close(pip[1]);
        InitSyscallLimits();
        read(pip[0], buf, 10);
        if(DEBUG)
            printf("child is already!\n");
        WatchChild(pid);
   }
}

int main(int argc, char* argv[]) {
    UserExecutableFilePath = argv[1];
    InDataFilePath = argv[2];
    OutDataFilePath = argv[3];
    //UserExecErrorFilePath = argv[4];
    //check
    for(int i = 1;i <= 2;i++) {
        if(CheckFileIsExist(argv[i]) == 0) {
            fprintf(stderr, "%d\n%s: %s\n",FileIsNotExist, ErrorMessage[FileIsNotExist], argv[i]);
            return 0;
        }
    }
    if(DEBUG)
        printf("####File all exist\n%s\n%s\n%s\n", InDataFilePath, OutDataFilePath, UserExecutableFilePath);
    Begin();
    return 0;
}
