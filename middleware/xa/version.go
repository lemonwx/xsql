/**
 *  author: lim
 *  data  : 18-3-24 下午4:54
 */

package xa

// xsql support xa transaction by import Version for each data
type Version uint64

func NextVersion() (uint64, error) {
	/*
		1. get next version from redis,
		2. let nexy version += 1
		3. write [next version ]into redis's active versions list belong to this xsql id
		4. commit to redis
	*/
	return 0, nil
}

func ActiveVersions() ([]uint64, error) {
	return []uint64{1, 2, 3}, nil
}

func ReleaseVersionInfo(version uint64) error {
	/*
		1. remove the param version from redis's active versions list belong to this xsql id
		2. commit to redis
	*/
	return nil
}
