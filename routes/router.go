package routes

import (
	"echo-gorm-mvc/api"

	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", api.Home)

	e.GET("/token-with-loop", api.TokenWithLoop)
	e.POST("/token-with-body", api.TokenWithBody)
	e.GET("/token-with-query-string", api.TokenWithQueryString)
	e.GET("/token-with-path-parameter/:email", api.TokenWithPathParameter)

	e.GET("/go-routine", api.GoRoutine)

	e.POST("/put-s3", api.PutS3)

	e.GET("/http-req", api.HttpReq)

	e.POST("/has-one", api.HasOne)
	e.GET("/master-parameter", api.GetMasterParameter)
	e.PUT("/master-parameter", api.AddMasterParameter)
	e.PATCH("/master-parameter/:id", api.UpdateMasterParameter)
	e.DELETE("/master-parameter/:id", api.DeleteMasterParameter)

	return e
}
