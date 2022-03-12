package gen

import (
    "errors"
    "testing"
)

func TestGetResultMsg(t *testing.T) {
    if GetResultMsg(UserExisted).Message != "用户已存在" {
        t.FailNow()
    }
}

func TestGetResultMsgWithString(t *testing.T) {
    if GetResultMsg(UserExisted, "用户未找到").Message == "用户已存在" {
        t.FailNow()
    }
}

func TestGetResultMsgWithErr(t *testing.T) {
    if GetResultMsg(UserExisted, errors.New("用户未找到")).Message == "用户已存在" {
        t.FailNow()
    }
}