package midf

import (
	"errors"
	"path"
	"strings"
)

type MidFile struct {
	FilePath string
	FileName string
}

func NewMidMif(fullpath string) MidMif {
	midMif := MidMif{}
	//dir, fname := path.Split(fullpath)
	midMif.Mif, _ = NewMifFile(fullpath + ".mif")
	midMif.Mid, _ = NewMidFile(fullpath + ".mid")
	midMif.Datas = make([]MiData, 0, 100)
	return midMif
}

func NewMidFile(fullpath string) (*MidFile, error) {
	file := &MidFile{}
	if strings.EqualFold(path.Ext(fullpath), ".mid") {
		file.FilePath, file.FileName = path.Split(fullpath)
	} else {
		return nil, errors.New("文件扩展名不是.mid")
	}
	return file, nil
}
