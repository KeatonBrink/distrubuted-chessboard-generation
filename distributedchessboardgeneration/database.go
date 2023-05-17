package distributedchessboardgeneration

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

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

func Download(url, path string) error {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error download: http.Get, url: %v", url)
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error download: io.ReadAll", err)
		return err
	}
	file, err := os.Create(path)
	if err != nil {
		fmt.Printf("Error download: os.Create, Path: %v", path)
		return err
	}
	defer file.Close()
	_, err = file.Write(body)
	if err != nil {
		fmt.Println("Error download: file.Write", err)
		return err
	}
	return nil
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
