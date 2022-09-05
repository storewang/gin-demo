package util

import (
	"log"
	"reflect"
)

func GetMethod(methodName string, controller interface{}) reflect.Value {
	v := reflect.ValueOf(controller)
	t := reflect.TypeOf(controller)
	mthd, b := t.MethodByName(methodName)
	if !b {
		log.Fatal("未获取到方法:", methodName)
		return reflect.Value{}
	}

	return v.Method(mthd.Index)
}
