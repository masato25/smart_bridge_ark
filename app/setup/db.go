package setup

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/masato25/smart_bridge_ark/app/model/ark"
	"github.com/masato25/smart_bridge_ark/app/model/blocks"
	"github.com/masato25/smart_bridge_ark/app/model/connections"
	"github.com/masato25/smart_bridge_ark/app/model/delegate"
	"github.com/masato25/smart_bridge_ark/app/model/ether"
	"github.com/masato25/smart_bridge_ark/config"
	log "github.com/sirupsen/logrus"
)

var db *gorm.DB

func GetConn() *gorm.DB {
	return db
}

func ConnDB(migrate ...bool) (err error) {
	dbconf := config.MyConfig().Database
	connpath := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s sslmode=disable password=%s",
		dbconf.Host, dbconf.Port, dbconf.User, dbconf.DBName, dbconf.Password)
	db, err = gorm.Open("postgres", connpath)
	if err != nil {
		log.Error(err)
		return
	}
	if dbconf.Debug {
		db.LogMode(true)
	}
	if len(migrate) != 0 {
		log.Debug(migrate[0])
		if migrate[0] {
			Migration()
		}
	}
	db.Model(&connections.ArkEther{}).Related(&ark.ArkTransaction{}).Related(&ether.EtherTransaction{})
	db.Model(&delegate.VoteProfit{}).Related(&delegate.Vote{}).Related(&blocks.Block{})
	return
}

func Migration() {
	db.DropTableIfExists(&ark.ArkTransaction{})
	db.DropTableIfExists(&ether.EtherTransaction{})
	db.DropTableIfExists(&connections.ArkEther{})
	db.DropTableIfExists(&blocks.Block{})
	db.DropTableIfExists(&delegate.Vote{})
	db.DropTableIfExists(&delegate.VoteProfit{})
	db.AutoMigrate(
		ark.ArkTransaction{},
		ether.EtherTransaction{},
		connections.ArkEther{},
		blocks.Block{},
		delegate.Vote{},
		delegate.VoteProfit{},
	)
}

func CloseDB() error {
	return db.Close()
}
