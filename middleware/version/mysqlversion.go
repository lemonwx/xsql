/**
 *  author: lim
 *  data  : 18-4-20 下午12:22
 */

package version

/*
import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lemonwx/log"
	"sync"
)

var conn string
var db *sql.DB
var batch uint64
var curV uint64
var batchMax uint64
var lock sync.RWMutex

func NewRpcPool(size int, addr string) {
	conn = "root:root@tcp(172.17.0.4:5518)/v"
	db, _ = sql.Open("mysql", conn)
	batch = 2000
	curV = 0

	batchMax = 0
	db.SetMaxIdleConns(10)
	db.SetMaxIdleConns(100)
}

func NextVersion() (uint64, error) {
	lock.Lock()
	if curV+1 <= batchMax {
		curV += 1
		defer lock.Unlock()
		return curV, nil
	}
	lock.Unlock()

	trx, err := db.Begin()
	if err != nil {
		return 0, err
	}

	var curId uint64 = 0
	rows := trx.QueryRow("select id from v for update")
	if err != nil {
		log.Debug(err)
		trx.Rollback()
		return 0, err
	}

	err = rows.Scan(&curId)
	if err != nil {
		trx.Rollback()
		return 0, err
	}

	log.Debug(curId)
	var rs sql.Result

	if curId != 0 {
		rs, err = trx.Exec("update v set id = ?", curId+batch)
	} else {
		rs, err = trx.Exec("insert into v(id) values (?)", curId+batch)
	}

	if err != nil {
		log.Debug(err, rs)
		trx.Rollback()
		return 0, err
	}

	stmt, err := trx.Prepare("insert into v_in_use (id) values (?)")
	for idx := curId + 1; idx <= curId+batch; idx += 1 {
		stmt.Exec(idx)
	}

	if err != nil {
		trx.Rollback()
		return 0, err
	}

	err = trx.Commit()
	if err != nil {
		return 0, err
		trx.Rollback()
	}

	lock.Lock()
	curV = curId + 1
	batchMax = curV + batch
	lock.Unlock()

	return curId + 1, nil
}

func ReleaseVersion(version uint64) error {
	_, err := db.Exec("delete from v_in_use where id = ?", version)
	return err
}

func VersionsInUse() (map[uint64]uint8, error) {
	rows, err := db.Query("select id from v_in_use")
	if err != nil {
		return nil, err
	}
	ret := make(map[uint64]uint8)
	var tmp uint64
	for rows.Next() {
		err = rows.Scan(&tmp)
		if err != nil {
			return nil, err
		}
		ret[tmp] = 1
	}

	return ret, nil
}

*/