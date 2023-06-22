export GOPROXY=https://goproxy.cn,direct
export CGO_ENABLED=0
TMP='/tmp'
OUTDIR='../output'
MAINNAME='ehfs'
MOD=$(go list ../src/)
BASEMOD=mjpclab.dev/ghfs/src
source ./build.inc.version.sh
LICENSE='../LICENSE'
getLdFlags() {
	echo "-s -w -X $BASEMOD/version.appVer=$VERSION -X $BASEMOD/version.appArch=${ARCH:-$(go env GOARCH)}"
}
