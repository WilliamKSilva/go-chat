#!/bin/bash

# Use this for development purposes only #
# Its so boring to call GO interpreter everytime # 

ROOTDIR=$(dirname "$0")
ENTRYPOINT_SERVER="${ROOTDIR}/cmd/server/main.go"
GIT='git --git-dir='$PWD'/.git'
GOCOMPILER='/usr/local/go/bin/go'

FIRST_COMPILE_DONE=false
LAST_CHANGES_SIZE=0
while true
do
    CHANGES=$($GIT diff)
    CHANGES_OUTPUT_LENGTH=$(printf "%s" "$CHANGES" | wc -c)
    if (($CHANGES_OUTPUT_LENGTH == $LAST_CHANGES_SIZE)) && $FIRST_COMPILE_DONE;
    then
        continue
    fi

    LAST_CHANGES_SIZE=$CHANGES_OUTPUT_LENGTH

<<<<<<< HEAD
    if (( !$FIRST_COMPILE_DONE )) || (($CHANGES_OUTPUT_LENGTH > 0)); then
        if (( !$FIRST_COMPILE_DONE )); then
=======
    if ((!$FIRST_COMPILE_DONE)) || (($CHANGES_OUTPUT_LENGTH > 0));
    then
        if ((!$FIRST_COMPILE_DONE));
        then
>>>>>>> parent of f7ab3c7 (fix: change recompile condition order)
            FIRST_COMPILE_DONE=true
        fi

        echo "Recompiled!"
        $GOCOMPILER run $ENTRYPOINT_SERVER &
    fi
done
