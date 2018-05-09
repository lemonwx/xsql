/**
 *  author: lim
 *  data  : 18-5-8 下午9:22
 */

package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	//db, err := sql.Open("mysql", "root:root@tcp(192.168.1.6:1234)/db")
	db, err := sql.Open("mysql", "root:root@tcp(172.17.0.3:5518)/db")
	handleErr(err)
	rs, err := db.Exec("begin;select * from db.tb where id = 1")
	handleErr(err)

	fmt.Println(rs)

}
