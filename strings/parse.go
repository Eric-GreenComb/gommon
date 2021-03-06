package strings

import ()

// ID2Province id 2 chinese province
func ID2Province(id string) string {
	switch id {
	case "11":
		return "北京市"
	case "12":
		return "天津市"
	case "13":
		return "河北省"
	case "14":
		return "山西省"
	case "15":
		return "内蒙古自治区"

	case "21":
		return "辽宁省"
	case "22":
		return "吉林省"
	case "23":
		return "黑龙江省"

	case "31":
		return "上海市"
	case "32":
		return "江苏省"
	case "33":
		return "浙江省"
	case "34":
		return "安徽省"
	case "35":
		return "福建省"
	case "36":
		return "江西省"
	case "37":
		return "山东省"

	case "41":
		return "河南省"
	case "42":
		return "湖北省"
	case "43":
		return "湖南省"
	case "44":
		return "广东省"
	case "45":
		return "广西壮族自治区"
	case "46":
		return "海南省"

	case "51":
		return "重庆市"
	case "52":
		return "四川省"
	case "53":
		return "贵州省"
	case "54":
		return "云南省"
	case "55":
		return "西藏自治区"

	case "61":
		return "陕西省"
	case "62":
		return "甘肃省"
	case "63":
		return "青海省"
	case "64":
		return "宁夏回族自治区"
	case "65":
		return "新疆维吾尔自治区"

	case "71":
		return "台湾省"

	case "81":
		return "香港特别行政区"
	case "82":
		return "澳门特别行政区"
	}
	return ""
}

// ParseIDCard parse province and year by idcard
func ParseIDCard(idcard string) (string, string) {
	provinceID := Substr(idcard, 0, 2)
	year := Substr(idcard, 6, 4)
	return ID2Province(provinceID), year
}
