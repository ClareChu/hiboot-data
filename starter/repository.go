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

package starter

import "errors"

var (

	// ErrInvalidDataSource invalid data source
	ErrInvalidDataSource = errors.New("invalid dataSource")

	// ErrInvalidDataModel invalid data model
	ErrInvalidDataModel = errors.New("invalid data model, must contains string field ID and assigns string value")

	// ErrNotImplemented method is not implemented
	ErrNotImplemented = errors.New("method is not implemented")
)

// Repository interface
type Repository interface {
	SetDataSource(dataSource interface{})
	DataSource() interface{}
	CloseDataSource() error
}

// BaseRepository base repository
type BaseRepository struct {
}

func (r *BaseRepository) SetDataSource(dataSource interface{}) {

}

func (r *BaseRepository) DataSource() interface{} {
	return ErrNotImplemented
}

func (r *BaseRepository) CloseDataSource() error {
	return ErrNotImplemented
}
