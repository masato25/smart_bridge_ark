package config

import (
	"bytes"
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

type Database struct {
	Debug    bool
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

type Web struct {
	Port int
}

type Ark struct {
	PublicKey      string
	UserName       string
	PassPhrasePath string
}

type Ether struct {
	KeyPath        string
	DisplayMessage string
	KeyPassword    string
	RPCHost        string
	GasFee         int64
}

func (self Ark) GetPassString() (content string, err error) {
	log.Debugf("GetPassString - path: %s", self.PassPhrasePath)
	var file *os.File
	file, err = os.Open(self.PassPhrasePath)
	defer file.Close()
	if err != nil {
		return
	}
	bufc := bytes.NewBuffer(nil)
	_, err = io.Copy(bufc, file)
	if err != nil {
		return
	}
	content = string(bufc.Bytes())
	return
}

func (self Ether) GetKeyString() (content string, err error) {
	log.Debugf("GetKeyString - path: %s", self.KeyPath)
	var file *os.File
	file, err = os.Open(self.KeyPath)
	defer file.Close()
	if err != nil {
		return
	}
	bufc := bytes.NewBuffer(nil)
	_, err = io.Copy(bufc, file)
	if err != nil {
		return
	}
	content = string(bufc.Bytes())
	return
}
