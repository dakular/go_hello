package main

import "github.com/gin-gonic/gin"
import "net/http"
import "fmt"
import "io/ioutil"
import "encoding/json"
import "log"
import "github.com/justinas/nosurf"

type Data struct {
	Datetime    string  `json:"datetime"`
	Pm25        float64 `json:"pm25"`
	Pm10        float64 `json:"pm10"`
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
}

func main() {
	router := gin.Default()
	router.Static("/static", "./static") // 注册静态资源目录
	router.LoadHTMLGlob("templates/*")   // 注册模板目录

	csrf := nosurf.New(router)
	csrf.SetFailureHandler(http.HandlerFunc(csrfFailHandler))

	router.GET("/", func(c *gin.Context) {
		// c.String(http.StatusOK, "Gin Home")
		c.HTML(200, "index.tmpl", gin.H{
			"title":      "Duckula + Gin",
			"content":    "Gin Home",
			"csrf_token": nosurf.Token(c.Request),
		})
	})

	router.GET("/ping", func(c *gin.Context) {
		data := getPM()
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "pong",
			"data": gin.H{
				"datetime": data["datetime"],
				"pm25":     data["pm25"],
			},
		})
	}, middleware1)

	user := router.Group("/user")

	user.GET("/:name", func(c *gin.Context) {
		// 使用 c.Param(key) 获取 url 参数
		log.Println("CSRF TOKEN: " + nosurf.Token(c.Request))
		name := c.Param("name")
		date := c.DefaultQuery("date", "2000-01-01")
		c.String(http.StatusOK, "Hello %s %s", name, date)
	}, middleware1)

	user.POST("/create", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.DefaultPostForm("password", "000000")
		c.String(http.StatusOK, "username %s password %s", username, password)
	}, middleware1)

	// router.Run(":8080") // listen and serve on 0.0.0.0:8080
	http.ListenAndServe(":8080", csrf)
}

func middleware1(c *gin.Context) {
	log.Println("exec middleware1")
	//你可以写一些逻辑代码
	// 执行该中间件之后的逻辑
	c.Next()
}

func csrfFailHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "[XSRF Error] %s\n", nosurf.Reason(r))
	log.Println("[XSRF Error] %s\n", nosurf.Reason(r))
}

func getPM() map[string]interface{} {
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
	return data2
}
