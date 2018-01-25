package config

import (
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
	arkc := viper.GetStringMapString("ark")
	self.Ark = Ark{
		PublicKey:      arkc["publickey"],
		UserName:       arkc["username"],
		PassPhrasePath: arkc["passphrase"],
	}
	return self
}

func (self ViperConfig) SetEther() ViperConfig {
	ethc := viper.GetStringMapString("ether")
	self.Ether = Ether{
		KeyPath:        ethc["keyjsonpath"],
		DisplayMessage: ethc["displaymessage"],
	}
	return self
}
