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
	"hidevops.io/hiboot-data/examples/bolt/entity"
	"hidevops.io/hiboot/pkg/app/web"
	"hidevops.io/hiboot/pkg/log"
	"net/http"
	"testing"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func TestCrdRequest(t *testing.T) {
	// TODO: mock UserService
	testApp := web.NewTestApp(newUserController).Run(t)

	t.Run("should add user with POST request", func(t *testing.T) {
		// First, let's Post User
		testApp.Post("/user").
			WithJSON(entity.User{Id: "1", Name: "Peter", Age: 18}).
			Expect().Status(http.StatusOK)
	})

	t.Run("should get user with GET request", func(t *testing.T) {
		// Then Get User
		testApp.Get("/user/id/{id}").
			WithPath("id", "1").
			Expect().Status(http.StatusOK)
	})

	t.Run("should return 404 if trying to find a record that does not exist", func(t *testing.T) {
		// Then Get User
		testApp.Get("/user/id/{id}").
			WithPath("id", "9999").
			Expect().Status(http.StatusNotFound)
	})

	t.Run("should delete the record with DELETE request", func(t *testing.T) {
		// Finally Delete User
		testApp.Delete("/user/id/{id}").
			WithPath("id", "1").
			Expect().Status(http.StatusOK)
	})

}
