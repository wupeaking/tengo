package storage

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"

	"github.com/d5/tengo/v2/common"
)

type DB interface {
	Set(key string, value []byte) error
	Get(key string) ([]byte, error)
}

type Storage struct {
	Address string // 合约地址
	db      DB
}

var (
	NotFoundErr = errors.New("not found")
)

func New(contract string, db DB) (*Storage, error) {
	return &Storage{Address: contract, db: db}, nil
}

func (s *Storage) SaveConstants(consts []common.Object) error {
	baseKey := s.Address + "_constants_"
	for i := range consts {
		key := fmt.Sprintf("%s%d", baseKey, i)
		var buf bytes.Buffer
		enc := gob.NewEncoder(&buf)
		if err := enc.Encode(consts[i]); err != nil {
			return err
		}
		if err := s.db.Set(key, buf.Bytes()); err != nil {
			return err
		}
	}
	return nil
}

func (s *Storage) SaveByteCode(codes []byte) error {
	baseKey := s.Address + "_bytecode"
	return s.db.Set(baseKey, codes)
}

func (s *Storage) LoadByteCode() ([]byte, error) {
	baseKey := s.Address + "_bytecode"
	v, err := s.db.Get(baseKey)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (s *Storage) LoadConstants(index int) (common.Object, error) {
	baseKey := s.Address + "_constants_"
	key := fmt.Sprintf("%s%d", baseKey, index)
	v, err := s.db.Get(key)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewReader(v)
	dec := gob.NewDecoder(buf)
	var obj common.Object
	err = dec.Decode(obj)
	return obj, err
}

func (s *Storage) SetGlobal(index int, value common.Object) error {
	// todo::
	return nil
}

func (s *Storage) GetGlobal(index int) (common.Object, error) {
	// todo::
	return nil, nil
}
