package file

import (
    "io/ioutil"
    "os"
    "path"
    "path/filepath"
    "strings"
)

// CheckFileIsExist 检查目录是否存在
func CheckFileIsExist(filename string) bool {
    var exist = true
    if _, err := os.Stat(filename); os.IsNotExist(err) {
        exist = false
    }
    return exist
}

// BuildDir 创建目录
// 生成多级目录
func BuildDir(absDir string) error {
    return os.MkdirAll(path.Dir(absDir), os.ModePerm)
}

// DeleteFile 删除文件或文件夹
func DeleteFile(absDir string) error {
    return os.RemoveAll(absDir)
}

// GetPathDirs 获取目录所有文件夹
func GetPathDirs(absDir string) (re []string) {
    if CheckFileIsExist(absDir) {
        files, _ := ioutil.ReadDir(absDir)
        for _, f := range files {
            if f.IsDir() {
                re = append(re, f.Name())
            }
        }
    }
    return
}

// GetPathFiles 获取目录所有文件
func GetPathFiles(absDir string) (re []string) {
    if CheckFileIsExist(absDir) {
        files, _ := ioutil.ReadDir(absDir)
        for _, f := range files {
            if !f.IsDir() {
                re = append(re, f.Name())
            }
        }
    }
    return
}

// GetModelPath 获取程序运行目录
func GetModelPath() string {
    dir, _ := os.Getwd()
    return strings.Replace(dir, "\\", "/", -1)
}

// GetCurrentDirectory 获取exe所在目录
func GetCurrentDirectory() string {
    dir, _ := os.Executable()
    exPath := filepath.Dir(dir)
    // fmt.Println(exPath)

    return strings.Replace(exPath, "\\", "/", -1)
}
