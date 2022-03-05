// Package dingmysql
// @Description: 仅对本地mysql进行操作，不涉及增删查改，仅windows可使用
package dingmysql

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/axgle/mahonia"
	"github.com/dingdinglz/dingtools/dinglog"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
)

const version = "v0.1"

func init() {
	logger := dinglog.NewLogger()
	logger.Info("dingdb", "dingmysql", "version:", version)
}

var mysqlpath = string("")

// SetMysqlPath
// 该函数是必须调用的，设置mysql在您机子上的绝对位置
func SetMysqlPath(dir string) {
	mysqlpath = dir
}

// Mysql_WriteBat
// 内部写出操作指令的函数
func Mysql_WriteBat(text string) {
	os.WriteFile(filepath.Join(mysqlpath, "bin", "run.bat"), []byte(text), 0666)
}

// ConvertToString
// 内部编码转换函数
func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

// Mysql_RunBat
// 内部，执行指令运行
func Mysql_RunBat() (error, string) {
	runMySqlBat := exec.Command("cmd", "/c", "run.bat")
	runMySqlBat.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	runMySqlBat.Dir = filepath.Join(mysqlpath, "bin")
	var out, errout bytes.Buffer
	runMySqlBat.Stdout = &out
	runMySqlBat.Stderr = &errout
	err := runMySqlBat.Run()
	if err != nil {
		return errors.New(ConvertToString(errout.String(), "gbk", "utf-8")), ConvertToString(out.String(), "gbk", "utf-8")
	}
	return nil, ConvertToString(out.String(), "gbk", "utf-8")
}

// Mysql_Init
// 初始化mysql
func Mysql_Init() (error, string) {
	var text_bat string = `@echo off
mysqld --initialize-insecure --user=mysql
mysqld install mysql`
	Mysql_WriteBat(text_bat)
	err, s := Mysql_RunBat()
	if strings.Count(s, "Install/Remove of the Service Denied!") != 0 {
		return errors.New("权限不够，请以管理员权限执行该命令！"), ""
	}
	return err, s
}

// Mysql_start
// 启动mysql
func Mysql_start() (error, string) {
	var text_bat string = `@echo off
net start mysql`
	Mysql_WriteBat(text_bat)
	return Mysql_RunBat()
}

// Mysql_updatePass
// 修改密码，初始化之后调用时oldpass为空
func Mysql_updatePass(oldpass string, newpass string) (error, string) {
	var text_bat string
	if oldpass == "" {
		text_bat = `mysqladmin -uroot password ` + newpass
	} else {
		text_bat = `mysqladmin -uroot -p` + oldpass + ` password ` + newpass
	}
	Mysql_WriteBat(text_bat)
	return Mysql_RunBat()
}

// Mysql_stop
// 停用mysql
func Mysql_stop() (error, string) {
	var text_bat string = `@echo off
net stop mysql`
	Mysql_WriteBat(text_bat)
	return Mysql_RunBat()
}

// Mysql_remove
// 移除mysql启动
func Mysql_remove() (error, string) {
	var text_bat string = `@echo off
mysqld remove`
	Mysql_WriteBat(text_bat)
	return Mysql_RunBat()
}

// MySql_CreateDatabase
// 新建一个数据库
func MySql_CreateDatabase(name string, pass string) (error, string) {
	var text_bat string = `@echo off
mysqladmin create ` + name + ` -uroot -p` + pass
	Mysql_WriteBat(text_bat)
	return Mysql_RunBat()
}

// Mysql_DropDatabase
// 删除一个数据库
func Mysql_DropDatabase(name string, pass string) (error, string) {
	var text_bat string = `@echo off
mysqladmin drop %s -uroot -p%s -f`
	text_bat = fmt.Sprintf(text_bat, name, pass)
	Mysql_WriteBat(text_bat)
	return Mysql_RunBat()
}

// Mysql_Clean
// 清理mysql数据，需要先停止，再使用
func Mysql_Clean() error {
	return os.RemoveAll(filepath.Join(mysqlpath, "data"))
}

func Mysql_GetDatabases(pass string) []string {
	var text_bat string = `@echo off
mysqlshow -uroot -p` + pass
	Mysql_WriteBat(text_bat)
	runMySqlBat := exec.Command("cmd", "/c", "run.bat")
	runMySqlBat.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	runMySqlBat.Dir = filepath.Join(mysqlpath, "bin")
	var out, errout bytes.Buffer
	runMySqlBat.Stdout = &out
	runMySqlBat.Stderr = &errout
	runMySqlBat.Run()
	databases := out.String()
	s1 := strings.Split(databases, "\n")
	var s2 []string
	for _, i := range s1 {
		if strings.HasPrefix(i, "|") {
			i = strings.ReplaceAll(i, "|", "")
			if !strings.HasPrefix(i, "   ") {
				i = strings.ReplaceAll(i, " ", "")
				s2 = append(s2, i)
			}
		}
	}
	return s2
}
