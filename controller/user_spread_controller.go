package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/wego-auth-manager/model"
	log "github.com/sirupsen/logrus"
)

// UserSpreadList 我的推广
func UserSpreadList(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//TODO
		user := model.GetUser(ctx)
		log.Error(user)
	}
}

// UserSpreadShare 我的分享
func UserSpreadShare(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//TODO
		user := model.GetUser(ctx)
		log.Error(user)
	}
}
