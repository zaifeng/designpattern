package strategy

import "os"

type FileSave struct {
}

func (fs *FileSave) Save(name string, data []byte) error {
	return os.WriteFile(name, data, 0755)
}
