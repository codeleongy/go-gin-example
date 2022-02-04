package app

import (
	"fmt"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"github.com/leong-y/go-gin-example/pkg/e"
)

// BindAndValid binds and validates data
func BindAndValid(c *gin.Context, form interface{}) (int, int) {
	err := c.ShouldBind(form)
	if err != nil {
		return http.StatusBadRequest, e.INVALID_PARAMS
	}
	fmt.Printf("form: %v\n", form)
	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		println("111")
		return http.StatusInternalServerError, e.ERROR
	}
	if !check {
		println("222")
		MarkErrors(valid.Errors)
		return http.StatusBadRequest, e.INVALID_PARAMS
	}

	return http.StatusOK, e.SUCCESS
}
