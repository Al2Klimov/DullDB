package dulldb

import (
	"io/ioutil"
	"os"
	"path"
	"testing"
)

var structure = map[string]interface{}{"sense_of_life": 42.0}

const persisted = `{"sense_of_life":42}
`

func TestReplace(t *testing.T) {
	dir, errTD := ioutil.TempDir("", "")
	if errTD != nil {
		t.Error(errTD)
		return
	}
	defer os.RemoveAll(dir)

	file := path.Join(dir, "db.json")

	if err := Replace(file, structure); err != nil {
		t.Errorf("Replace(%#v, %#v): got %#v, expected nil", file, structure, err)
		return
	}

	written, errRF := ioutil.ReadFile(file)
	if errRF != nil {
		t.Errorf("Replace(%#v, %#v): can't read written data: %#v", file, structure, errRF)
		return
	}

	if string(written) != persisted {
		t.Errorf("Replace(%#v, %#v): written %#v, expected %#v", file, structure, string(written), persisted)
	}
}
