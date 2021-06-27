package dulldb

import (
	"io/ioutil"
	"os"
	"path"
	"reflect"
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

func TestSelect(t *testing.T) {
	dir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Error(err)
		return
	}
	defer os.RemoveAll(dir)

	file := path.Join(dir, "db.json")
	if err := ioutil.WriteFile(file, []byte(persisted), 0600); err != nil {
		t.Error(err)
		return
	}

	noPerm := path.Join(dir, "noperm.json")
	if err := ioutil.WriteFile(noPerm, []byte(persisted), 0000); err != nil {
		t.Error(err)
		return
	}

	invalid := path.Join(dir, "invalid.json")
	if err := ioutil.WriteFile(invalid, nil, 0600); err != nil {
		t.Error(err)
		return
	}

	noSuch := path.Join(dir, "nosuch.json")
	const wontVanish = "won't vanish"
	var data interface{} = wontVanish

	if err := Select(noSuch, &data); err == nil {
		if data != wontVanish {
			t.Errorf("Select(%#v, %#v): changed var passed by pointer to %#v", noSuch, &data, data)
		}
	} else {
		t.Errorf("Select(%#v, %#v): got %#v, expected nil", noSuch, &data, err)
	}

	if err := Select(file, &data); err == nil {
		if !reflect.DeepEqual(data, structure) {
			t.Errorf("Select(%#v, %#v): got %#v, expected %#v", noSuch, &data, data, structure)
		}
	} else {
		t.Errorf("Select(%#v, %#v): got %#v, expected nil", file, &data, err)
	}

	if Select(noPerm, &data) == nil {
		t.Errorf("Select(%#v, %#v): got nil, expected error", noPerm, &data)
	}

	if Select(invalid, &data) == nil {
		t.Errorf("Select(%#v, %#v): got nil, expected error", invalid, &data)
	}
}
