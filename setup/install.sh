#!/bin/bash

# installing gor on your local system
# for macOS
curl https://github.com/buger/goreplay/archive/refs/tags/v2.0.0-rc2.tar.gz -O -J -L

# unzip the downloaded file
tar -xf goreplay-2.0.0-rc2.tar.gz

# remove zip file
rm -rf goreplay-2.0.0-rc2.tar.gz

# make directory
mkdir /usr/local/bin/gor

# shellcheck disable=SC2164
cd goreplay-2.0.0-rc2

# build golang executable file
go build -o ./gor

# move the files to bin
# shellcheck disable=SC2035
mv * /usr/local/bin/gor

# shellcheck disable=SC2103
cd ..

# remove empty directory
rm -rf goreplay-2.0.0-rc2