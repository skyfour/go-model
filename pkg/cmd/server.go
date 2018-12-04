package cmd

import (
	"github.com/skyfour/go-model/pkg/protocol/https"
	"github.com/skyfour/go-model/pkg/utils"
)

func RunServer() error {
	// get configuration

	port := utils.Config.GetString("gin.port")

	return https.RunServer(port)
}
