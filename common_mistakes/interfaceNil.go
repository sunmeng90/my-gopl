package main

import (
	"bytes"
	"io"
)

const debug = false

func main(){
	var buf *bytes.Buffer
	if debug {
		buf = new(bytes.Buffer)
	}
	// assign a nil pointer of type *bytes.Buffer to out parameter
	f(buf)
	if debug{
		//
	}
}

func f(out io.Writer){
	if out != nil {// guard check for nil
		out.Write([]byte("done\n")) // panic: nil pointer dereference
	}
}

////////////////////////////////
// type: *bytes.Buffer
// value: nil
// A non-nil interface containing a nil pointer
///////////////////////////////

//nil receiver call print hello world