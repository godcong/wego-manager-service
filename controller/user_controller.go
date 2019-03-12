package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/wego-auth-manager/config"
	"github.com/godcong/wego-auth-manager/model"
	"github.com/godcong/wego-auth-manager/util"
	log "github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
)

// UserLogin godoc
// @Summary Login user
// @Description Login user
// @Tags default
// @Accept  json
// @Produce  json
// @Param account body Login true "user update info"
// @success 200 {object} util.WebToken
// @Failure 400 {object} controller.CodeMessage
// @Router /login [post]
func UserLogin(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var u model.Login
		e := ctx.BindJSON(&u)

		if e != nil {
			Error(ctx, e)
			return
		}
		user := model.User{
			Username: u.Username,
		}

		b, e := user.Get()
		if e != nil {
			log.Info("get error1")
			Error(ctx, e)
			return
		}

		if !b {
			log.Info("get error2")
			Error(ctx, xerrors.New("username password is not correct"))
			return
		}

		b = user.Validate(&u, config.Config().WebToken.Key)
		if !b {
			log.Info("validate error")
			Error(ctx, xerrors.New("username password is not correct"))
			return
		}
		token := util.NewWebToken(user.ID, user.Username)
		token.Nickname = user.Nickname
		t, e := util.ToToken(config.Config().WebToken.Key, token)
		if e != nil {
			log.Info(e)
			Error(ctx, xerrors.New("username password is not correct"))
			return
		}

		user.Token = t
		i, e := user.Update("token")
		if e != nil || i != 1 {
			log.Info(e, i)
			Error(ctx, xerrors.New("unknown login error"))
			return
		}
		Success(ctx, gin.H{
			"token": t,
		})
		return
	}
}

// UserRegister godoc
// @Summary register user
// @Description register user
// @Tags default
// @Accept  json
// @Produce  json
// @Param account body User true "user update info"
// @success 200 {object} util.WebToken
// @Failure 400 {object} controller.CodeMessage
// @Router /register [post]
func UserRegister(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user model.User
		err := ctx.BindJSON(&user)
		if err != nil {
			Error(ctx, err)
			return
		}
		user.Salt = util.GenerateRandomString(16)
		user.Password = util.SHA256(user.Password, config.Config().WebToken.Key, user.Salt)
		_, err = model.Insert(nil, &user)
		if err != nil {
			Error(ctx, err)
			return
		}
		Success(ctx, user)
	}
}

// UserList godoc
// @Summary List users
// @Description List users
// @Tags admin
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @success 200 {array} model.Paginate
// @Failure 400 {object} controller.CodeMessage
// @Router /admin/user [get]
func UserList(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user model.User
		users, err := user.Users()
		if err != nil {
			Error(ctx, err)
			return
		}
		log.Info(users)
		Success(ctx, users)
	}
}

// UserAdd godoc
// @Summary Add user
// @Description Add user
// @Tags admin
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param account body User true "user update info"
// @success 200 {object} model.User
// @Failure 400 {object} controller.CodeMessage
// @Router /admin/user [post]
func UserAdd(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user model.User
		err := ctx.BindJSON(&user)
		if err != nil {
			Error(ctx, err)
			return
		}
		user.Password = util.SHA256(user.Password, config.Config().WebToken.Key, util.GenerateRandomString(16))
		_, err = model.Insert(nil, &user)
		if err != nil {
			Error(ctx, err)
			return
		}
		Success(ctx, user)
	}
}

// UserUpdate godoc
// @Summary Update user
// @Description Update user
// @Tags admin
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param id path string true "User ID"
// @Param account body User true "user update info"
// @success 200 {object} model.User
// @Failure 400 {object} controller.CodeMessage
// @Router /admin/user/{id} [post]
func UserUpdate(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		user := model.NewUser(id)
		b, err := user.Get()
		if err != nil || !b {
			Error(ctx, xerrors.Errorf("no users:%w", err))
			return
		}
		err = ctx.BindJSON(user)
		if err != nil {
			Error(ctx, err)
			return
		}

		_, err = model.Update(nil, id, user)
		if err != nil {
			Error(ctx, err)
			return
		}
		Success(ctx, user)
	}
}

// UserReset godoc
// @Summary reset user password
// @Description reset user password
// @Tags admin
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param id path string true "User ID"
// @Param account body User true "user update info"
// @success 200 {object} model.User
// @Failure 400 {object} controller.CodeMessage
// @Router /admin/user/{id}/reset [post]
func UserReset(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		user := model.NewUser(id)
		b, err := user.Get()
		if err != nil || !b {
			Error(ctx, xerrors.Errorf("no users:%w", err))
			return
		}
		login := model.Login{}
		err = ctx.BindJSON(&login)
		if err != nil {
			Error(ctx, err)
			return
		}
		if login.Password == "" {
			login.Password = "123456"
		}

		user.Password = util.SHA256(login.Password, config.Config().WebToken.Key, user.Salt)
		_, err = model.UpdateWithColumn(nil, id, user, "password")
		if err != nil {
			Error(ctx, err)
			return
		}
		Success(ctx, user)
	}
}

// UserShow godoc
// @Summary Show user
// @Description Show user
// @Tags admin
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param id path string true "User ID"
// @success 200 {object} model.User
// @Failure 400 {object} controller.CodeMessage
// @Router /admin/user/{id} [get]
func UserShow(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		user := model.NewUser(id)
		_, err := model.Get(nil, user)
		if err != nil {
			Error(ctx, err)
			return
		}
		Success(ctx, user)
	}
}

// UserInfo ...
func UserInfo(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userToken := model.GetUser(ctx)
		user := model.NewUser(userToken.ID)
		_, err := model.Get(nil, user)
		if err != nil {
			Error(ctx, err)
			return
		}
		Success(ctx, user)
	}
}
