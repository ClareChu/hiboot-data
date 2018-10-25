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

package data

import "github.com/hidevopsio/hiboot/pkg/utils/reflector"

// KVRepository is the Key/Value Repository interface
type KVRepository interface {
	Repository
	Put(params ...interface{}) error
	Get(params ...interface{}) error
	Delete(params ...interface{}) error
}

type BaseKVRepository struct {
	BaseRepository
}

func (r *BaseKVRepository) getKey(value interface{}, id string) string {
	k := reflector.GetFieldValue(value, id)
	if k.IsValid() && k.CanInterface() {
		return k.Interface().(string)
	}
	return ""
}

func (r *BaseKVRepository) Parse(params ...interface{}) ([]byte, []byte, interface{}, error) {
	var key string
	var value interface{}
	if len(params) == 2 {
		key = params[0].(string)
		value = params[1]
	} else {
		value = params[0]
		key = r.getKey(value, "ID")
		if key == "" {
			key = r.getKey(value, "Id")
		}
		if key == "" {
			return nil, nil, nil, InvalidDataModelError
		}
	}

	bucket := reflector.GetLowerCamelName(value)
	if bucket == "" {
		return nil, nil, nil, reflector.ErrInvalidInput
	}

	return []byte(bucket), []byte(key), value, nil
}

// Put inserts a key:value pair into the database
func (r *BaseKVRepository) Put(params ...interface{}) error {
	return nil
}

// Get retrieves a key:value pair from the database
func (r *BaseKVRepository) Get(params ...interface{}) error {
	return nil
}

// Delete removes a key:value pair from the database
func (r *BaseKVRepository) Delete(params ...interface{}) error {
	return nil
}
