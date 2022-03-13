package file

import (
    "github.com/gocarina/gocsv"
    "os"
)

// ParseCsv 解析
func ParseCsv(filename string, dst interface{}) error {
    f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, os.ModePerm)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    return gocsv.UnmarshalFile(f, dst) // to dst
}
