# DullDB

The most stupid (that's to say simple) actually useful database.

## Features

* K.I.S.S.
* NoSQL (but keep in mind that "schemaless" is a myth!)
* fully ACID
* multiple producer`*`, multiple consumer and lock-free
* web-scale ;-) (no limits on producers`*` or consumers amount)
* cross-platform (even a DullDB database by itself)
* open and well-known storage format
* read performance benefits if the OS caches disk data (e.g. Linux)

`*` if the OS allows to overwrite a file by renaming
while it's open and being read (e.g. *nix)

## How it works

The complete database is a single JSON file with just the payload.
It's written atomically, i.e. by `write(2)`, `fsync(2)`,
`close(2)` and `rename(2)`. This way:

* Changes to the data are **atomic** due to `rename(2)`.
* The data stays **consistent** due to `fsync(2)`.
* "Transactions" (that's to say writes) are **isolated** from each other
  due to `rename(2)` (the last rename wins).
* "Commits" (that's to say writes) are **durable** due to `fsync(2)`.
* **Multiple consumers** can `read(2)` the file at the same time,
  even while **multiple producers** are performing `write(2)` or `rename(2)`
  **without locking** anything.

(The [features](#Features) not explained by the bullet points above
are implied by the database format itself.)

## Drawbacks

The whole database has to be read or written at once.

## Usage

```bash
go get github.com/Al2Klimov/DullDB/v1
```

```go
package main

import "github.com/Al2Klimov/DullDB/v1"

func main() {
	err := dulldb.Replace(
		"persons.json",
		[]struct{ Name string }{{"Alice"}, {"Bob"}, {"Carl"}},
	)
	if err != nil {
		panic(err)
	}
}
```

```go
package main

import (
	"fmt"
	"github.com/Al2Klimov/DullDB/v1"
)

func main() {
	var persons []struct{ Name string }
	if err := dulldb.Select("persons.json", &persons); err != nil {
		panic(err)
	}
	fmt.Print(persons)
}
```

Result: `[{Alice} {Bob} {Carl}]`
