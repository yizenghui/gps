package gps

import (
	"fmt"
	"testing"
)

func Test_Demo(t *testing.T) {
	// 百度地图广州南站
	latB1 := 22.995057
	lngB1 := 113.275872
	// 百度转火星
	latG1, lngG1 := BD09ToGCJ02(latB1, lngB1)
	// 火星转国际
	latW1, lngW1 := GCJ02ToWGS84(latG1, lngG1)

	// 百度地图广州站
	latB2 := 23.155005
	lngB2 := 113.264057
	// 百度转火星
	latG2, lngG2 := BD09ToGCJ02(latB2, lngB2)
	// 火星转国际
	latW2, lngW2 := GCJ02ToWGS84(latG2, lngG2)

	fmt.Println("百度坐标(BD-09)上广州南站到广州广州站距离：")
	fmt.Println(Distance(latB1, lngB1, latB2, lngB2), latB1, lngB1, latB2, lngB2)
	fmt.Println("腾讯高德坐标(GCJ-02)上广州南站到广州广州站距离：")
	fmt.Println(Distance(latG1, lngG1, latG2, lngG2), latG1, lngG1, latG2, lngG2)
	fmt.Println("国际坐标(WGS-84)地图上广州南站到广州广州站距离：")
	fmt.Println(Distance(latW1, lngW1, latW2, lngW2), latW1, lngW1, latW2, lngW2)
}

func Test_BDToGCJToWGS(t *testing.T) {
	// 百度地图广州站
	latB1 := 23.155005
	lngB1 := 113.264057
	// 百度转火星
	fmt.Println("BD09:", latB1, lngB1)
	latG1, lngG1 := BD09ToGCJ02(latB1, lngB1)
	fmt.Println("BD09ToGCJ02:", latG1, lngG1)
	latB1, lngB1 = GCJ02ToBD09(latG1, lngG1)
	fmt.Println("GCJ02ToBD09:", latB1, lngB1)
	// 火星转国际
	latW1, lngW1 := GCJ02ToWGS84(latG1, lngG1)
	fmt.Println("GCJ02ToWGS84:", latW1, lngW1)
	//
	latG1, lngG1 = WGS84ToGCJ02(latW1, lngW1)
	fmt.Println("WGS84ToGCJ02:", latG1, lngG1)

}
