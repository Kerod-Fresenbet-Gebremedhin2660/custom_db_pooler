package db

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"sync"
	"test_custom_db_pool/ds"
	"time"
)

type ConnDB struct {
	conn    *sql.DB
	acquire sync.Mutex
	timeout time.Duration
}

var (
	Conns ds.Stack[ConnDB]
	once  sync.Once
)

func (c *ConnDB) Acquire() (*sql.DB, error) {
	locked := c.acquire.TryLock()
	if locked == false {
		return nil, errors.New("conn is locked")
	} else {
		return c.conn, nil
	}
}

func (c *ConnDB) Release() {
	c.acquire.Unlock()
	Conns.Push(ConnDB{
		conn:    c.conn,
		acquire: sync.Mutex{},
	})
}

func InitConnPool() {
	var size uint64 = 150
	once.Do(func() {
		var i uint64 = 0
		Conns = ds.NewStack[ConnDB](size)
		for ; i < size; i++ {
			openConn, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/postgres")
			if err != nil {
				fmt.Println(fmt.Sprintf("ERROR InitConnPool: %s", err))
				continue
			}
			Conns.Push(ConnDB{
				conn:    openConn,
				acquire: sync.Mutex{},
			})
		}
	})
}

func AcquireFromConnPool() (*sql.DB, error) {
	if Conns.Size() == 0 {
		return nil, errors.New("pool not initialized / not populated")
	}
	for i := 0; i < Conns.Size(); i++ {
		connDB := Conns.Pop()
		conn, err := connDB.Acquire()
		if err != nil {
			return conn, err
		}
	}
	return nil, errors.New("all connections busy")
}
