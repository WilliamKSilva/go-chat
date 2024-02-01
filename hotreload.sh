#!/bin/bash
GIT='git --git-dir='$PWD'/.git'

CHANGES=$($GIT diff)
echo "${CHANGES}"
