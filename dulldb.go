package dulldb

import (
	"bytes"
	"encoding/json"
	"github.com/Al2Klimov/FUeL.go"
	"github.com/natefinch/atomic"
	"os"
)

// Replace saves the data $value into the database file $into (as with json.Marshal).
func Replace(into string, value interface{}) fuel.ErrorWithStack {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)

	enc.SetEscapeHTML(false)

	if err := enc.Encode(value); err != nil {
		return fuel.AttachStackToError(err, 0)
	}

	return fuel.AttachStackToError(atomic.WriteFile(into, buf), 0)
}

// Select loads the database file $from into *$to (as with json.Unmarshal).
func Select(from string, to interface{}) fuel.ErrorWithStack {
	f, err := os.Open(from)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}

		return fuel.AttachStackToError(err, 0)
	}
	defer f.Close()

	return fuel.AttachStackToError(json.NewDecoder(f).Decode(to), 0)
}
