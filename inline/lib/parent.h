#pragma once
#include <sys/types.h>

void InitSyscallLimits();
void WatchChild(pid_t c_pid);
