package datetime

import (
    "github.com/stretchr/testify/assert"
    "testing"
    "time"
)

func TestFormatTimeToStr(t *testing.T) {
    assert.Equal(t, "2022-03-14", FormatTimeToStr(time.Now(), "yyyy-mm-dd"))
}

func TestFormatStrToTime(t *testing.T) {
    toTime, err := FormatStrToTime("2022-03-14", "yyyy-mm-dd")
    assert.Equal(t, nil, err)
    assert.Equal(t, 2022, toTime.Year())
    assert.Equal(t, 3, int(toTime.Month()))
    assert.Equal(t, 14, toTime.Day())
}
