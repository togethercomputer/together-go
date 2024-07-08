// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package together

import (
	"bufio"
	"bytes"
	"encoding/json"
	"net/http"
)

type Decoder struct {
	Event Event

	res *http.Response
	scn *bufio.Scanner
	err error
}

type Event struct {
	Type string
	Data []byte
}

func NewDecoder(res *http.Response) *Decoder {
	if res == nil || res.Body == nil {
		return nil
	}
	scn := bufio.NewScanner(res.Body)
	return &Decoder{
		res: res,
		scn: scn,
	}
}

func (s *Decoder) Next() bool {
	if s.err != nil {
		return false
	}

	event := ""
	data := bytes.NewBuffer(nil)

	for s.scn.Scan() {
		txt := s.scn.Bytes()

		// Dispatch event on an empty line
		if len(txt) == 0 {
			s.Event = Event{
				Type: event,
				Data: data.Bytes(),
			}
			return true
		}

		// Split a string like "event: bar" into name="event" and value=" bar".
		name, value, _ := bytes.Cut(txt, []byte(":"))

		// Consume an optional space after the colon if it exists.
		if len(value) > 0 && value[0] == ' ' {
			value = value[1:]
		}

		switch string(name) {
		case "":
			// An empty line in the for ": something" is a comment and should be ignored.
			continue
		case "event":
			event = string(value)
		case "data":
			_, s.err = data.Write(value)
			if s.err != nil {
				break
			}
			_, s.err = data.WriteRune('\n')
			if s.err != nil {
				break
			}
		}
	}

	return false
}

func (s *Decoder) Close() error {
	return s.res.Body.Close()
}

type Stream[T any] struct {
	decoder *Decoder
	cur     T
	err     error
	done    bool
}

func (s *Stream[T]) Next() bool {
	if s.err != nil {
		return false
	}

	for s.decoder.Next() {
		if s.done {
			return false
		}

		if bytes.HasPrefix(s.decoder.Event.Data, []byte("[DONE]")) {
			s.done = true
			return false
		}

		if s.decoder.Event.Type == "" {
			s.err = json.Unmarshal(s.decoder.Event.Data, &s.cur)
			if s.err != nil {
				return false
			}
			return true
		}
	}

	return false
}

func (s *Stream[T]) Current() T {
	return s.cur
}

func (s *Stream[T]) Err() error {
	return s.err
}

func (s *Stream[T]) Close() error {
	return s.decoder.Close()
}
