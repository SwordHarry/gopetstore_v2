package persistence

import (
	"errors"
	"gopetstore_v2/src/domain"
	"gopetstore_v2/src/util"
)

const (
	getSequenceSQL    = `SELECT name, nextid FROM SEQUENCE WHERE NAME = ?`
	updateSequenceSQL = `UPDATE SEQUENCE SET NEXTID = ? WHERE NAME = ?`
)

func GetSequence(name string) (*domain.Sequence, error) {
	d, err := util.GetConnection()
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = d.Close()
	}()
	s := new(domain.Sequence)
	err = d.Get(s, getSequenceSQL, name)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func UpdateSequence(s *domain.Sequence) error {
	d, err := util.GetConnection()
	if err != nil {
		return err
	}
	defer func() {
		_ = d.Close()
	}()
	r, err := d.NamedExec(updateSequenceSQL, s)
	if err != nil {
		return err
	}
	row, err := r.RowsAffected()
	if err != nil {
		return err
	}
	if row > 0 {
		return nil
	}
	return errors.New("can not update sequence")
}
