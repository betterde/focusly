package database

import (
	"github.com/betterde/focusly/config"
	"github.com/surrealdb/surrealdb.go"
	"sync"
	"time"
)

var (
	db   *surrealdb.DB
	once sync.Once
)

// GetInstance Get SurrealDB instance.
func GetInstance() (*surrealdb.DB, error) {
	var err error

	once.Do(func() {
		db, err = surrealdb.New(config.Conf.Database.URL, surrealdb.WithTimeout(time.Second*30))
		if err != nil {
			return
		}

		_, err = db.Signin(map[string]interface{}{
			"user": config.Conf.Database.Username,
			"pass": config.Conf.Database.Password,
		})
		if err != nil {
			return
		}
	})

	return db, err
}
