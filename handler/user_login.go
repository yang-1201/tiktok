package handler

import (
	"context"
	"github.com/cloudwego/hertz/cmd/hz/util/logs"
	"github.com/cloudwego/hertz/pkg/app"
	"go_tiktok_project/common"
	"go_tiktok_project/common/dal/mysql"
	"go_tiktok_project/common/dal/rediss"
	pb "go_tiktok_project/idl/pb"
	"net/http"
<<<<<<< HEAD
	"reflect"
=======
>>>>>>> 2bb4450e1ea238fa6bab7b3bbf098c30b1fd617d
)

func UserLogin(ctx context.Context, c *app.RequestContext) {
	//defer func() {
	//	recover()
	//}()

	path := c.Request.Path()
	logs.Info("req path: ", path)

<<<<<<< HEAD
	req := new(pb.DouyinUserLoginRequest)
	resp := new(pb.DouyinUserLoginResponse)
	if err := c.BindAndValidate(&req); err != nil {
		c.JSON(400, err.Error())
		return
	}

	username := req.GetUsername()
	password := req.GetPassword()
=======
	resp := new(pb.DouyinUserLoginResponse)

	username := c.Query("Username")
	password := c.Query("Password")
>>>>>>> 2bb4450e1ea238fa6bab7b3bbf098c30b1fd617d

	user, err := mysql.FindUserByNameAndPass(username, password)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
<<<<<<< HEAD
=======
		return
>>>>>>> 2bb4450e1ea238fa6bab7b3bbf098c30b1fd617d
	}

	token, err := rediss.GetTokenByName(ctx, username)
	if err != nil {
<<<<<<< HEAD
		c.JSON(400, err.Error())
=======
		c.JSON(http.StatusBadRequest, err.Error())
>>>>>>> 2bb4450e1ea238fa6bab7b3bbf098c30b1fd617d
		return
	}

	// return
<<<<<<< HEAD
	*resp.StatusCode = common.LoginSuccess
	*resp.StatusMsg = common.LoginSuccessMsg
	*resp.UserId = cvt2id(user.UserID)
	*resp.Token = token

	c.JSON(http.StatusOK, resp)
}

// cvt2id
func cvt2id(num any) (id int64) {
	return reflect.ValueOf(num).Int()
}
=======

	loginsuccode := int32(common.LoginSuccess)
	loginsucmsg := common.LoginSuccessMsg

	resp.StatusCode = &loginsuccode
	resp.StatusMsg = &loginsucmsg
	resp.UserId = &user.UserID
	resp.Token = &token

	c.JSON(http.StatusOK, resp)
}
>>>>>>> 2bb4450e1ea238fa6bab7b3bbf098c30b1fd617d
