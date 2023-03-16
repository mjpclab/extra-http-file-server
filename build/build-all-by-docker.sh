#!/bin/bash

cd "$(dirname "$0")"

# init variable `builds`
source ./build-all.inc.sh

prefix=$(realpath ../)
ehfs=/go/src/mjpclab.dev/ehfs

rm -rf "$prefix/output/"

buildByDocker() {
  gover="$1"
  shift
  docker pull golang:"$gover"

  docker run \
    --rm \
    -v "$prefix":"$ehfs" \
    -e EX_UID="$(id -u)" \
    -e EX_GID="$(id -g)" \
    golang:"$gover" \
    /bin/bash -c '
      sed -i -e "s;://[^/ ]*;://mirrors.aliyun.com;" /etc/apt/sources.list;
      apt-get update;
      apt-get install -yq git zip;
      git config --global safe.directory "*"
      /bin/bash '"$ehfs"'/build/build.sh "$@";
      chown -R $EX_UID:$EX_GID '"$ehfs"'/output;
    ' \
    'argv_0_placeholder' \
    "$@"
}

gover=latest
buildByDocker "$gover" "${builds[@]}"

#gover=1.20
#builds=()
#builds+=('windows 386 -7-8' 'windows amd64 -7-8')
#builds+=('windows amd64,v2 -7-8' 'windows amd64,v3 -7-8')
#builds+=('darwin amd64 -10.13-high-sierra')
#buildByDocker "$gover" "${builds[@]}"

#gover=1.16
#builds=('darwin amd64 -10.12-sierra')
#buildByDocker "$gover" "${builds[@]}"
