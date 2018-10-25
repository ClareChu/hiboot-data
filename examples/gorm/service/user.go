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

package service

import (
	"errors"
	"github.com/hidevopsio/hiboot-data/examples/gorm/entity"
	"github.com/hidevopsio/hiboot-data/starter/gorm"
	"github.com/hidevopsio/hiboot/pkg/app"
	"github.com/hidevopsio/hiboot/pkg/utils/idgen"
)

type UserService interface {
	AddUser(user *entity.User) (err error)
	GetUser(id uint64) (user *entity.User, err error)
	GetAll() (user *[]entity.User, err error)
	DeleteUser(id uint64) (err error)
}

type userServiceImpl struct {
	// add UserService, it means that the instance of UserServiceImpl can be found by UserService
	UserService
	repository gorm.Repository
}

func init() {
	// register UserServiceImpl
	app.Register(newUserService)
}

// will inject BoltRepository that configured in github.com/hidevopsio/hiboot/pkg/starter/data/bolt
func newUserService(repository gorm.Repository) UserService {
	repository.AutoMigrate(&entity.User{})
	return &userServiceImpl{
		repository: repository,
	}
}

func (s *userServiceImpl) AddUser(user *entity.User) (err error) {
	if user == nil {
		return errors.New("user is not allowed nil")
	}
	if user.Id == 0 {
		user.Id, _ = idgen.Next()
	}
	err = s.repository.Create(user).Error()
	return
}

func (s *userServiceImpl) GetUser(id uint64) (user *entity.User, err error) {
	user = &entity.User{}
	err = s.repository.Where("id = ?", id).First(user).Error()
	return
}

func (s *userServiceImpl) GetAll() (users *[]entity.User, err error) {
	users = &[]entity.User{}
	err = s.repository.Find(users).Error()
	return
}

func (s *userServiceImpl) DeleteUser(id uint64) (err error) {
	err = s.repository.Where("id = ?", id).Delete(entity.User{}).Error()
	return
}
