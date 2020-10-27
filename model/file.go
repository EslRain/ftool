package model

import (
	"path/filepath"
	"strings"
)

type File struct {
	FullPath string
	Folder   string
	File     string
	FileName string
	Ext      string
}

func ConstructFile(path string) File {
	base := filepath.Base(path) //文件名带后缀
	ext := filepath.Ext(base)

	return File{
		FullPath: path,
		File:     base,
		Folder:   strings.TrimSuffix(path, base),
		FileName: strings.TrimSuffix(base, ext),
		Ext:      ext,
	}
}
