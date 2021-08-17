#!/bin/sh

for file in ./*.ogg; do
	ffmpeg -i $file -f s16le -ar 48000 -ac 2 pipe:1 | dca >./dca/${file%.*}.dca
done
