package util

import (
	"io/fs"
)

type renamedFileInfo struct {
	name string
	fs.FileInfo
}

func (info renamedFileInfo) Name() string {
	return info.name
}

func CreateRenamedFileInfo(name string, fileInfo fs.FileInfo) renamedFileInfo {
	return renamedFileInfo{name, fileInfo}
}
