// Copyright 2013 gcblog authors.
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

package controllers

import (
	"github.com/Unknwon/gcblog/models"
)

type HomeController struct {
	baseController
}

func (this *HomeController) Get() {
	this.Data["IsHome"] = true
	this.TplNames = "home.html"
	this.Data["RecentArchives"] = models.GetRecentPosts()
}
