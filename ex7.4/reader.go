package reader

import (
	"io"
)

type stringReader struct {
	s string
}

func (r *stringReader) Read(p []byte) (n int, err error) {
	n = copy(p, r.s)//assign directly to named return var
	r.s = r.s[n:]
	if len(r.s) == 0 {
		err = io.EOF
	}
	return
}

//NewReader stringReader implements Reader interface
func NewReader(s string) io.Reader {//io.Reader is an interface
	// TODO: why the return signature can't be *io.Reader
	//https://stackoverflow.com/questions/13511203/why-cant-i-assign-a-struct-to-an-interface
	return &stringReader{s} //stringReader is a struct, 
}
