package keju

import (
	"compress/zlib"
	"io"
	"os"
)

const RepositoryFile = `./question.gob`

// flag parameter is the same as os.OpenFile
func OpenRepositoryFile(flag int) (*os.File, error) {
	return os.OpenFile(RepositoryFile, flag, 0755)
}

func OpenRepositoryFileForRead() (io.ReadCloser, error) {
	f, err := OpenRepositoryFile(os.O_RDONLY)
	if err != nil {
		return f, err
	}
	return zlib.NewReader(f)
}

func OpenRepositoryFileForWriteNewRepo() (io.WriteCloser, error) {
	f, err := OpenRepositoryFile(os.O_WRONLY | os.O_CREATE | os.O_TRUNC)
	if err != nil {
		return f, err
	}
	return zlib.NewWriterLevel(f, zlib.BestCompression)
}
