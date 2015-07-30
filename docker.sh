#!/bin/bash

mkdir -p docker-build-tmp
chmod +s docker-build-tmp

# Build all executables in the golang-godep container.  Output statically linked binaries to docker-build-tmp
docker run -e "GOBIN=/usr/src/go/src/github.com/GeoNet/haz/docker-build-tmp"  -e "CGO_ENABLED=0" -e "GOOS=linux" --rm -v \
"$PWD":/usr/src/go/src/github.com/GeoNet/haz -w /usr/src/go/src/github.com/GeoNet/haz quay.io/geonet/golang-godep godep go install -a -installsuffix cgo ./...

mkdir -p docker-build-tmp/common/etc/ssl/certs
mkdir -p docker-build-tmp/common/usr/share

cp passwd.minimal docker-build-tmp/common/etc/passwd
cp /etc/ssl/certs/ca-certificates.crt docker-build-tmp/common/etc/ssl/certs
# An alternative is to use $GOROOT/lib/time/zoneinfo.zip
rsync --archive /usr/share/zoneinfo docker-build-tmp/common/usr/share

# For geonet-rest
rsync --archive geonet-rest/tmpl docker-build-tmp

# Copy in and rename the Dockerfiles.  Exclude the db Dockerfile.
find . -name 'Dockerfile' ! -path "./docker-build-tmp/*" | awk -F '/' '{print "cp "$0" docker-build-tmp/"$3"-"$2}' | sh
# Copy in the json config file for each excutable
find . -name 'Dockerfile' ! -path "./docker-build-tmp/*" | awk -F '/' '{print "cp "$2"/"$2".json docker-build-tmp/"}' | sh

build=(`ls docker-build-tmp/Dockerfile-* | awk -F 'Dockerfile-' '{print $2}'`)

for i in "${build[@]}"
do
    :
    docker build --rm=true -t quay.io/geonet/haz:$i -f docker-build-tmp/Dockerfile-$i docker-build-tmp
done

