package version

import (
	"mjpclab.dev/ghfs/src/version"
	"os"
)

func PrintVersion() {
	os.Stdout.WriteString("EHFS: Extra HTTP File Server\n")
	version.PrintVersion()
}
