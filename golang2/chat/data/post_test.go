package data

import (
	"ChatServe/utils"
	"fmt"
	"testing"
)

func Test_PostUser(t *testing.T) {
	post := Post{
		Userid: 1,
	}
	if user, err := post.User(); err != nil {
		t.Error(err, "函数不通过")
	} else {
		fmt.Println("user:", user)
	}
}

func Test_GetPost(t *testing.T) {
	if post, err := GetPost(1); err != nil {
		t.Error(err, "函数不通过")
	} else {
		t.Log("测试通过", post)
	}
}

func Test_CreatePost(t *testing.T) {
	if err := CreatePost(1, 5, "正在测试回帖。\n请无视此回帖", 4); err != nil {
		t.Error(err, "函数不通过")
	}
}

func Test_AddRepNum(t *testing.T) {
	if err := AddRepNum(1); err != nil {
		t.Error(err, "函数不通过")
	}
}

func Test_IsCollected(t *testing.T) {
	collection, err := IsCollected(2, 1)
	if err != nil {
		t.Error(err)
	} else {
		utils.Logger.Println(collection)
	}
}

func Test_Collect(t *testing.T) {
	if num, err := Collect(1, 5); err != nil {
		t.Error(err)
	} else {
		utils.Logger.Println(num)
	}
}

func Test_Cancel(t *testing.T) {
	if _, err := Cancel(1, 5); err != nil {
		t.Error(err)
	}
}
