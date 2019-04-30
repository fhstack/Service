#include <stdio.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/stat.h>
#include "utils.h"

int CheckFileIsExist(const char* path) {
    FILE* f = fopen(path, "r");
    if(f == NULL) {
        return 0;
    } else {
        fclose(f);
        return 1;
    }
}

int GetFileSize(char* path) {
    struct stat f_stat;
    if(stat(path, &f_stat) == -1) {
        return 0;
    }
    return f_stat.st_size;
}

//MB
int GetMemory(int pid) {
    char buf[100] = {0};
    sprintf(buf, "sh /home/fh/goWorkspace/src/github.com/l-f-h/judge/inline/getmemory.sh %d", pid);
    FILE* pipstream = popen(buf, "r");
    int memory;
    fscanf(pipstream, "%d", &memory);
    return (memory << 10);
}
