package ex7_5

import "io"

type limitReader struct {
	reader   io.Reader
	n, limit int
}

func (l *limitReader) Read(p []byte) (n int, err error) {
	n, err = l.reader.Read(p[:l.limit])
	l.n += n
	if l.n > l.limit {
		err = io.EOF
	}
	return
}

func LimitReader(reader io.Reader, n int) io.Reader {
	return &limitReader{reader: reader, limit: n}
}
