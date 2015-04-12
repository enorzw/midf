package midf

import (
	"errors"
	"path"
	"strings"
)

/*
MIF文件封装
*/

type MifFile struct {
	FilePath string
	FileName string
	Head     MifHead
}

func NewMifFile(fullpath string) (*MifFile, error) {
	file := &MifFile{}
	file.Head = NewMifHead()
	if strings.EqualFold(path.Ext(fullpath), ".mif") {
		file.FilePath, file.FileName = path.Split(fullpath)
	} else {
		return nil, errors.New("文件扩展名不是.mif")
	}
	return file, nil
}
