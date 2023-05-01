package distributedchessboardgeneration

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDatabase(path string) (*sql.DB, error) {
	options :=
		"?" + "_busy_timeout=10000" +
			"&" + "_case_sensitive_like=OFF" +
			"&" + "_foreign_keys=ON" +
			"&" + "_journal_mode=OFF" +
			"&" + "_locking_mode=NORMAL" +
			"&" + "mode=rw" +
			"&" + "_synchronous=OFF"
	db, err := sql.Open("sqlite3", path+options)
	if err != nil {
		// handle the error here
		return nil, err
	}
	return db, nil
}

func PrintSizeOfDataInTable(dbPath, table string) error {
	db, err := OpenDatabase(dbPath)
	if err != nil {
		return err
	}
	results, err := db.Query("SELECT COUNT(*) FROM ?", table)
	if err != nil {
		return err
	}
	results.Next()
	var tableSize int
	err = results.Scan(&tableSize)
	if err != nil {
		return err
	}
	log.Printf("Size of table: %v in file: %v, %v", table, dbPath, tableSize)
	return nil
}
