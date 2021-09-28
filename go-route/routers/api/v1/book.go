package v1

import (
	"book-server/controllers"
	"book-server/db/models"
	"book-server/pkg/app"
	"book-server/pkg/e"
	"fmt"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func ADDBook(c *gin.Context) {
	ag := app.Gin{C: c}
	book := new(models.BookList)
	err := c.BindJSON(book)
	if err != nil {
		ag.Response(500, 500, err)
	}
	mc := controllers.BookList{}
	if err := mc.ADD(book); err != nil {
		ag.Response(500, 500, err.Error())
		return
	}
	ag.Response(200, 200, nil)
	return
}

func UpdateBook(c *gin.Context) {
	ag := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	fmt.Println(id)
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		ag.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	mac := new(models.BookList)
	err := c.BindJSON(mac)
	if err != nil {
		ag.Response(http.StatusBadRequest, e.INVALID_PARAMS, err.Error())
		return
	}
	mc := controllers.BookList{}
	mc.ID = id
	if err := mc.GetBookList(); err != nil {
		fmt.Println(fmt.Sprintf("get book by id:%d, err: %v", id, err))
		ag.Response(500, e.ERROR, err.Error())
		return
	}

	mac.ID = id
	if err := mc.Update(mac); err != nil {
		fmt.Println("update book error :", err)
		ag.Response(500, e.ERROR, err.Error())
		return
	}
	ag.Response(200, e.SUCCESS, mac)
	return
}
