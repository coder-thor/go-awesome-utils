package system

import (
	"os"
	"path/filepath"

	"gopkg.in/ini.v1"
)

func IniConfLoader(path string) (*ini.File, error) {
	filename, err := os.Executable()
	if err != nil {
		return nil, err
	}

	dirname := filepath.Dir(filename)

	file, lookupErr := ini.Load(dirname + "/" + path)

	if lookupErr != nil {
		return nil, lookupErr
	}

	return file, lookupErr
}
