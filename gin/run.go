package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v8"
	"net/http"
	"reflect"
	"sync/atomic"
	"time"
)

//binding 绑定一些验证请求参数,自定义标签bookabledate表示可预约的时期
type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckOut,bookabledate" time_format:"2006-01-02"`
}

type Get struct {
	Page  int    `json:"page" form:"page" `
	Limit int    `json:"limit" form:"limit"`
	Str   string `json:"str" form:"str" binding:"required`
}

//定义bookabledate标签对应的验证方法
func bookableDate(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	if date, ok := field.Interface().(time.Time); ok {
		today := time.Now()

		if date.Unix() > today.Unix() {
			return true
		}
	}
	return false
}

var total int64 = 0
var i100 int64 = 0
var i50 int64 = 0
var i40 int64 = 0
var i20 int64 = 0

func main() {
	route := gin.Default()

	route.Use(func(context *gin.Context) {
		start := time.Now()
		atomic.AddInt64(&total, 1)
		context.Next()
		t := time.Since(start).Milliseconds()
		if t > 100 {
			atomic.AddInt64(&i100, 1)
		} else if t > 50 {
			atomic.AddInt64(&i50, 1)
		} else if t > 40 {
			atomic.AddInt64(&i40, 1)
		} else if t > 20 {
			atomic.AddInt64(&i20, 1)
		}
		fmt.Println(fmt.Sprintf("耗时 %d ms", t))
	})
	//将验证方法注册到验证器中
	/*if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", bookableDate)
	}
	*/
	go func() {
		t := time.Tick(10 * time.Second)
		for true {
			select {
			case <-t:
				fmt.Println(fmt.Sprintf("total=%d, i100=%d, i50=%d, i40=%d,i20=%d ", total, i100, i50, i40, i20))
			}
		}
	}()
	route.GET("/bookable", getBookable)
	route.GET("/get", get)
	route.GET("/read", read)
	route.Run(":8080")
}

func getBookable(c *gin.Context) {
	var b Booking
	if err := c.ShouldBindQuery(&b); err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "Booking dates are valid!"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func get(c *gin.Context) {
	var b Get
	err := c.ShouldBindQuery(&b)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func read(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
