// Copyright 2013 The StudyGolang Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// http://studygolang.com
// Author: polaris	polaris@studygolang.com

package model

const (
	LinkForm    = "只是链接"
	ContentForm = "包括内容"
)

// 资源信息
type Resource struct {
	Id      int    `json:"id" xorm:"pk autoincr"`
	Title   string `json:"title"`
	Form    string `json:"form"`
	Content string `json:"content"`
	Url     string `json:"url"`
	Uid     int    `json:"uid"`
	Catid   int    `json:"catid"`
	CatName string `json:"-" xorm:"-"`
	Ctime   string `json:"ctime,omitempty" xorm:"created"`
	Mtime   string `json:"mtime,omitempty" xorm:"<-"`
}

// 资源扩展（计数）信息
type ResourceEx struct {
	Id      int    `json:"-" xorm:"pk"`
	Viewnum int    `json:"viewnum"`
	Cmtnum  int    `json:"cmtnum"`
	Likenum int    `json:"likenum"`
	Mtime   string `json:"mtime" xorm:"<-"`
}

type ResourceInfo struct {
	Resource   `xorm:"extends"`
	ResourceEx `xorm:"extends"`
}

func (*ResourceInfo) TableName() string {
	return "resource"
}

// 资源分类信息
type ResourceCat struct {
	Catid int    `json:"catid" xorm:"pk autoincr"`
	Name  string `json:"name"`
	Intro string `json:"intro"`
	Ctime string `json:"ctime" xorm:"<-"`
}

func (*ResourceCat) TableName() string {
	return "resource_category"
}
