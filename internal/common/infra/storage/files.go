package storage

import (
	"fmt"
	"strings"

	"github.com/charmingruby/mvplease/pkg/errors"
)

func NewFile(name, mimetype string, size int64, validMimetypes []string, maxSizeInBytes int64) (*File, error) {
	f := File{
		Name:     name,
		Mimetype: mimetype,
		Size:     size,
	}

	if err := f.validate(validMimetypes, maxSizeInBytes); err != nil {
		return nil, err
	}

	return &f, nil
}

type File struct {
	Name     string `json:"name"`
	Mimetype string `json:"mimetype"`
	Size     int64  `json:"size"` // size in bytes
}

func (f *File) validate(validMimetypes []string, maxSizeInBytes int64) error {
	var matchAMimetype bool

	for _, mimetype := range validMimetypes {
		if f.Mimetype == mimetype {
			matchAMimetype = true
		}
	}

	if !matchAMimetype {
		return errors.NewInvalidMimetypeError(nil, f.Mimetype, validMimetypes)
	}

	if f.Size > maxSizeInBytes {
		return errors.NewExceedsMaxSizeError(nil, f.Size, maxSizeInBytes)
	}

	return nil
}

func SplitFileData(filename string) (string, string, error) {
	agg := strings.Split(filename, ".")

	if len(agg) < 2 {
		return "", "", fmt.Errorf("invalid file")
	}

	file := agg[0]
	mimetype := agg[1]

	return file, mimetype, nil
}

func MBToBytes(value int) int {
	MBInBytes := 1000000
	return MBInBytes * value
}
