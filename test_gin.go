package main

import "github.com/gin-gonic/gin"
import "net/http"
import "fmt"
import "io/ioutil"
import "encoding/json"

type Data struct {
	Datetime    string  `json:"datetime"`
	Pm25        float64 `json:"pm25"`
	Pm10        float64 `json:"pm10"`
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
}

func main() {
	api := "http://duckula.net:66/env"
	data := Data{}
	var data2 map[string]interface{}
	resp, _ := http.Get(api)
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	json.Unmarshal([]byte(body), &data)
	json.Unmarshal(body, &data2)
	fmt.Println("===========================")
	fmt.Println(string(body))
	fmt.Println(fmt.Sprintf("%+v", data))
	fmt.Println(data2["datetime"], "PM2.5:", data2["pm25"], "PM10:", data2["pm10"])
	fmt.Println("===========================")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":    200,
			"message": "pong",
			"data": gin.H{
				"datetime": data2["datetime"],
				"pm25":     data2["pm25"],
			},
		})
	})
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
