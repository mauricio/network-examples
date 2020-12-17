package redis

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScanLines(t *testing.T) {
	tt := []struct {
		data []byte
		eof bool
		advance int
		token []byte
		err string
	} {
		{
			data: []byte("+PING\r\n"),
			advance: 7,
			token: []byte("+PING"),
		},
		{
			data: []byte("+PING"),
			eof: true,
			err: "incomplete stream: all redis messages should end with a CRLF but this line did not have it and we're at EOF, this is a broken stream: [+PING]",
		},
		{
			data: []byte("+PING"),
		},
		{
			data: []byte("+PING\r\n+PONG"),
			advance: 7,
			token: []byte("+PING"),
		},
	}

	for _, ts := range tt {
		t.Run(string(ts.data), func(t *testing.T) {
			advance, token, err := ScanLines(ts.data, ts.eof)
			assert.Equal(t, ts.advance, advance)
			assert.Equal(t, ts.token, token)
			if err != nil || ts.err != "" {
				assert.EqualError(t, err, ts.err)
			}
		})
	}
}