/**
 *  author: lim
 *  data  : 18-4-24 下午9:52
 */

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(172.17.0.4:5518)/db")
	//db, err := sql.Open("mysql", "root:root@tcp(192.168.1.5:1234)/db")
	handleErr(err)

	//srvstmt, err := db.Prepare("insert into tb(id, tb.name) values (?, ?)")
	//stmt, err := db.Prepare("select tb.name, tb.id, tb.name,  tb.id, tb.id,  tb.name, tb.name, tb.name, tb.name, tb.name, tb.name, tb.name, tb.name, tb.name, tb.name, tb.name, tb.name, tb.name, tb.name, tb.name, tb.name from tb join tt on tb.id = tt.id")
	//snapstmt, err := db.Prepare("select tb.name, tb.id, tb.name,  tb.id, tb.id,  tb.name, tb.name, tb.name, tb.name, tb.name, tb.name, tb.name, tb.name, tb.name, tb.name, tb.name, tb.name, tb.name, tb.name, tb.name, tb.name from tb")
	stmt, err := db.Prepare("select * from tb")
	handleErr(err)
	fmt.Println("prepare done.")

	rs, err := stmt.Query()
	handleErr(err)

	var id int
	var name *string
	//var v uint64
	fmt.Println("---------------================")
	for rs.Next() {
		err = rs.Scan(&id, name)
		fmt.Println(err)
		fmt.Println("----", id, name)
	}
	err = stmt.Close()
	fmt.Println(err)
	err = stmt.Close()
	fmt.Println(err)
	fmt.Println("===========")
	err = stmt.Close()
	fmt.Println(err)

}
