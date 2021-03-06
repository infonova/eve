#!/bin/sh

# get parent directory
source="${BASH_SOURCE[0]}"
while [[ -h "${source}" ]]; do 
    source="$(readlink "${source}")"
done
dir="$(cd -P "$(dirname "${source}")/.." && pwd )"

cd ${dir}

# remove old folders and create new ones
echo "==> Removing old folders"
rm -rf bin/*
rm -rf log/*
mkdir -p bin/
mkdir -p log/

# get version
git_commit=$(git rev-parse HEAD)

# build
echo "==> Building..."
source ./version.info
go build
cp eve bin/
cp scripts/eve.sh bin/
echo ${eve_version} > etc/eve.version
echo ${git_commit} >> etc/eve.version
cd ..
tar -czf eve-${eve_version}.tar.gz eve/bin eve/log eve/etc/eve.toml.sample eve/etc/eve.version --transform s/eve/eve-${eve_version}/
cd -
mv ../eve-${eve_version}.tar.gz .
