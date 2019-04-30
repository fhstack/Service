#include <stdio.h>
#include <ctype.h>
#include <stdlib.h>
#include "lib/utils.h"
//Nonspace and '\n'
void FindNextNonspace(int* c1, int* c2, FILE* f1, FILE* f2) {
    while(isspace(*c1) || isspace(*c2)) {
        if(*c1 != *c2) {
            if(*c1 == EOF) {
                do {
                    *c2= fgetc(f2);
                } while(isspace(*c1));
                continue;
            } else if (*c2 == EOF) {
                do {
                    *c1 = fgetc(f1);
                } while(isspace(*c1));
                continue;
            } else if(isspace(*c1) && isspace(*c2)) {   //ignore the end of line space
                while(*c2 == '\n' && isspace(*c1) && *c1 != '\n')
                    *c1 = fgetc(f1);
                while(*c1 == '\n' && isspace(*c2) && *c2 != '\n')
                    *c2 = fgetc(f2);
            } else {
                fprintf(stdout, "%d", 0);
                exit(0);
            }
        }
        if(isspace(*c1))
            *c1 = fgetc(f1);
        if(isspace(*c2))
            *c2 = fgetc(f2);
    }
}

void Compare(const char* file1, const char* file2) {
    FILE* f1, *f2;
    f1 = fopen(file1, "re");
    f2 = fopen(file2, "re");
    int c1, c2;
    int isAC = 1;
    //must exist
    for(;;) {
        c1 = fgetc(f1);
        c2 = fgetc(f2);
        FindNextNonspace(&c1, &c2, f1, f2);
        for(;;) {
            while((!isspace(c1) && c1) || (!isspace(c2) && c2)) {
                if(c1 == EOF && c2 == EOF) {
                    goto end;
                }
                //only one EOF, continue
                if(c1 == EOF || c2 == EOF) {
                    break;
                }
                if(c1 != c2) {
                    isAC = 0;
                    goto end;
                }
                c1 = fgetc(f1);
                c2 = fgetc(f2);
            }
            FindNextNonspace(&c1, &c2, f1, f2);
            if(c1 == EOF && c2 == EOF) {
                goto end;
            }
            if(c1 == EOF || c2 == EOF) {
                isAC = 0;
                goto end;
            }
            if((c1 == '\n' || !c1 ) && (c2 == '\n' || !c2)) {
                break;
            }
        }
    }
end:
    fprintf(stdout, "%d", isAC);
    fclose(f1);
    fclose(f2);
}

int main(int argc, const char* argv[]) {
    if(argc != 3) {
        fprintf(stderr, "parmas error");
        return 0;
    }
    const char* OutDataFilePath = argv[1];
    const char* StdDataFilePath = argv[2];
    if(CheckFileIsExist(OutDataFilePath) == 0) {
        fprintf(stderr, "OutDataFile is not exist!");
        return 0;
    }
    if(CheckFileIsExist(StdDataFilePath) == 0) {
        fprintf(stderr, "StdDataFile is not exist!");
        return 0;
    }    
    Compare(OutDataFilePath, StdDataFilePath);
}
