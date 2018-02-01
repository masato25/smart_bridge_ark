package config

import (
	"strconv"

	"github.com/spf13/viper"
)

type ViperConfig struct {
	Database Database
	Web      Web
	Ark      Ark
	Ether    Ether
}

func (self ViperConfig) SetAll() ViperConfig {
	self = self.SetDatabase()
	self = self.SetWeb()
	self = self.SetArk()
	self = self.SetEther()
	return self
}

func (self ViperConfig) SetDatabase() ViperConfig {
	dbc := viper.GetStringMap("database")
	self.Database = Database{
		Debug:    dbc["debug"].(bool),
		Host:     dbc["host"].(string),
		Port:     dbc["port"].(int),
		User:     dbc["user"].(string),
		Password: dbc["password"].(string),
		DBName:   dbc["dbname"].(string),
	}
	return self
}

func (self ViperConfig) SetWeb() ViperConfig {
	webc := viper.GetStringMap("web")
	self.Web = Web{
		Port: webc["port"].(int),
	}
	return self
}

func (self ViperConfig) SetArk() ViperConfig {
	arkc := viper.GetStringMap("ark")
	self.Ark = Ark{
		Enable:         arkc["enable"].(bool),
		Host:           arkc["host"].(string),
		DevNet:         arkc["dev_net"].(string),
		PublicKey:      arkc["publickey"].(string),
		UserName:       arkc["username"].(string),
		PassPhrasePath: arkc["passphrase"].(string),
		Address:        arkc["address"].(string),
	}
	return self
}

func (self ViperConfig) SetEther() ViperConfig {
	ethc := viper.GetStringMapString("ether")
	gasfee, _ := strconv.Atoi(ethc["gasfee"])
	self.Ether = Ether{
		Enable:         ethc["enable"] == "true",
		Address:        ethc["address"],
		KeyPath:        ethc["keyjsonpath"],
		DisplayMessage: ethc["displaymessage"],
		KeyPassword:    ethc["keypassword"],
		RPCHost:        ethc["ethereum_rpc_hist"],
		GasFee:         int64(gasfee),
	}
	return self
}
