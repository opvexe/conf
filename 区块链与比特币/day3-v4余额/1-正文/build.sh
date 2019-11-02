#!/bin/bash
gitbook init
if [ -e SUMMARY.md ]; then
    rm SUMMARY.md
fi
echo "begin ..."

README="README.md"
for file in `ls  *.md*`
do
    displayName=${file%.md}
    echo -e "* [$displayName]($file)" >>SUMMARY.md
    echo -e "* [$displayName]($file)"
done

savepath=~/Desktop/区块链课程离线课件pdf版

filename=$savepath/$1/$2.pdf

echo "file : $filename"
gitbook pdf .  $filename

rm README.md SUMMARY.md
