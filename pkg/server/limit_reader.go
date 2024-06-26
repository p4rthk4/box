// U2SMTP - smtp server
//
// Licensed under the MIT License for individual use or a commercial
// license for enterprise use. See LICENSE.txt and COMMERCIAL_LICENSE.txt.

package server

import (
	"errors"
	"io"
)

type LimitReader struct {
	r         io.Reader
	lineLimit int

	curLineLength int
}

var errTooLongLine = errors.New("smtp: too long a line in input stream")
func (lr *LimitReader) Read(b []byte) (int, error) {
	if lr.curLineLength > lr.lineLimit && lr.lineLimit > 0 {
		return 0, errTooLongLine
	}

	n, err := lr.r.Read(b)
	if err != nil {
		return n, err
	}

	if lr.lineLimit == 0 {
		return n, nil
	}

	for _, chr := range b[:n] {
		if chr == '\n' {
			lr.curLineLength = 0
		}
		lr.curLineLength++

		if lr.curLineLength > lr.lineLimit {
			return 0, errTooLongLine
		}
	}

	return n, nil
}
