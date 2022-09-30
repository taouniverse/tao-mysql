// Copyright 2022 huija
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mysql

import (
	"fmt"
	"github.com/taouniverse/tao"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/**
import _ "github.com/taouniverse/tao-mysql"
*/

// M config of mysql
var M = new(Config)

func init() {
	err := tao.Register(ConfigKey, M, setup)
	if err != nil {
		panic(err.Error())
	}
}

// DB orm client of mysql
var DB *gorm.DB

// setup unit with the global config 'M'
// execute when init tao universe
func setup() (err error) {
	var datetimePrecision = 2

	// https://github.com/go-sql-driver/mysql#dsn-data-source-name
	dsn := fmt.Sprintf(
		"gorm:gorm@tcp(%s:%d)/gorm?charset=%s&parseTime=%+v&loc=%s",
		M.Host, M.Port, M.Charset, M.ParseTime, M.Location)

	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DefaultDatetimePrecision:  &datetimePrecision,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})
	if err != nil {
		return tao.NewErrorWrapped("mysql: fail to create gorm client", err)
	}

	return
}
