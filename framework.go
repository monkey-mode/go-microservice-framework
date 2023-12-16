package framework

import (
	"os"
)

type Config struct {
	ConfigPath  string
	SecretPath  string
	MaskingPath string
}

func Init(cf Config) {
	_, err := os.Getwd()
	if err != nil {
		panic(err)
	}
}
