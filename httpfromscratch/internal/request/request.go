package request

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

type state int

const (
	stateInit state = 0
	stateDone state = 1
)

type Request struct {
	RequestLine RequestLine
	state       state
}

func NewRequest() *Request {
	return &Request{
		state: stateInit,
	}
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

var SEPARATOR = []byte("\r\n")
var SEPARATOR_LEN = len(SEPARATOR)

func parseRequestLine(data []byte) (*RequestLine, int, error) {
	var rl RequestLine

	line, _, found := bytes.Cut(data, []byte("\r\n"))
	if !found {
		return nil, len(line) + SEPARATOR_LEN, nil
	}

	parts := bytes.Split(line, []byte(" "))
	if len(parts) > 3 {
		return nil, len(line) + SEPARATOR_LEN, fmt.Errorf("invalid request line: %s", data)
	} else if len(parts) < 3 {
		return nil, len(line) + SEPARATOR_LEN, nil
	}

	rl.Method = string(parts[0])
	rl.RequestTarget = string(parts[1])
	rl.HttpVersion = string(parts[2])

	rl.HttpVersion = strings.TrimPrefix(rl.HttpVersion, "HTTP/")

	return &rl, len(line) + SEPARATOR_LEN, nil
}

func (r *Request) parse(data []byte) (int, error) {
	read := 0
outer:
	switch r.state {
	case stateInit:
		rl, n, err := parseRequestLine(data[read:])
		if err != nil {
			return 0, err
		}
		if rl == nil {
			break outer
		}
		r.RequestLine = *rl
		read += n
		r.state = stateDone
	case stateDone:
		break outer
	}
	return read, nil
}

func (r *Request) done() bool {
	return r.state == stateDone
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	request := NewRequest()

	buf := make([]byte, 8)
	bufLen := 0

	for !request.done() {
		n, err := reader.Read(buf[bufLen:])
		if err != nil {
			return nil, err
		}
		bufLen += n
		if bufLen > len(buf)/2 {
			newBuf := make([]byte, len(buf)*2)
			copy(newBuf, buf[:bufLen])
			buf = newBuf
		}
		readN, err := request.parse(buf[:bufLen+n])
		if err != nil {
			return nil, err
		}

		copy(buf, buf[readN:bufLen])
		bufLen -= readN
	}

	return request, nil
}
