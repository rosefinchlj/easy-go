package util

import (
    "testing"
)

func TestGetRandomString(t *testing.T) {
    randomString := GetRandomString(10)
    t.Log(randomString)

    if len(randomString) != 10 {
        t.FailNow()
    }
}

func TestGetRangeNumString(t *testing.T) {
    numString := GetRangeNumString(5)
    t.Log(numString)

    if len(numString) != 5 {
        t.FailNow()
    }
}

func TestGetRangeNum(t *testing.T) {
    num := GetRangeNum(6)
    t.Log(num)
}

func TestGetRandInt(t *testing.T) {
    num := GetRandInt(5, 100)
    t.Log(num)
}