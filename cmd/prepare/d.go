/**
 *  author: lim
 *  data  : 18-5-4 下午12:23
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
	//db, err := sql.Open("mysql", "root:root@tcp(172.17.0.2:5518)/db")
	db, err := sql.Open("mysql", "root:root@tcp(192.168.1.6:1234)/db")
	handleErr(err)

	//srvstmt, err := db.Prepare("insert into tb(id, name) values (?, ?)")
	stmt, err := db.Prepare("delete from tb where id = ?")
	handleErr(err)
	fmt.Println(stmt)

	x := 0
	for idx := x; idx < x+20; idx += 1 {
		rs, err := stmt.Exec(idx)

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
