package api

import (
	"echo-gorm-mvc/db"
	"echo-gorm-mvc/models"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

func HasOne(c echo.Context) error {
	res := make(map[string]interface{})
	req := make(map[string]interface{})

	c.Bind(&req)

	db := db.DbManager()
	db.LogMode(true)

	reqSells := []models.RequestSell{}

	db.Model(&models.RequestSell{}).Joins("JOIN tm_kabupaten on tm_kabupaten.KABUPATEN_ID = tr_jubel_request_sell.KABUPATEN_ID").Preload("Kabupaten").Where("tm_kabupaten.PROVINSI_ID = ?", req["provinsi_id"]).Find(&reqSells)

	res["req_sell"] = reqSells
	return c.JSON(http.StatusOK, res)
}

func GetMasterParameter(c echo.Context) error {
	res := make(map[string]interface{})

	masterParameter := []models.MasterParameter{}
	db := db.DbManager()
	db.LogMode(true)
	db.Model(&masterParameter).Find(&masterParameter)

	res["data"] = masterParameter
	return c.JSON(http.StatusOK, res)
}

func AddMasterParameter(c echo.Context) error {
	res := make(map[string]interface{})

	uniqID, _ := uuid.NewRandom()
	stringUniqID := uniqID.String()

	//method 1
	// req := make(map[string]interface{})
	// c.Bind(&req)
	// req["created_by"] = uniqID
	// reqEncoded, _ := json.Marshal(req)
	// db := db.DbManager()
	// db.Exec("ALTER TABLE tm_master_parameter AUTO_INCREMENT = 1;")
	// db.LogMode(true)
	// masterParameter := &models.MasterParameter{}
	// json.Unmarshal([]byte(reqEncoded), &masterParameter)
	// db.Create(&masterParameter)

	//method 2
	masterParameter := &models.MasterParameter{}
	c.Bind(&masterParameter)
	fmt.Println(uniqID)
	masterParameter.CreatedBy = &stringUniqID
	db := db.DbManager()
	db.Exec("ALTER TABLE tm_master_parameter AUTO_INCREMENT = 1;")
	db.LogMode(true)
	db.Create(masterParameter)

	res["status"] = 200

	return c.JSON(http.StatusOK, res)
}

func UpdateMasterParameter(c echo.Context) error {
	res := make(map[string]interface{})
	req := make(map[string]interface{})

	c.Bind(&req)

	db := db.DbManager()
	db.LogMode(true)
	masterParameter := models.MasterParameter{}

	//method 1
	// db.Model(&models.MasterParameter{}).Where("parameter_id = ?", req["parameter_id"]).Find(&masterParameter)
	// masterParameter.ParameterLabel = req["parameter_label"].(string)
	// masterParameter.ParameterDescription = req["parameter_description"].(string)
	// db.Save(&masterParameter)

	//method 2
	db.Model(&masterParameter).Where("parameter_id = ?", c.Param("id")).Update(req)

	res["status"] = 200
	return c.JSON(http.StatusOK, res)
}

func DeleteMasterParameter(c echo.Context) error {
	res := make(map[string]interface{})

	db := db.DbManager()
	db.LogMode(true)
	db.Where("parameter_id = ?", c.Param("id")).Delete(&models.MasterParameter{}, "test")

	res["status"] = 200
	return c.JSON(http.StatusOK, res)
}
