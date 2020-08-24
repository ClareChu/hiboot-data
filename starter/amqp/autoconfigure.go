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

package amqp

import (
	"hidevops.io/hiboot/pkg/app"
)

const Profile = "amqp"

type configuration struct {
	app.Configuration
	// the properties member name must be amqp if the mapstructure is amqp,
	// so that the reference can be parsed
	Properties *Properties
}

func newConfiguration(properties *Properties) *configuration {
	return &configuration{Properties: properties}
}

func init() {
	app.Register(newConfiguration, new(Properties))
}

// Repository method name must be unique
func (c *configuration) Channel() (chn *Channel) {
	chn = newChannel()
	err := chn.Connect(c.Properties)
	if err != nil {
		return nil
	}
	return chn
}

type NewChannel func() (chn *Channel)

func (c *configuration) NewChannel() NewChannel {
	return c.Channel
}