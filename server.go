package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/", homepage)

	http.HandleFunc("/hello", hello)

	http.HandleFunc("/weather/", weather)

	http.HandleFunc("/city/", city)

	fmt.Printf("Server running%s\n", "...")
	println("Server running...")

	http.ListenAndServe(":8008", nil)

}

func homepage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("I am Go server"))
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello!"))
}

func weather(w http.ResponseWriter, r *http.Request) {
	id := strings.SplitN(r.URL.Path, "/", 3)[2]
	println(r.URL.Path)
	println("id : " + id)

	data, err := query(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(data)
}

func query(id string) (weatherData, error) {
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/forecast?APPID=46025ed068eca3d54d729dc0e2024b32&id=" + id)
	if err != nil {
		return weatherData{}, err
	}

	defer resp.Body.Close()

	var d weatherData

	println(json.NewDecoder(resp.Body))

	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return weatherData{}, err
	}

	return d, nil
}

func city(w http.ResponseWriter, r *http.Request) {
	var city_id int
	var city_country string
	city_name := strings.SplitN(r.URL.Path, "/", 3)[2]

	// 读取JSON文件内容 返回字节切片
	jsondata, err := ioutil.ReadFile("./city.list.json")
	if err != nil {
		return
	}

	// 打印时需要转为字符串
	// fmt.Println("*** data.json content: ***")
	// fmt.Println(string(data))

	// 将字节切片映射到指定结构上
	var citydata cityData
	json.Unmarshal(jsondata, &citydata)

	// 遍历json
	for k, v := range citydata {
		if strings.ToLower(v.Name) == strings.ToLower(city_name) {
			fmt.Println(k, v.Name)
			city_id = v.ID
			city_country = v.Country
			break
		}
	}

	// 打印对象结构
	// fmt.Println("*** unmarshal result: ***")
	// fmt.Println(citydata)

	// 打印格式化JSON结果
	// newBytes, _ := json.Marshal(&citydata)
	// newBytes, _ := json.MarshalIndent(&citydata, "", "  ") // 格式化输出
	// fmt.Println("*** update content: ***")
	// fmt.Println(string(newBytes))

	// 输出结果
	// arr := []string{"hello", "apple", "python", "golang", "base", "peach", "pear"}
	arr := map[string]string{"city": city_name, "id": strconv.Itoa(city_id), "country": city_country, "api": fmt.Sprintf("/weather/%d", city_id)}
	lang, err := json.Marshal(arr)
	if err == nil {
		fmt.Println(string(lang))
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Write([]byte(string(lang)))
		return
	}

	w.Write([]byte("city " + city_name + " id " + strconv.Itoa(city_id)))
}

type weatherData struct {
	Cod     string  `json:"cod"`
	Message float64 `json:"message"`
	Cnt     int     `json:"cnt"`
	List    []struct {
		Dt   int `json:"dt"`
		Main struct {
			Temp      float64 `json:"temp"`
			TempMin   float64 `json:"temp_min"`
			TempMax   float64 `json:"temp_max"`
			Pressure  float64 `json:"pressure"`
			SeaLevel  float64 `json:"sea_level"`
			GrndLevel float64 `json:"grnd_level"`
			Humidity  int     `json:"humidity"`
			TempKf    float64 `json:"temp_kf"`
		} `json:"main"`
		Weather []struct {
			ID          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
		Clouds struct {
			All int `json:"all"`
		} `json:"clouds"`
		Wind struct {
			Speed float64 `json:"speed"`
			Deg   float64 `json:"deg"`
		} `json:"wind"`
		Sys struct {
			Pod string `json:"pod"`
		} `json:"sys"`
		DtTxt string `json:"dt_txt"`
		Rain  struct {
		} `json:"rain,omitempty"`
		Snow struct {
			ThreeH float64 `json:"3h"`
		} `json:"snow,omitempty"`
	} `json:"list"`
	City struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Coord struct {
			Lat float64 `json:"lat"`
			Lon float64 `json:"lon"`
		} `json:"coord"`
		Country  string `json:"country"`
		Timezone int    `json:"timezone"`
		Sunrise  int    `json:"sunrise"`
		Sunset   int    `json:"sunset"`
	} `json:"city"`
}

type cityData []struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
	Coord   struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
}
