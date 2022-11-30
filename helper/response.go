package helper

import (
    "strings"
)

type Response struct {
    Status  bool        `json:"status"`
    Message string      `json:"message"`
    Error   interface{} `json:"errors"`
    Data    interface{} `json:"data"`
}

type EmptyObj struct{}

func BuildResponse(status bool, message string, data interface{}) Response {
    res := Response{
        Status:  status,
        Message: message,
        Data:    data,
    }

    return res
}

func BuildErrorResponse(err string, message string, data interface{}) Response {
    splitted_errors := strings.Split(err, "\n")
    res := Response{
        Status:  false,
        Message: message,
        Data:    data,
        Error:   splitted_errors,
    }

    return res
}
