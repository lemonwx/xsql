package version
/**
 *  author: lim
 *  data  : 18-3-24 下午4:54
 */

 /*
package version

import (
	"strconv"

	"github.com/lemonwx/xsql/config"
	"github.com/garyburd/redigo/redis"
	"github.com/lemonwx/log"
)

// xsql support xa transaction by import Version for each data
type Version uint64
var Cfg *config.Conf
var pool redis.Pool


func NextVersion() ([]byte, error) {
	/*
		1. get next version from redis,
		2. let nexy version += 1
		3. write [next version ]into redis's active versions list belong to this xsql id
		4. commit to redis
	*/
	/*

	for {
		conn := pool.Get()
		defer conn.Close()

		// watch next_version
		reply, err := conn.Do("WATCH", "next_version")
		log.Debug(reply, err)
		if err != nil {
			return nil, err
		}

		//get next_version
		curV, err := conn.Do("get", "next_version")
		if err != nil {
			return nil, err
		}

		// convert to uint64
		curVersion, err := redis.Bytes(curV, err)
		if err != nil {
			return nil, err
		}

		nextV, err := strconv.ParseUint(string(curVersion), 10, 64)
		// add 1
		nextV += 1
		log.Debugf("get next v: %v", nextV)

		// next_version + 1 then add to VersionsInUse set
		err = conn.Send("MULTI")
		if err != nil {
			return nil, err
		}

		err = conn.Send("set", "next_version", nextV)
		if err != nil {
			log.Debug(err)
			return nil, err
		}

		err = conn.Send("sadd", "VersionsInUse", nextV)
		if err != nil {
			log.Debug(err)
			return nil, err
		}

		reply, err = conn.Do("EXEC")
		if err != nil {
			log.Debug(err)
		}

		if reply != nil {
			log.Debug("exec success ")
			return curVersion, nil
		}
	}
}


	/*
func VersionsInUse() ([][]byte, error) {
	conn := pool.Get()
	defer conn.Close()
	reply, err := conn.Do("SMEMBERS", "VersionsInUse")
	if err != nil {
		return nil, err
	}

	return redis.ByteSlices(reply, err)
}

func ReleaseVersion(version []byte) error {
	/*
		1. remove the param version from redis's active versions list belong to this xsql id
		2. commit to redis
	*/
	/*
	conn := pool.Get()
	defer  conn.Close()

	reply, err := conn.Do("SREM", "VersionsInUse", string(version))
	log.Debug(reply, err)
	return err
}

/*
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

*/