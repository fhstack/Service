#!/bin/bash
ps -aux | grep $1 | awk '{print $6}' | head -1
