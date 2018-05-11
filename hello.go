package main

import (
    "fmt"
    "github.com/apodemakeles/ugo/time"
    "github.com/astaxie/beego"
    "github.com/bitly/go-simplejson"
    "runtime"
)

func main() {
    fmt.Println(utime.NowUnixTS())
    maxCPU := runtime.NumCPU()
    runtime.GOMAXPROCS(maxCPU)
    strJson := `{"announcer": {"nickname": "非议讲史", "kind": "user", "created_at": 1494904539000, "updated_at": 1494983507000, "track_id": 38088960}}`
    mapJson,_:= simplejson.NewJson([]byte(strJson))
	v := mapJson.Get("announcer").Get("created_at")
	val, _ := v.Int64()
	println(val)
    beego.Run()
}
