#!/usr/bin/env bash


cwd=$(dirname $0)
cd $cwd
tstDir=$(pwd)

projectWorkdir=$(dirname $tstDir)


test -f "$tstDir/result.html" && \rm "$tstDir/result.html"
test -f "$tstDir/result.md"  && \rm "$tstDir/result.md"

# Generate controls.yaml file with template controls.tmpl
sed "s#TSTDIR#$tstDir#g" "$tstDir/controls.tmpl" > "$tstDir/controls.yaml"

cd "$projectWorkdir/src"
go run . -f "$tstDir/controls.yaml" -o "$tstDir/result.pdf"
rc=$?

if [ "$rc" != "0" ]
then
    echo "Error when executing go run ..."
    exit 1
fi
