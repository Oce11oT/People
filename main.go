package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)
func main() {
	connStr:="user=postgres password=people1234 dbname=people sslmode=disable"
	db, err := sql.Open("postgres",connStr)

	if err!= nil{
		panic(err)
	}
	defer db.Close()
	result, err := db.Exec("insert into people (uname,sex,age)" +
		" values ('Emilia','female',21)")
	if err !=nil{
		panic(err)
	}
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())
}
