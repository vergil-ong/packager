package util

import "github.com/kataras/iris"

func BuildIrisMap(result bool,msg string, mapObj iris.Map) iris.Map  {
	var retMap iris.Map
	if result {
		retMap = iris.Map{
			"code": 200,
			"msg": msg,
		}
		if mapObj != nil {
			for k,v := range mapObj {
				retMap[k] = v;
			}
		}

	}else{
		retMap = iris.Map{
			"code": 404,
			"msg": msg,
		}
	}
	return retMap
}
