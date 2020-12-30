package controller

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/olekukonko/tablewriter"
	"log"
	"os"
	"strconv"
	"time"
)

func NewCtt(title string, createdAt int64, modifiedAt int64, content string) (int64, error) {
	db := Connect()
	insertCttSQL := `INSERT INTO ctt(title, createdAt, modifiedAt,content) VALUES (?, ?, ?,?)`
	statement, err := db.Prepare(insertCttSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	res, err := statement.Exec(title, createdAt, modifiedAt, content)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return res.LastInsertId()
}

func GetCttAll() {
	db := Connect()
	row, err := db.Query("SELECT * FROM ctt")
	if err != nil {
		fmt.Println("Error")
		log.Fatal(err)
	}
	defer row.Close()
	data := [][]string{}
	for row.Next() { // Iterate and fetch the records from result cursor
		var id string
		var title string
		var createdAt string
		var modifiedAt string
		var content string
		row.Scan(&id, &title, &createdAt, &modifiedAt, &content)
		fmt.Println("")
		tt, _ := strconv.ParseInt(modifiedAt, 10, 64)
		data = append(data, []string{id, title, time.Unix(tt, 0).String()})
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Code", "Title", "Changed"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render() // Send output
}
