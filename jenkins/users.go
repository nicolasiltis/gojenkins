// Copyright 2015 Vadim Kravcenko
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package gojenkins

type User struct {
	Raw     *UserResponse
	Jenkins *Jenkins
	Base    string
}

type asynchPeople struct {
	Class       string 			`json:"_class"`
	Users       []*UserResponse `json:"users"`
}

type UserResponse struct {
	LastChange string `json:"lastChange"`
	project string `json:"project"`
	user struct {
		absoluteUrl string `json:"absoluteUrl"`
		fullName string `json:"fullName"`
	} `json:"user"`
}