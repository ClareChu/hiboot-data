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
	"errors"
	"fmt"
	"github.com/hidevopsio/gorm"
	_ "github.com/hidevopsio/gorm/dialects/mssql"
	_ "github.com/hidevopsio/gorm/dialects/mysql"
	_ "github.com/hidevopsio/gorm/dialects/postgres"
	_ "github.com/hidevopsio/gorm/dialects/sqlite"
	"hidevops.io/hiboot/pkg/log"
	"hidevops.io/hiboot/pkg/utils/crypto/rsa"
	"strings"
	"sync"
)

type Repository interface {
	gorm.Repository
}

type DataSource interface {
	Open(p *Properties) error
	IsOpened() bool
	Close() error
	Repository() gorm.Repository
}

type dataSource struct {
	repository gorm.Repository
}

var DatabaseIsNotOpenedError = errors.New("database is not opened")
var ds *dataSource
var dsOnce sync.Once

func GetDataSource() DataSource {
	dsOnce.Do(func() {
		ds = new(dataSource)
	})
	return ds
}

func (d *dataSource) Init(repository Repository) {
	d.repository = repository
}

func (d *dataSource) Open(p *Properties) error {
	var err error
	password := p.Password
	if p.Config.Decrypt {
		pwd, err := rsa.DecryptBase64([]byte(password), []byte(p.Config.DecryptKey))
		if err == nil {
			password = string(pwd)
		}
	}
	loc := strings.Replace(p.Loc, "/", "%2F", -1)
	databaseName := strings.Replace(p.Database, "-", "_", -1)
	parseTime := "False"
	if p.ParseTime {
		parseTime = "True"
	}
	source := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=%v&loc=%v",
		p.Username, password, p.Host, p.Port, databaseName, p.Charset, parseTime, loc)

	d.repository, err = gorm.Open(p.Type, source)
	db := d.repository.SqlDB()
	db.SetConnMaxLifetime(p.ConnMaxLifetime)
	db.SetMaxIdleConns(p.MaxIdleConns)
	db.SetMaxOpenConns(p.MaxOpenConns)
	if err != nil {
		log.Errorf("dataSource connection failed: %v (%v)", err, p)
		defer func() {
			d.repository.Close()
			d.repository = nil
		}()
		return err
	} else {
		log.Infof("connected to dataSource %v@%v:%v/%v", p.Username, p.Host, p.Port, databaseName)
	}

	return nil
}

func (d *dataSource) IsOpened() bool {
	return d.repository != nil
}

func (d *dataSource) Close() error {
	if d.repository != nil {
		err := d.repository.Close()
		d.repository = nil
		return err
	}
	return DatabaseIsNotOpenedError
}

func (d *dataSource) Repository() gorm.Repository {
	return d.repository
}
