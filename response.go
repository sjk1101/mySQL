package mySQL

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Resp struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Page    interface{} `json:"page,omitempty"`
	Status  int         `json:"status"`
}

func CommonResponse(c *gin.Context, data interface{}, err error) {
	CommonPageResponse(c, data, nil, err)
}

func CommonPageResponse(c *gin.Context, data, page interface{}, err error) {
	message := "操作"
	statusCode := http.StatusOK
	status := 1
	if err != nil {
		message += "失敗: " + err.Error()
		status = -1
		// status = http.StatusInternalServerError
	} else {
		message += "成功"
	}
	result := Resp{
		Message: message,
		Data:    data,
		Page:    page,
		Status:  status,
	}
	c.JSON(statusCode, result)
}

func CommonResponseParamError(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, Resp{Message: "參數錯誤: " + err.Error(), Status: -1})
}

func CommonResponseInternalServerError(c *gin.Context, err error) {
	// 目前前端回應只接受 200 400
	c.JSON(http.StatusBadRequest, Resp{Message: "服務器錯誤: " + err.Error(), Status: -1})
}

func CommonResponseForbiddenError(c *gin.Context, err error) {
	c.JSON(http.StatusForbidden, Resp{Message: err.Error(), Status: -1})
}

func ConvertStruct(one interface{}, another interface{}) (interface{}, error) {
	bytes, _ := json.Marshal(one)
	err := json.Unmarshal(bytes, &another)
	if err != nil {
		return nil, err
	}
	return another, nil
}
