package api

import (
	"echo-gorm-mvc/db"
	"echo-gorm-mvc/models"
	"fmt"
	_ "io/ioutil"
	"net/http"
	"strings"
	"time"
	"github.com/labstack/echo"
)

func Home(c echo.Context) error {

	return c.String(http.StatusOK, "Welcome Echo.labstack")

}

func TokenWithLoop(c echo.Context) error {
	db := db.DbManager()
	tokens := []models.Token{}
	db.Where("email like ?", "%johanes%").Find(&tokens)

	res := make(map[string]interface{})

	for index, row := range tokens {
		t, err := time.Parse("2006-01-02 15:04:05", row.CreatedAt.String())
		if err != nil {
			res["status"] = "failed"
			res["message"] = err.Error()
			return c.JSON(http.StatusOK, res)
		}

		tokens[index].Formatted = t.Format("Mon Jan _2 15:04:05 2006")
	}
	res["status"] = "200"
	res["message"] = "success"
	res["tokens"] = tokens

	return c.JSON(http.StatusOK, res)
}

func TokenWithBody(c echo.Context) error {
	res := make(map[string]interface{})
	req := make(map[string]interface{})

	c.Bind(&req)

	/*if you want to handle error
	if err != nil {
	 	res["status"] = "failed"
	 	res["message"] = err.Error()
	 	return c.JSON(http.StatusOK, res)
	}*/

	res["status"] = "200"
	res["message"] = "success"
	res["request"] = map[string]interface{}{"a": req["a"], "b": req["b"]}

	return c.JSON(http.StatusOK, res)
}

func TokenWithQueryString(c echo.Context) error {
	res := make(map[string]interface{})

	res["status"] = "200"
	res["message"] = "success"
	res["query_param"] = c.QueryParam("email")

	return c.JSON(http.StatusOK, res)
}

func TokenWithPathParameter(c echo.Context) error {
	res := make(map[string]interface{})

	res["status"] = "200"
	res["message"] = "success"
	res["path_param"] = c.Param("email")

	return c.JSON(http.StatusOK, res)
}

func GoRoutine(c echo.Context) error {
	res := make(map[string]interface{})

	go background()

	for i := 0; i < 5; i++ {
		fmt.Printf("going process %v \n", i)
	}

	res["status"] = "200"
	res["message"] = "success"

	return c.JSON(http.StatusOK, res)
}

func background() {
	time.Sleep(5 * time.Second)

	for i := 0; i < 5; i++ {
		fmt.Printf("background process %v \n", i)
	}
}

func HttpReq(c echo.Context) error {
	res := make(map[string]interface{})

	url := "http://mockbin.com/request?foo=bar&foo=baz"
	payload := strings.NewReader("{\"foo\": \"bar\"}")
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("cookie", "foo=bar; bar=baz")
	req.Header.Add("x-pretty-print", "2")

	resHttp, _ := http.DefaultClient.Do(req)

	defer resHttp.Body.Close()

	// body, _ := ioutil.ReadAll(resHttp.Body)

	res["status"] = "200"
	res["message"] = "success"
	return c.JSON(http.StatusOK, resHttp)
}
