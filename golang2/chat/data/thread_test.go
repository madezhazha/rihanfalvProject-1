package data

import (
	"fmt"
	"testing"
)

func Test_GetThread(t *testing.T) {
	if thread, err := GetThread(); err != nil {
		t.Error(err, "函数不通过")
	} else {
		t.Log("测试通过", thread)
	}
}

func Test_ThreadUser(t *testing.T) {
	thread := Thread{
		Userid: 1,
	}
	if user, err := thread.User(); err != nil {
		t.Error(err, "函数不通过")
	} else {
		fmt.Println("user:", user)
	}

}

func Test_ThreadByTopicID(t *testing.T) {
	if userAndThread, err := ThreadByTopicID(1); err != nil {
		t.Error(err, "函数不通过")
	} else {
		t.Log(userAndThread)
	}

}
