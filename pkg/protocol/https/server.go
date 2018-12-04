package https

import (
	"github.com/gin-gonic/gin"
	"os"
	"path"
	"runtime"
	"github.com/skyfour/go-model/pkg/api/v1/user"
	"github.com/skyfour/go-model/pkg/utils"
)

func RunServer(port string) error {
	defer utils.DB.Close()
	defer utils.Log.OSFile.Close()
	r := gin.Default()
	g := r.Group("/")
	g.Use()
	{
		g.GET("user/:uid",v1.GetUser)
		g.POST("user",v1.InsertUser)
		g.DELETE("user/:uid",v1.DeleteUser)
	}


	_, fileName, _, _ := runtime.Caller(0)
	rootPath := path.Join(fileName, "../../../../configs/")
	err := os.Chdir(rootPath)
	if err != nil {
		panic(err)
	}
	r.Run(":" + "8080")
	return err
}