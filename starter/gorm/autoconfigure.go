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

package gorm

import (
	"github.com/hidevopsio/hiboot/pkg/app"
)

type configuration struct {
	app.Configuration
	// the properties member name must be Gorm if the mapstructure is gorm,
	// so that the reference can be parsed
	GormProperties properties `mapstructure:"gorm"`
}

func init() {
	app.Register(new(configuration))
}

func (c *configuration) dataSource() DataSource {
	dataSource := GetDataSource()
	if !dataSource.IsOpened() {
		dataSource.Open(&c.GormProperties)
	}
	return dataSource
}

// Repository method name must be unique
func (c *configuration) Repository() Repository {
	dataSource := c.dataSource()
	return dataSource.Repository()
}
