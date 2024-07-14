package main

import (
	"github.com/GabrielBruno7/gopportunities/config"
	"github.com/GabrielBruno7/gopportunities/router"
)


func main() {
	err := config.Init();

	if err != nil {
		println(err)
		return
	}

	router.Initialize();
}

//gorm
//gin-gonic
//handler
//ponteiros
//go mod tidy //Reset the packs and install 