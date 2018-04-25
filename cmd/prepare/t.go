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
	db, err := sql.Open("mysql", "root:root@tcp(192.168.1.7:1234)/db")
	handleErr(err)

	stmt, err := db.Prepare("select * from tb where id = ?")
	handleErr(err)
	fmt.Println(stmt)

	rs, err := stmt.Exec(1)
	handleErr(err)

	rs.RowsAffected()
	rs.LastInsertId()

}
