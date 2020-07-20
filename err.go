package tinycmdb

import (
	"log"
)

// e is a shorty for an fatal error check
func e(err error) {
	if err != nil {
		log.Fatal(err)
	}

}

func dbe(err error, sqlStmt string) {
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}
