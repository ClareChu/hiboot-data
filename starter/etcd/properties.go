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

package etcd

import "hidevops.io/hiboot/pkg/at"

type cert struct {
	CertFile      string `json:"cert_file" default:"config/certs/etcd.pem"`
	KeyFile       string `json:"key_file" default:"config/certs/etcd-key.pem"`
	TrustedCAFile string `json:"trusted_ca_file" default:"config/certs/ca.pem"`
}

type Properties struct {
	at.ConfigurationProperties `value:"etcd"`

	DialTimeout    int64    `json:"dial_timeout"`
	RequestTimeout int64    `json:"request_timeout"`
	Endpoints      []string `json:"endpoints"`
	Cert           cert     `json:"cert"`
}
