package dingnuts

import (
	"errors"
	"github.com/dingdinglz/dingtools/dinglog"
	"github.com/xujiajun/nutsdb"
)

const version = "v0.1"

func init() {
	logger := dinglog.NewLogger()
	logger.Info("dingdb", "dingnuts", "version:", version)
}

// Version
// 返回dingnuts的版本号
func Version() string {
	return version
}

type DingNuts struct {
	db     *nutsdb.DB
	opt    nutsdb.Options
	bucket string
}

// NewDingNuts
// 构建一个新的dingnuts对象
func NewDingNuts(dir string) (*DingNuts, error) {
	_opt := nutsdb.DefaultOptions
	_opt.Dir = dir
	_db, err := nutsdb.Open(_opt)
	return &DingNuts{db: _db, opt: _opt, bucket: ""}, err
}

// NewDingNutsFromOpt
// 从opt构建一个新的dingnuts对象
func NewDingNutsFromOpt(youropt nutsdb.Options) (*DingNuts, error) {
	_db, err := nutsdb.Open(youropt)
	return &DingNuts{db: _db, opt: youropt, bucket: ""}, err
}

// GetOpt
// 获取db的设置
func (d *DingNuts) GetOpt() nutsdb.Options {
	return d.opt
}

// Close
// 关闭数据库
func (d *DingNuts) Close() error {
	return d.db.Close()
}

// Use
// 指定操作bucket
func (d *DingNuts) Use(name string) {
	d.bucket = name
}

// GetBucket
// 传回一个新的bucket对象
func (d *DingNuts) GetBucket(name string) *DingNuts {
	return &DingNuts{db: d.db, opt: d.opt, bucket: name}
}

// SetValue
// 设置字段
func (d *DingNuts) SetValue(name []byte, value []byte) error {
	return d.db.Update(func(tx *nutsdb.Tx) error {
		return tx.Put(d.bucket, name, value, 0)
	})
}

// GetValue
// 取字段值
func (d *DingNuts) GetValue(name []byte) ([]byte, error) {
	var vaule []byte
	err := d.db.View(func(tx *nutsdb.Tx) error {
		_vaule, err := tx.Get(d.bucket, name)
		vaule = _vaule.Value
		return err
	})
	return vaule, err
}

// BackUp
// 备份数据库，不可单独备份一个bucket
func (d *DingNuts) BackUp(dir string) error {
	if d.bucket != "" {
		return errors.New("can't backup a bucket")
	}
	return d.db.Backup(dir)
}
