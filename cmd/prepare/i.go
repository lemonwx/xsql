/**
 *  author: lim
 *  data  : 18-4-24 下午9:52
 */

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(172.17.0.2:5518)/db")
	//db, err := sql.Open("mysql", "root:root@tcp(192.168.1.6:1234)/db")
	handleErr(err)

	//srvstmt, err := db.Prepare("insert into tb(id, name) values (?, ?)")
	stmt, err := db.Prepare("insert into tb (id, name) values (?, ?)")
	handleErr(err)
	fmt.Println(stmt)

	x := 1
	for idx := x; idx < x+20; idx += 1 {
		rs, err := stmt.Exec(idx, "name")
		handleErr(err)

		fmt.Println("---------------", idx)
		af, err := rs.RowsAffected()
		handleErr(err)
		id, err := rs.LastInsertId()
		handleErr(err)

		fmt.Println("---------------", af, id, idx)
		db.Exec("commit")
	}

	for {
		time.Sleep(time.Second * 10)
	}

}
