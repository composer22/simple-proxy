#!/bin/bash

# Build all executables
#
# ex: ./build.sh
#
# 386 = 32 bit
# amd = 64 bit
#
# $GOOS			$GOARCH
# ====================
# darwin		386
# darwin		amd64 *
# linux			386
# linux			amd64 *
# windows		386
# windows		amd64 *
# darwin	    386
# darwin	    amd64
# darwin	    arm
# darwin	    arm64
# dragonfly	    amd64
# freebsd	    386
# freebsd	    amd64
# freebsd	    arm
# linux	        386
# linux	       amd64
# linux	        arm
# linux	    arm64
# linux	    ppc64
# linux	    ppc64le
# linux	    mips64
# linux	    mips64le
# netbsd	    386
# netbsd	    amd64
# netbsd	    arm
# openbsd	    386
# openbsd	    amd64
# openbsd	    arm
# plan9	    386
# plan9	    amd64
# solaris	    amd64
# windows	    386
# windows	    amd64

type setopt >/dev/null 2>&1 && setopt shwordsplit
PLATFORMS="darwin/amd64 linux/amd64 windows/amd64"

function go-compile {
	local GOOS=${1%/*}
	local GOARCH=${1#*/}
	GOOS=${GOOS} GOARCH=${GOARCH} go build -o simple-proxy-${GOOS}-${GOARCH} -i
}

function run {
	for PLATFORM in $PLATFORMS; do
			local CMD="go-compile ${PLATFORM}"
			echo "$CMD"
			$CMD
	done
}

run
mv simple-proxy-windows-amd64 simple-proxy-windows-amd64.exe
