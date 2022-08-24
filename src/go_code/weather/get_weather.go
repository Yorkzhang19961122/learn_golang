package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Weather struct {
	Status    string     `json:"status"返回状态`
	Count     string     `json:"count"返回结果总条数`
	Info      string     `json:"info"返回的状态信息`
	Infocode  string     `json:"infocode"返回状态说明`
	Forecasts []Forecast `json:"forecasts"预报天气信息数据`
}
type Forecast struct {
	City       string `json:"city"城市名称`
	Adcode     string `json:"adcode"城市编码`
	Province   string `json:"province"省份`
	Reporttime string `json:"reporttime"预报时间`
	Casts      []Cast `json:casts预报数据`
}
type Cast struct {
	Date         string `json:"date"日期`
	Week         string `json:"week"星期`
	Dayweather   string `json:"dayweather"白天天气`
	Nightweather string `json:"nightweather"晚上天气`
	Daytemp      string `json:"daytemp"白天温度`
	Nighttemp    string `json:"nighttemp"晚上温度`
	Daywind      string `json:"daywind"白天风向`
	Nightwind    string `json:"nightwind"晚上风向`
	Daypower     string `json:"daypower"白天风力`
	Nightpower   string `json:"nightpower"晚上风力`
}

//func verity(dayweather, nightweather string) string {
//	var sub string
//	rain := "雨"
//	snow := "雪"
//	sub = "收藏点赞投硬币，新的一天会有好事发生哝 ||今日天气预报"
//	if strings.Contains(dayweather, rain) || strings.Contains(nightweather, rain) {
//		sub = sub + "今天将降雨，出门请别忘带伞"
//	}
//	if strings.Contains(dayweather, snow) || strings.Contains(nightweather, snow) {
//		sub = sub + "    下雪了"
//	}
//	return sub
//}

//func NumToStr(str string) string {
//	switch str {
//	case "1":
//		return "一"
//	case "2":
//		return "二"
//	case "3":
//		return "三"
//	case "4":
//		return "四"
//	case "5":
//		return "五"
//	case "6":
//		return "六"
//	case "7":
//		return "日"
//	}
//	return ""
//}

func doHttpGetRequest(url string) (rlt string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		} else {
			return string(body), err
		}
	}
}

func getWeather() (string, error) {
	url := "https://restapi.amap.com/v3/weather/weatherInfo?"
	key := "67eb2f124ead53f6c51e88542f29d0a7"
	city := "330110"
	ext := "all"
	// rlt: 天气接口原始返回结果
	rlt, err := doHttpGetRequest(url + "city=" + city + "&key=" + key + "&extensions=" + ext)
	fmt.Println("rlt：", rlt)

	var data Weather
	var fore Forecast
	//var cast Cast
	// data: 按照Weather的type解析json
	json.Unmarshal([]byte(rlt), &data)
	if err != nil {
		return err.Error(), err
	} else {
		fmt.Println("data：", data)
		fore = data.Forecasts[0]
		fmt.Println("fore：", fore)

		output := "杭州市" + fore.City + ": " + "[今日]" + fore.Casts[0].Dayweather + "-" + fore.Casts[0].Nightweather + " " + fore.Casts[0].Nighttemp + "°C" + "-" + fore.Casts[0].Daytemp + "°C " +
			"[明日]" + fore.Casts[1].Dayweather + "-" + fore.Casts[1].Nightweather + " " + fore.Casts[1].Nighttemp + "°C" + "-" + fore.Casts[1].Daytemp + "°C " +
			"更新时间: " + fore.Reporttime
		fmt.Println("output：", output)

		//output := fore.Province + fore.City + " 预报时间：" + fore.Reporttime + "\n"
		//var str, subject string
		//for i := 0; i < len(fore.Casts); i++ {
		//	cast = fore.Casts[i]
		//	str += "日期：" + cast.Date + "\t星期" + NumToStr(cast.Week) +
		//		"\n白天：【天气：" + cast.Dayweather + "\t	温度：" + cast.Daytemp + "\t	风向：" + cast.Daywind + "\t	风力：" + cast.Daypower + "】" +
		//		"\n夜晚：【天气：" + cast.Nightweather + "\t	温度：" + cast.Nighttemp + "\t\t	风向：" + cast.Nightwind + "\t	风力：" + cast.Nightpower + "】\r\n"
		//}
		//subject = verity(fore.Casts[0].Dayweather, fore.Casts[0].Nightweather)
		//fmt.Println("天气：", output+str)
		return output, nil
	}
}

func main() {
	getWeather()
}
