#!/bin/bash
set -x

echo "Windows"
GOOS=windows GOARCH=amd64 GO_EXTLINK_ENABLED=0 CGO_ENABLED=1 go build
if [[ $? -eq 0 ]]
then
  mv main.exe mp4dumpgo.exe
  compress -f mp4dumpgo.exe
fi

echo "linux"
GOOS=linux GOARCH=amd64 GO_EXTLINK_ENABLED=0 CGO_ENABLED=1 go build  
if [[ $? -eq 0 ]]
then
  mv main mp4dumpgo_linux
  compress -f mp4dumpgo_linux
fi

echo "mac"
GOOS=darwin GOARCH=amd64 GO_EXTLINK_ENABLED=0 CGO_ENABLED=1 go build 
if [[ $? -eq 0 ]]
then
  mv main mp4dumpgo_mac
  compress -f mp4dumpgo_mac
fi


