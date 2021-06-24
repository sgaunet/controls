#!/usr/bin/env bash


cwd=$(dirname $0)
cd $cwd
tstDir=$(pwd)

projectWorkdir=$(dirname $tstDir)

# Generate controls.yaml file with template controls.tmpl
sed "s#TSTDIR#$tstDir#g" "$tstDir/controls.tmpl" > "$tstDir/controls.yaml"

cd "$projectWorkdir/src"
go run . -f "$tstDir/controls.yaml" -o "$tstDir/result.md"


# mdtohtml is also a project maintained by the author (forker originaly)
# This command read markdown format and generate html (use wkhtmltopdf after to get a PDF)
# https://github.com/sgaunet/mdtohtml
which mdtohtml > /dev/null 2>&1
rc=$?

if [ "$rc" = "0" ]
then
    mdtohtml "$tstDir/result.md" "$tstDir/result.html"
    firefox "$tstDir/result.html"
fi
