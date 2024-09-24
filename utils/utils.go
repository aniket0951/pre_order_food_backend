package utils

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Response struct {
	Status       bool        `json:"status,omitempty"`
	Message      string      `json:"message,omitempty"`
	Error        string      `json:"error,omitempty"`
	ResponseBody interface{} `json:"data,omitempty"`
}

func BuildSuccessResponse(msg string, responseData interface{}) *Response {
	return &Response{
		Status:       true,
		Message:      msg,
		ResponseBody: responseData,
	}
}

func BuildFailedResponse(errorMsg string) Response {

	response := Response{
		Status:  false,
		Message: "Failed to process the request",
		Error:   errorMsg,
	}

	return response
}

// UUID with prefix without -
func UUIDWithPrefix(prefix string) string {
	id := uuid.New().String()
	id = prefix + "_" + id
	id = strings.ReplaceAll(id, "-", "")
	return id
}

func IsValidaUUID(ctx *gin.Context, id, tag string) uuid.UUID {
	objID, err := uuid.Parse(id)
	if err != nil {
		response := BuildFailedResponse(fmt.Sprintf("invalid %s id", tag))
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
	return objID
}
