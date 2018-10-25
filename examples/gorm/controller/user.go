// Copyright 2018 John Deng (hi.devops.io@gmail.com).
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package controller

import (
	"github.com/hidevopsio/hiboot-data/examples/gorm/entity"
	"github.com/hidevopsio/hiboot-data/examples/gorm/service"
	"github.com/hidevopsio/hiboot/pkg/app"
	"github.com/hidevopsio/hiboot/pkg/app/web"
	"github.com/hidevopsio/hiboot/pkg/model"
	"github.com/hidevopsio/hiboot/pkg/utils/copier"
	"net/http"
)

type userRequest struct {
	model.RequestBody
	Id       uint64 `json:"id"`
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Age      uint   `json:"age" validate:"gte=0,lte=130"`
	Gender   uint   `json:"gender" validate:"gte=0,lte=2"`
}

// RestController
type userController struct {
	web.Controller
	userService service.UserService
}

func init() {
	app.Register(newUserController)
}

// newUserController inject userService automatically
func newUserController(userService service.UserService) *userController {
	return &userController{
		userService: userService,
	}
}

// Post POST /user
func (c *userController) Post(request *userRequest) (model.Response, error) {
	var user entity.User
	copier.Copy(&user, request)
	err := c.userService.AddUser(&user)
	response := new(model.BaseResponse)
	response.SetData(user)
	return response, err
}

// GetById GET /id/{id}
func (c *userController) GetById(id uint64) (response model.Response, err error) {
	user, err := c.userService.GetUser(id)
	response = new(model.BaseResponse)
	if err != nil {
		response.SetCode(http.StatusNotFound)
	} else {
		response.SetData(user)
	}
	return
}

// GetById GET /id/{id}
func (c *userController) GetAll() (response model.Response, err error) {
	users, err := c.userService.GetAll()
	response = new(model.BaseResponse)
	response.SetData(users)
	return
}

// DeleteById DELETE /id/{id}
func (c *userController) DeleteById(id uint64) (response model.Response, err error) {
	err = c.userService.DeleteUser(id)
	response = new(model.BaseResponse)
	return
}
