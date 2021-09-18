package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"wechat-go/service/user"
)

type LoginInParams struct {
	User	*user.User `json:"user"`
}

type LoginOut struct {
	Uid		string
}

func Login(ctx *gin.Context)  {
	var in LoginInParams
	var out LoginOut
	if err := ctx.ShouldBindJSON(&in); err != nil {
		fmt.Println(err)
		return
	}
	inter := user.CheckUser(in.User.NickName)
	if inter == nil {
		fmt.Println("Login Check Error")
		return
	}
	out.Uid = inter.(string)
	ctx.JSON(http.StatusOK, gin.H{ "status": "200", "uid":out.Uid})
}
