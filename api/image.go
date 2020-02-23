package api

import (
	"echo-gorm-mvc/config"
	"net/http"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/labstack/echo"
)

func PutS3(c echo.Context) error {
	configuration := config.GetConfig()
	res := make(map[string]interface{})
	req := make(map[string]interface{})
	c.Bind(&req)

	session, err := session.NewSession(&aws.Config{
		Region: aws.String(configuration.AWS_REGION),
	})
	if err != nil {
		println(err)
	}

	// Upload Files
	now := time.Now()
	sec := now.Unix()
	filename := "golang/uniqid" + strconv.FormatInt(sec, 10) + ".jpg"
	base64Req := req["base64"].(string)
	err = UploadFileS3(session, "sumo365-s3-dev-bucket", filename, base64Req)
	if err != nil {
		res["status"] = 500
		res["message"] = "UploadFileS3 return err"
		WriteLog(err)
		return c.JSON(http.StatusOK, res)
	}

	res["status"] = "200"
	res["message"] = "success"

	return c.JSON(http.StatusOK, res)
}
