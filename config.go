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
	"context"
	"github.com/taouniverse/tao"
)

// ConfigKey for this repo
const ConfigKey = "mysql"

// Config implements tao.Config
// declare the configuration you want & define some default values
type Config struct {
	Host      string   `json:"host"`
	Port      int      `json:"port"`
	User      string   `json:"user"`
	Password  string   `json:"password"`
	DB        string   `json:"db"`
	Charset   string   `json:"charset"`
	ParseTime *bool    `json:"parse_time"`
	Location  string   `json:"local"`
	RunAfters []string `json:"run_after,omitempty"`
}

var trueVar = true

var defaultMysql = &Config{
	Host:      "localhost",
	Port:      3306,
	User:      "tao",
	Password:  "123456qwe",
	Charset:   "utf8mb4,utf8",
	ParseTime: &trueVar,
	Location:  "UTC",
	RunAfters: []string{},
}

// Name of Config
func (m *Config) Name() string {
	return ConfigKey
}

// ValidSelf with some default values
func (m *Config) ValidSelf() {
	if m.Host == "" {
		m.Host = defaultMysql.Host
	}
	if m.Port == 0 {
		m.Port = defaultMysql.Port
	}
	if m.User == "" {
		m.User = defaultMysql.User
	}
	if m.Password == "" {
		m.Password = defaultMysql.Password
	}
	if m.Charset == "" {
		m.Charset = defaultMysql.Charset
	}
	if m.ParseTime == nil {
		m.ParseTime = defaultMysql.ParseTime
	}
	if m.Location == "" {
		m.Location = defaultMysql.Location
	}
	if m.RunAfters == nil {
		m.RunAfters = defaultMysql.RunAfters
	}
}

// ToTask transform itself to Task
func (m *Config) ToTask() tao.Task {
	return tao.NewTask(
		ConfigKey,
		func(ctx context.Context, param tao.Parameter) (tao.Parameter, error) {
			// non-block check
			select {
			case <-ctx.Done():
				return param, tao.NewError(tao.ContextCanceled, "%s: context has been canceled", ConfigKey)
			default:
			}
			// JOB code run after RunAfters, you can just do nothing here
			db, err := DB.DB()
			if err != nil {
				return param, err
			}

			err = db.Ping()
			return param, err
		})
}

// RunAfter defines pre task names
func (m *Config) RunAfter() []string {
	return m.RunAfters
}
