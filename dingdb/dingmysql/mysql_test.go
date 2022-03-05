package dingmysql

import (
	"github.com/dingdinglz/dingtools/dinglog"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMysql(t *testing.T) {
	logger := dinglog.NewLogger()
	SetMysqlPath("D:\\mysql")
	err, s := Mysql_Init()
	assert.NoError(t, err)
	logger.Info(s)
	err, s3 := Mysql_start()
	assert.NoError(t, err)
	logger.Info(s3)
	err, s2 := Mysql_updatePass("", "dinglznb")
	assert.NoError(t, err)
	logger.Info(s2)
}

func TestStopAndCleanMysql(t *testing.T) {
	logger := dinglog.NewLogger()
	SetMysqlPath("D:\\mysql")
	err, s := Mysql_stop()
	assert.NoError(t, err)
	logger.Info(s)
	err, s2 := Mysql_remove()
	assert.NoError(t, err)
	logger.Info(s2)
	err = Mysql_Clean()
	assert.NoError(t, err)
	logger.Info("清理完成！")
}

func TestDatabases(t *testing.T) {
	logger := dinglog.NewLogger()
	SetMysqlPath("D:\\mysql")
	err, s := MySql_CreateDatabase("wolong", "dinglznb")
	assert.NoError(t, err)
	logger.Info(s)
	s1 := Mysql_GetDatabases("dinglznb")
	for _, i := range s1 {
		logger.Info(i)
	}
	err, s2 := Mysql_DropDatabase("wolong", "dinglznb")
	assert.NoError(t, err)
	logger.Info(s2)
	s3 := Mysql_GetDatabases("dinglznb")
	for _, i := range s3 {
		logger.Info(i)
	}
}

func TestShowDatabases(t *testing.T) {
	SetMysqlPath("D:\\mysql")
	logger := dinglog.NewLogger()
	s := Mysql_GetDatabases("dinglznb")
	for _, i := range s {
		logger.Info(i)
	}
}
