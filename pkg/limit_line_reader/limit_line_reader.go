// U2SMTP - limit line reader
//
// Licensed under the MIT License for individual use or a commercial
// license for enterprise use. See LICENSE.txt and COMMERCIAL_LICENSE.txt.

package limitlinereader

import (
	"errors"
	"io"
)

// read lines but any line size is more
// then LimitLineReader than return
// ErrTooLongLine error
type LimitLineReader struct {
	Reader      io.Reader
	MaxLineSize int

	curLineLength int
}

var ErrTooLongLine = errors.New("smtp: too long a line in input stream")

func (lr *LimitLineReader) Read(b []byte) (int, error) {
	if lr.curLineLength > lr.MaxLineSize && lr.MaxLineSize > 0 {
		return 0, ErrTooLongLine
	}

	n, err := lr.Reader.Read(b)
	if err != nil {
		return n, err
	}

	if lr.MaxLineSize == 0 {
		return n, nil
	}

	for _, chr := range b[:n] {
		if chr == '\n' {
			lr.curLineLength = 0
		}
		lr.curLineLength++

		if lr.curLineLength > lr.MaxLineSize {
			return 0, ErrTooLongLine
		}
	}

	return n, nil
}
