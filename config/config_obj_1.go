package config

import "os"

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
}

func (self Ark) GetPassString() (content string, err error) {
	var file *os.File
	file, err = os.Open(self.PassPhrasePath)
	if err != nil {
		return
	}
	var contentb []byte
	_, err = file.Read(contentb)
	if err != nil {
		return
	}
	content = string(contentb)
	return
}

func (self Ether) GetKeyString() (content string, err error) {
	var file *os.File
	file, err = os.Open(self.KeyPath)
	if err != nil {
		return
	}
	var contentb []byte
	_, err = file.Read(contentb)
	if err != nil {
		return
	}
	content = string(contentb)
	return
}
