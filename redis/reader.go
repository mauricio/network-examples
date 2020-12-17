package redis

import (
	"bytes"
	"fmt"
)

var (
	crlfBytes = []byte{'\r','\n'}
)

func ScanLines(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// empty array, ask for more data
	if len(data) == 0 {
		return 0, nil, nil
	}

	nextCrlf := bytes.Index(data, crlfBytes)

	if nextCrlf >= 0 {
		// here's a full crlf-terminated line.
		return nextCrlf + 2, data[0:nextCrlf], nil
	}
	// If we're at EOF and there's no crlf, this is a broken stream, error it
	if atEOF {
		return 0, nil, fmt.Errorf("incomplete stream: all redis messages should end with a CRLF but this line did not have it and we're at EOF, this is a broken stream: [%v]", string(data))
	}

	// need more data
	return 0, nil, nil
}