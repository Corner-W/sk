package telnet

import (
	"io"
)

// LongWriteByte trys to write the byte from 'b' to the writer 'w', such that it deals
// with "short writes" where w.Write would return an error of io.ErrShortWrite and
// n < 1.
//
// Note that LongWriteByte still could return the error io.ErrShortWrite; but this
// would only be after trying to handle the io.ErrShortWrite a number of times, and
// then eventually giving up.
func LongWriteByte(w io.Writer, b byte) error {
	var buffer [1]byte
	p := buffer[:]

	buffer[0] = b

	numWritten, err := LongWrite(w, p)
	if 1 != numWritten {
		return io.ErrShortWrite
	}

	return err
}
func LongWriteString(w io.Writer, s string) (int64, error) {

	numWritten := int64(0)
	for {
		//@TODO: Should check to make sure this doesn't get stuck in an infinite loop writting nothing!
		n, err := io.WriteString(w, s)
		numWritten += int64(n)
		if nil != err && io.ErrShortWrite != err {
			return numWritten, err
		}

		if !(n < len(s)) {
			break
		}

		s = s[n:]

		if len(s) < 1 {
			break
		}
	}

	return numWritten, nil
}

func LongWrite(w io.Writer, p []byte) (int64, error) {

	numWritten := int64(0)
	for {
		//@TODO: Should check to make sure this doesn't get stuck in an infinite loop writting nothing!
		n, err := w.Write(p)
		numWritten += int64(n)
		if nil != err && io.ErrShortWrite != err {
			return numWritten, err
		}

		if !(n < len(p)) {
			break
		}

		p = p[n:]

		if len(p) < 1 {
			break
		}
	}

	return numWritten, nil
}
