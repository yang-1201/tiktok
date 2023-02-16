package handler

import (
	"context"
<<<<<<< HEAD
=======
	"github.com/cloudwego/hertz/cmd/hz/util/logs"
>>>>>>> 2bb4450e1ea238fa6bab7b3bbf098c30b1fd617d
	"go_tiktok_project/common"
	"go_tiktok_project/common/dal/mysql"
	"go_tiktok_project/common/dal/rediss"
	pb "go_tiktok_project/idl/pb"
	"go_tiktok_project/service"
	"net/http"

<<<<<<< HEAD
	"github.com/cloudwego/hertz/cmd/hz/util/logs"
=======
>>>>>>> 2bb4450e1ea238fa6bab7b3bbf098c30b1fd617d
	"github.com/cloudwego/hertz/pkg/app"
)

func UserRegister(ctx context.Context, c *app.RequestContext) {
	path := c.Request.Path()

	logs.Info("req path: ", path)

<<<<<<< HEAD
	req := new(pb.DouyinUserRegisterRequest)

	if err := c.BindAndValidate(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if res := service.IsUsernameLegal(req.GetUsername()); !res {
=======
	//req := new(pb.DouyinUserRegisterRequest)

	username := c.Query("Username")
	password := c.Query("Password")
	//fmt.Println(username)
	//fmt.Println(password)

	if res := service.IsUsernameLegal(username); !res {
>>>>>>> 2bb4450e1ea238fa6bab7b3bbf098c30b1fd617d
		c.JSON(http.StatusBadRequest, "username is illegal")
		return
	}

<<<<<<< HEAD
	res, err := service.IsUsernameExist(req.GetUsername())
=======
	res, err := service.IsUsernameExist(username)
>>>>>>> 2bb4450e1ea238fa6bab7b3bbf098c30b1fd617d
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if res {
		c.JSON(http.StatusBadRequest, "username already used")
		return
	}

<<<<<<< HEAD
	userID, err := mysql.CreateUser(req.GetUsername(), req.GetPassword())
=======
	userID, err := mysql.CreateUser(username, password)
>>>>>>> 2bb4450e1ea238fa6bab7b3bbf098c30b1fd617d
	if err != nil {
		c.JSON(http.StatusBadRequest, "insert to mysql failed")
		return
	}

<<<<<<< HEAD
	token, err := service.GenerateToken(uint64(userID), req.GetUsername())
=======
	token, err := service.GenerateToken(uint64(userID), username)
>>>>>>> 2bb4450e1ea238fa6bab7b3bbf098c30b1fd617d
	if err != nil {
		c.JSON(http.StatusBadRequest, "generate token failed")
		return
	}

<<<<<<< HEAD
	err = rediss.SetToken(ctx, req.GetUsername(), token)
=======
	err = rediss.SetToken(ctx, username, token)
>>>>>>> 2bb4450e1ea238fa6bab7b3bbf098c30b1fd617d
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	resp := new(pb.DouyinUserRegisterResponse)

<<<<<<< HEAD
	*resp.StatusCode = common.LoginSuccess
	*resp.StatusMsg = "register success"
	// userid uint64 or int64?
	*resp.UserId = cvt2id(userID)
	*resp.Token = token
=======
	statuscode := int32(common.RegisterSucces)
	statusmsg := common.RegisterSueecssMsg

	resp.StatusCode = &statuscode
	resp.StatusMsg = &statusmsg
	resp.UserId = &userID
	resp.Token = &token
>>>>>>> 2bb4450e1ea238fa6bab7b3bbf098c30b1fd617d

	c.JSON(http.StatusOK, resp)
}
