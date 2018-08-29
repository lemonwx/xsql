/**
 *  author: lim
 *  data  : 18-8-28 下午10:43
 */

package server

import "time"

type field struct {
	t time.Duration
	c int64
}

func (f *field) avg() time.Duration {
	if f.c != 0 {
		return time.Duration(int64(f.t) / f.c)
	}
	return time.Duration(0)
}

func (f *field) add(t time.Duration) {
	f.c += 1
	f.t += t
}

type Stat struct {
	SqlparseT *field
	RouteT    *field
	VersionT  *field
	ExecT     *field
	ChkInuseT *field
	ClearT    *field
	GetConn   *field
	PutConn   *field
	Dispatch  *field
}

func newStat() *Stat {
	return &Stat{
		SqlparseT: &field{},
		RouteT:    &field{},
		VersionT:  &field{},
		ExecT:     &field{},
		ChkInuseT: &field{},
		ClearT:    &field{},
		GetConn:   &field{},
		PutConn:   &field{},
		Dispatch:  &field{},
	}
}

func (s *Stat) getTheoryAvg() time.Duration {
	max := time.Duration(0)
	vt := s.VersionT.avg()
	et := s.ExecT.avg()
	if vt > et {
		max = vt
	} else {
		max = et
	}

	return s.SqlparseT.avg() + s.RouteT.avg() + max + s.ChkInuseT.avg() + s.ClearT.avg()
}
