package gps

import (
	"fmt"
	"testing"
)

func Test_Distance(t *testing.T) {

	lat1 := 40.0439
	lng1 := 116.414

	lat2 := 40.0672
	lng2 := 116.415
	fmt.Println(Distance(lat1, lng1, lat2, lng2))
}

func Test_OutOfChina(t *testing.T) {

	lat1 := 40.0439
	lng1 := 116.414

	lat2 := 40.0672
	lng2 := 116.415

	if OutOfChina(lat1, lng1) == IsInChina(lat1, lng1) {
		panic("check in china error")
	}
	if OutOfChina(lat2, lng2) == IsInChina(lat2, lng2) {
		panic("check in china error")
	}
}

// GCJ02ToWGS84 火星转世界
func Test_gcj_decrypt(t *testing.T) {

	lat1 := 40.0439
	lng1 := 116.414

	lat2 := 40.0672
	lng2 := 116.415

	fmt.Println(GCJ02ToWGS84(lat1, lng1))
	fmt.Println(GCJ02ToWGS84Exact(lat1, lng1))

	fmt.Println(GCJ02ToWGS84(lat2, lng2))
	fmt.Println(GCJ02ToWGS84Exact(lat2, lng2))

}

// WGS84ToGCJ02 世界转火星
func Test_gcj_encrypt(t *testing.T) {

	lat1 := 40.0439
	lng1 := 116.414

	lat2 := 40.0672
	lng2 := 116.415

	fmt.Println(WGS84ToGCJ02(lat1, lng1))

	fmt.Println(WGS84ToGCJ02(lat2, lng2))

}

// BD09ToGCJ02 百度转火星
func Test_bd_encrypt(t *testing.T) {

	lat1 := 40.0439
	lng1 := 116.414

	lat2 := 40.0672
	lng2 := 116.415

	fmt.Println(BD09ToGCJ02(lat1, lng1))

	fmt.Println(BD09ToGCJ02(lat2, lng2))

}

// BD09ToWGS84 百度转火星
func Test_bj2_encrypt(t *testing.T) {

	lat1 := 40.0439
	lng1 := 116.414

	lat2 := 40.0672
	lng2 := 116.415

	fmt.Println(lat1, lng1)
	fmt.Println(GCJ02ToWGS84(BD09ToGCJ02(lat1, lng1)))

	fmt.Println(lat2, lng2)
	fmt.Println(GCJ02ToWGS84(BD09ToGCJ02(lat2, lng2)))

}
