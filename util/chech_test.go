package util

import "testing"

func TestCheckNullParam(t *testing.T) {
    if !CheckNullParam("", "") {
        t.FailNow()
    }
}

func TestIsPhone(t *testing.T) {
    if !IsPhone("18565768888") {
        t.FailNow()
    }

    if IsPhone("28565768888") {
        t.FailNow()
    }

    if IsPhone("185657688880") {
        t.FailNow()
    }

    if IsPhone("1856576888a") {
        t.FailNow()
    }

    if IsPhone("1856576888.") {
        t.FailNow()
    }
}

func TestIsMail(t *testing.T) {
    if !IsMail("380654493@qq.com") {
        t.FailNow()
    }

    if !IsMail("380654493@sina.com") {
        t.FailNow()
    }

    if !IsMail("380654493@gmail.com") {
        t.FailNow()
    }

    if !IsMail("380654493@outlook.com.cn") {
        t.FailNow()
    }

    if !IsMail("380654493@outlook.cn") {
        t.FailNow()
    }

    if IsMail("380654493@outlook") {
        t.FailNow()
    }

    if IsMail("380654493.com") {
        t.FailNow()
    }

    if IsMail("380654493@@qq.com") {
        t.FailNow()
    }
}