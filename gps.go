package gps

import (
	"math"
)

// x_pi ...
const x_pi = math.Pi * 3000.0 / 180.0

func Delta(lat, lon float64) (float64, float64) {
	var a = 6378245.0               //  a: 卫星椭球坐标投影到平面地图坐标系的投影因子。
	var ee = 0.00669342162296594323 //  ee: 椭球的偏心率。
	var dLat = TransformLat(lon-105.0, lat-35.0)
	var dLon = TransformLon(lon-105.0, lat-35.0)
	var radLat = lat / 180.0 * math.Pi
	var magic = math.Sin(radLat)
	magic = 1 - ee*magic*magic
	var sqrtMagic = math.Sqrt(magic)
	dLat = (dLat * 180.0) / ((a * (1 - ee)) / (magic * sqrtMagic) * math.Pi)
	dLon = (dLon * 180.0) / (a / sqrtMagic * math.Cos(radLat) * math.Pi)
	return dLat, dLon
}

//WGS-84 to GCJ-02
func WGS84ToGCJ02(wgsLat, wgsLon float64) (float64, float64) {
	//if OutOfChina(wgsLat, wgsLon)
	if !IsInChina(wgsLat, wgsLon) {
		return wgsLat, wgsLon
	}

	var lat, lon = Delta(wgsLat, wgsLon)

	return wgsLat + lat, wgsLon + lon
}

//GCJ-02 to WGS-84
func GCJ02ToWGS84(gcjLat, gcjLon float64) (float64, float64) {

	//if (this.outOfChina(gcjLat, gcjLon))
	if !IsInChina(gcjLat, gcjLon) {
		return gcjLat, gcjLon
	}
	var lat, lon = Delta(gcjLat, gcjLon)
	return gcjLat - lat, gcjLon - lon
}

//GCJ-02 to WGS-84 exactly
func GCJ02ToWGS84Exact(gcjLat, gcjLon float64) (float64, float64) {
	initDelta := 0.01
	threshold := 0.000000001
	dLat := initDelta
	dLon := initDelta
	mLat := gcjLat - dLat
	mLon := gcjLon - dLon
	pLat := gcjLat + dLat
	pLon := gcjLon + dLon
	var i int
	var wgsLat, wgsLon, tmplat, tmplon float64
	for i < 10000 {
		wgsLat = (mLat + pLat) / 2
		wgsLon = (mLon + pLon) / 2
		tmplat, tmplon = WGS84ToGCJ02(wgsLat, wgsLon)
		dLat = tmplat - gcjLat
		dLon = tmplon - gcjLon
		if (math.Abs(dLat) < threshold) && (math.Abs(dLon) < threshold) {
			i = 999999
		}

		if dLat > 0 {
			pLat = wgsLat
		} else {
			mLat = wgsLat
		}

		if dLon > 0 {
			pLon = wgsLon
		} else {
			mLon = wgsLon
		}
		i++
	}
	return wgsLat, wgsLon
}

//GCJ-02 to BD-09
func GCJ02ToBD09(gcjLat, gcjLon float64) (float64, float64) {
	var x, y = gcjLon, gcjLat
	var z = math.Sqrt(x*x+y*y) + 0.00002*math.Sin(y*x_pi)
	var theta = math.Atan2(y, x) + 0.000003*math.Cos(x*x_pi)
	bdLon := z*math.Cos(theta) + 0.0065
	bdLat := z*math.Sin(theta) + 0.006
	return bdLat, bdLon
}

//BD-09 to GCJ-02
func BD09ToGCJ02(bdLat, bdLon float64) (float64, float64) {
	x := bdLon - 0.0065
	y := bdLat - 0.006
	var z = math.Sqrt(x*x+y*y) - 0.00002*math.Sin(y*x_pi)
	var theta = math.Atan2(y, x) - 0.000003*math.Cos(x*x_pi)
	var gcjLon = z * math.Cos(theta)
	var gcjLat = z * math.Sin(theta)
	return gcjLat, gcjLon
}

//WGS-84 to Web mercator
//mercatorLat -> y mercatorLon -> x
func MercatorEncrypt(wgsLat, wgsLon float64) (float64, float64) {
	var x = wgsLon * 20037508.34 / 180.0
	var y = math.Log(math.Tan((90.+wgsLat)*math.Pi/360.)) / (math.Pi / 180.)
	y = y * 20037508.34 / 180.
	return y, x

	/*
		if math.Abs(wgsLon) > 180 || math.Abs(wgsLat) > 90 {
			return 0.0, 0.0
		}

		var x = 6378137.0 * wgsLon * 0.017453292519943295
		var a = wgsLat * 0.017453292519943295
		var y = 3189068.5 * math.Log((1.0+math.Sin(a))/(1.0-math.Sin(a)))
		return y, x
	*/
	//
}

// Web mercator to WGS-84
// mercatorLat -> y mercatorLon -> x
func MercatorDecrypt(mercatorLat, mercatorLon float64) (float64, float64) {
	var x = mercatorLon / 20037508.34 * 180.
	var y = mercatorLat / 20037508.34 * 180.
	y = 180 / math.Pi * (2*math.Atan(math.Exp(y*math.Pi/180.)) - math.Pi/2)
	return y, x
	/*
		if math.Abs(mercatorLon) < 180 && math.Abs(mercatorLat) < 90 {
			return 0.0, 0.0
		}
		if (math.Abs(mercatorLon) > 20037508.3427892) || (math.Abs(mercatorLat) > 20037508.3427892) {
			return 0.0, 0.0
		}
		var a = mercatorLon / 6378137.0 * 57.295779513082323
		var x = a - (math.Floor(((a + 180.0) / 360.0)) * 360.0)
		var y = (1.5707963267948966 - (2.0 * math.Atan(math.Exp((-1.0*mercatorLat)/6378137.0)))) * 57.295779513082323
		return y, x
		//*/
}

// Distance 计算距离
func Distance(latA, lonA, latB, lonB float64) float64 {
	radius := 6371000.0 // 6378137
	rad := math.Pi / 180.0

	latA = latA * rad
	lonA = lonA * rad
	latB = latB * rad
	lonB = lonB * rad

	theta := lonB - lonA
	dist := math.Acos(math.Sin(latA)*math.Sin(latB) + math.Cos(latA)*math.Cos(latB)*math.Cos(theta))

	return dist * radius
}

// Rects 距形集合
type Rects []Rect

// Rect 距形结构
type Rect struct {
	west  float64
	north float64
	east  float64
	south float64
}

// Rectangle 获取一个距形范围
func Rectangle(lng1, lat1, lng2, lat2 float64) Rect {
	var rect Rect
	rect.west = math.Min(lng1, lng2)
	rect.north = math.Max(lat1, lat2)
	rect.east = math.Max(lng1, lng2)
	rect.south = math.Min(lat1, lat2)
	return rect
}

// IsInRect 点在距形内
func IsInRect(rect Rect, lon, lat float64) bool {
	return rect.west <= lon && rect.east >= lon && rect.north >= lat && rect.south <= lat
}

// IsInChina 在中国内(粗略)
func IsInChina(lat, lon float64) bool {
	//China region - raw data
	//http://www.cnblogs.com/Aimeast/archive/2012/08/09/2629614.html
	var region = Rects{
		Rectangle(79.446200, 49.220400, 96.330000, 42.889900),
		Rectangle(109.687200, 54.141500, 135.000200, 39.374200),
		Rectangle(73.124600, 42.889900, 124.143255, 29.529700),
		Rectangle(82.968400, 29.529700, 97.035200, 26.718600),
		Rectangle(97.025300, 29.529700, 124.367395, 20.414096),
		Rectangle(107.975793, 20.414096, 111.744104, 17.871542),
	}

	//China excluded region - raw data
	var exclude = Rects{
		Rectangle(119.921265, 25.398623, 122.497559, 21.785006),
		Rectangle(101.865200, 22.284000, 106.665000, 20.098800),
		Rectangle(106.452500, 21.542200, 108.051000, 20.487800),
		Rectangle(109.032300, 55.817500, 119.127000, 50.325700),
		Rectangle(127.456800, 55.817500, 137.022700, 49.557400),
		Rectangle(131.266200, 44.892200, 137.022700, 42.569200),
	}

	for _, r := range region {
		if IsInRect(r, lon, lat) {
			for _, e := range exclude {
				if IsInRect(e, lon, lat) {
					return false
				}
			}
			return true
		}
	}

	return false
}

// OutOfChina 中国范围外
func OutOfChina(lat, lon float64) bool {

	if lon < 72.004 || lon > 137.8347 {
		return true
	}
	if lat < 0.8293 || lat > 55.8271 {
		return true
	}
	return false
}

// TransformLat ...
func TransformLat(x, y float64) float64 {
	var ret = -100.0 + 2.0*x + 3.0*y + 0.2*y*y + 0.1*x*y + 0.2*math.Sqrt(math.Abs(x))
	ret += (20.0*math.Sin(6.0*x*math.Pi) + 20.0*math.Sin(2.0*x*math.Pi)) * 2.0 / 3.0
	ret += (20.0*math.Sin(y*math.Pi) + 40.0*math.Sin(y/3.0*math.Pi)) * 2.0 / 3.0
	ret += (160.0*math.Sin(y/12.0*math.Pi) + 320*math.Sin(y*math.Pi/30.0)) * 2.0 / 3.0
	return ret
}

// TransformLon ...
func TransformLon(x, y float64) float64 {
	var ret = 300.0 + x + 2.0*y + 0.1*x*x + 0.1*x*y + 0.1*math.Sqrt(math.Abs(x))
	ret += (20.0*math.Sin(6.0*x*math.Pi) + 20.0*math.Sin(2.0*x*math.Pi)) * 2.0 / 3.0
	ret += (20.0*math.Sin(x*math.Pi) + 40.0*math.Sin(x/3.0*math.Pi)) * 2.0 / 3.0
	ret += (150.0*math.Sin(x/12.0*math.Pi) + 300.0*math.Sin(x/30.0*math.Pi)) * 2.0 / 3.0
	return ret
}
