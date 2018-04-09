/**
 *  author: lim
 *  data  : 18-3-24 下午4:54
 */

package xa

import (
	"github.com/lemonwx/xsql/config"
	"github.com/garyburd/redigo/redis"
)

// xsql support xa transaction by import Version for each data
type Version uint64
var Cfg *config.Conf
var pool redis.Pool

func NextVersion() (uint64, error) {
	/*
		1. get next version from redis,
		2. let nexy version += 1
		3. write [next version ]into redis's active versions list belong to this xsql id
		4. commit to redis
	*/
	return 0, nil
}

func VersionsInUse() ([]uint64, error) {
	return []uint64{1, 2, 3}, nil
}

func ReleaseVersionInfo(version uint64) error {
	/*
		1. remove the param version from redis's active versions list belong to this xsql id
		2. commit to redis
	*/
	return nil
}


func InitPool() {
	pool = redis.Pool{
		MaxIdle : 1,
		MaxActive: 1000,
		IdleTimeout: 10,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "127.0.0.1:6379")
			if err != nil {
				return nil, err
			}
			return c, nil
		},

	}
}