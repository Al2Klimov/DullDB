package dulldb

import (
	"bytes"
	"encoding/json"
	"github.com/Al2Klimov/FUeL.go"
	"github.com/natefinch/atomic"
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
