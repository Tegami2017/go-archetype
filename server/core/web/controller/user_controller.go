package controller

import (
	"github.com/gin-gonic/gin"
	"go_archetype/core/model/converter"
	"go_archetype/core/model/entity"
	"go_archetype/core/model/req"
	"go_archetype/core/model/vo"
	"go_archetype/core/storage/dao"
	"go_archetype/core/util"
	"strconv"
)

type UserController struct {
}

var UserCtrl = &UserController{}

func (r UserController) GetDetail(c *gin.Context) {
	userId, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		panic("invalid id")
	}
	userEntity := dao.User.SelectById(&userId)
	if userEntity == nil {
		panic("User not found.")
	}
	success(c, converter.UserConv.ToVO(userEntity))
}

func (r UserController) List(c *gin.Context) {
	userIdList := util.Array.ToIntArr(c.Query("ids"), ",")
	if len(userIdList) == 0 {
		panic("ids can not be empty")
	}
	users := dao.User.ListByIds(userIdList)
	userVOList := make([]*vo.User, len(*users))
	for i, user := range *users {
		userVOList[i] = converter.UserConv.ToVO(&user)
	}
	res := map[string]any{
		"list": userVOList,
	}
	success(c, res)
}

func (r UserController) Create(c *gin.Context) {
	var createUserReq req.CreateUserReq
	if err := c.ShouldBindJSON(&createUserReq); err != nil {
		panic(err.Error())
	}
	newUser := entity.User{
		Name:    &createUserReq.Name,
		Age:     &createUserReq.Age,
		Country: &createUserReq.Country,
	}
	dao.User.Create(&newUser)
	res := map[string]any{
		"Id": newUser.Id,
	}
	success(c, res)
}

func (r UserController) BatchCreate(c *gin.Context) {
	var reqBody []req.CreateUserReq
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		panic(err.Error())
	}
	newUsers := make([]entity.User, len(reqBody))
	for i, createUserReq := range reqBody {
		newUsers[i] = entity.User{
			Name:    &createUserReq.Name,
			Age:     &createUserReq.Age,
			Country: &createUserReq.Country,
		}
	}

	dao.User.BatchCreate(&newUsers)
	ids := make([]int, len(reqBody))
	for i, user := range newUsers {
		ids[i] = *user.Id
	}
	res := map[string]any{
		"Ids": ids,
	}
	success(c, res)
}
