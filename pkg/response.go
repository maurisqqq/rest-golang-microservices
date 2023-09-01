package pkg

import (
	"encoding/json"
	generalModel "jojonomic/app/config"
)

type ResponseParam struct {
	Val    bool
	Method string
	Data   interface{}
	Err    error
}

func GeneralResponse(param ResponseParam) []byte {
	var res interface{}
	if param.Val == false && param.Method == "updt" {
		res = generalModel.PostResponseSuccess{
			Error:  false,
			ReffId: "risksdfxx",
		}
	} else if param.Val == false || param.Method == "GET" {
		res = generalModel.GetResponseSuccess{
			Error: false,
			Data:  StructToMap(param.Data),
		}
	} else {
		res = generalModel.ResponseError{
			Error:   true,
			ReffId:  "xkdfd",
			Message: param.Err.Error(),
		}
	}
	resByte, _ := json.Marshal(res)
	return resByte
}

func StructToMap(obj interface{}) (newMap interface{}) {
	data, err := json.Marshal(obj)

	if err != nil {
		return
	}

	err = json.Unmarshal(data, &newMap)
	return
}
