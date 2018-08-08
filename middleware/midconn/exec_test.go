/**
 *  author: lim
 *  data  : 18-8-8 下午9:10
 */

package midconn

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func BenchmarkCaseWhenUpdate(b *testing.B) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/db",
		"root", "root", "172.17.0.2", 5518))
	if err != nil {
		b.Error(err)
	}

	tx, _ := db.Begin()

	for idx := 0; idx < b.N; idx += 1 {
		caseWhen := "case when version < 100000 and version not in (1,2,3,4,5,5,6,7,8,9,10,31,2,3,4,5,5,6,7,8,9,10,31,2,3,4,5,5,6,7,8,9,10,31,2,3,4,5,5,6,7,8,9,10,31,2,3,4,5,5,6,7,8,9,10,31,2,3,4,5,5,6,7,8,9,10,31,2,3,4,5,5,6,7,8,9,10,31,2,3,4,5,5,6,7,8,9,10,31,2,3,4,5,5,6,7,8,9,10,31,2,3,4,5,5,6,7,8,9,10,31,2,3,4,5,5,6,7,8,9,10,31,2,3,4,5,5,6,7,8,9,10,31,2,3,4,5,5,6,7,8,9,10,31,2,3,4,5,5,6,7,8,9,10,31,2,3,4,5,5,6,7,8,9,10,31,2,3,4,5,5,6,7,8,9,10,31,2,3,4,5,5,6,7,8,9,10,31,2,3,4,5,5,6,7,8,9,10,31,2,3,4,5,5,6,7,8,9,10,31,2,3,4,5,5,6,7,8,9,10,3)"
		_, err := db.Exec(fmt.Sprintf("update tb set version = (%s then 1007 end), name = (%s then 3 end) where id = 2", caseWhen, caseWhen))
		if err != nil {
			b.Error(err)
		}
		//db.Exec("update tb set name = 3 where id = 2")
	}

	tx.Rollback()
	db.Close()
}

func BenchmarkOrigalUpdate(b *testing.B) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/db",
		"root", "root", "172.17.0.2", 5518))
	if err != nil {
		b.Error(err)
	}

	tx, _ := db.Begin()

	for idx := 0; idx < b.N; idx += 1 {

		r, _ := tx.Query("select version from tb where id = 2")
		for r.Next() {
		}
		db.Exec("update tb set name = 3 where id = 2")
	}

	tx.Rollback()
	db.Close()
}
