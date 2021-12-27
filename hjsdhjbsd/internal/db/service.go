package db

import (
	"context"
	"database/sql"
	"log"
	"sync"
	"time"

	// _ "github.com/go-sql-driver/mysql"
)

type Config struct {
	mu sync.Mutex
	ConnectString int
}
  
type ServiceDB struct {
	once sync.Once
	db *sql.DB
	config *Config
}

func NewService(cfg *Config) *ServiceDB {
	return &ServiceDB{
		config: cfg,
	}
}

func (s *ServiceDB) Exec(ctx context.Context) {
	for{
		select {
		case <-ctx.Done():
			log.Println("context done")
			return
		default:
			log.Println("!!!!!")
		}
		time.Sleep(time.Millisecond*30)
	}
}

//func (db *ServiceDB) DbMutex() (*sql.DB, error) {
//	db, err := db.getDB()
//	if err != nil {
//		return nil, err
//	}
//
//	err = db.Ping()
//	if err != nil {
//	   return nil, err
//	}
//
//	return db, nil
//}
//
//func (db *ServiceDB) getDB() (*sql.DB, error) {
//	lock.RLock()
//
//	if db != nil {
//		defer lock.RUnlock()
//
//		return db, nil
//	}
//
//	lock.RUnlock()
//
//	lock.Lock()
//  	defer lock.Unlock()
//
//	var err error
//
//	db, err = sql.Open("mysql", "root:test@tcp(127.0.0.1:3306)/test")
//	if err != nil {
//	   return nil, err
//	}
//
//	return db, nil
//}
