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
	db, err := sql.Open("mysql", "root:root@tcp(172.17.0.3:5518)/db")
	//snapdb, err := sql.Open("mysql", "root:root@tcp(192.168.1.7:1234)/db")
	handleErr(err)

	stmt, err := db.Prepare("select * from tb where id = ?")
	handleErr(err)
	fmt.Println(stmt)

	rs, err := stmt.Query(2)
	handleErr(err)

	var id int
	var name string
	fmt.Println("---------------")
	for rs.Next() {
		err = rs.Scan(&id, &name)
		fmt.Println(err)
		fmt.Println("----", id, name)
	}

}
