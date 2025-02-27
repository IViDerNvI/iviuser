package mysql

import (
	"sync"

	"github.com/ividernvi/iviuser/internal/apiserver/store"
	v1 "github.com/ividernvi/iviuser/model/v1"
	"github.com/ividernvi/iviuser/pkg/db"
	"gorm.io/gorm"
)

type datastore struct {
	db *gorm.DB
}

var _ store.Store = &datastore{}
var (
	Once     sync.Once
	MySQLIns store.Store
)

func GetMySQLInstanceOr(opts db.DBOptions) (store.Store, error) {
	if MySQLIns != nil {
		return MySQLIns, nil
	}

	var err error
	Once.Do(func() {
		var database *gorm.DB
		database, err = db.New(opts)
		if err != nil {
			panic(err)
		}
		DropAndMigrate(database)
		MySQLIns = &datastore{db: database}
	})
	return MySQLIns, err
}

func (d *datastore) Users() store.UserStore {
	return newUserStore(d.db)
}

func (d *datastore) Close() {
	sqlDB, err := d.db.DB()
	if err != nil {
		return
	}
	sqlDB.Close()
}

func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(&v1.User{})
	if err != nil {
		panic(err)
	}
}

func AutoDrop(db *gorm.DB) {
	err := db.Migrator().DropTable(&v1.User{})
	if err != nil {
		panic(err)
	}
}

func DropAndMigrate(db *gorm.DB) {
	AutoDrop(db)
	AutoMigrate(db)
}
