#!/bin/sh

if [ -z "${GOROOT}" ]; then
	echo "$GOROOT is required" 
	exit 1
fi

while getopts s:j:d:u o
   do case "$o" in
		d)  DIR="$OPTARG";;
	   u)  UPLOAD="true";;
	   s)  SUFFIX="$OPTARG";;
		\?)  echo "Usage: $0 -d dist_dir [-s suffix]" && exit 1;;
	esac
done


#rm -rf php
#mkdir golang

#
# OS=`uname -s`
# if test "$OS" = "Darwin"; then
# cp ../../../build/usr/lib/libdjon-client.dylib ../../build/usr/lib/libdjon-client.dylib .
# else
# cp ../../../build/usr/lib/libdjon-client.so .
# fi

swig -c++ -go -intgosize 32 -o djongodriver.cpp driver-go.i 

gcc -c -fpic -I includes djongodriver.cpp
gcc -shared djongodriver.o -lstdc++ -ldjon-client -o djondbgo.so
6g djondb.go
6c -I $GOROOT/src/pkg/runtime/ djondb_gc.c
gopack grc djondb.a djondb.6 djondb_gc.6
exit 1

