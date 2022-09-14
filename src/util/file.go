package util

import (
	ghfsUtil "mjpclab.dev/ghfs/src/util"
	"os"
)

func GetFileInfoType(filename string) (file *os.File, info os.FileInfo, contentType string, err error) {
	file, err = os.Open(filename)
	if err != nil {
		return
	}

	info, err = file.Stat()
	if err != nil {
		file.Close()
		file = nil
		return
	}

	contentType, err = ghfsUtil.GetContentType(filename, file)
	if err != nil {
		file.Close()
		file = nil
		return
	}

	return
}
