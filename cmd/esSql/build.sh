#!/bin/bash

PROJECT=`pwd | awk -F / ' {print $NF}'`

echo PROJECT: $PROJECT
sleep 5

goos=(
linux
windows
darwin
)

for os in ${goos[*]}
do
if [ $os == "windows" ];then
suffix=.exe
else
unset suffix
fi

echo CGO_ENABLED=0 GOOS=$os GOARCH=amd64 go build -o $PROJECT-$os$suffix
CGO_ENABLED=0 GOOS=$os GOARCH=amd64 go build -o $PROJECT-$os$suffix
done
