package dingnuts

import "github.com/xujiajun/nutsdb"

type DingNuts struct {
	db  *nutsdb.DB
	opt nutsdb.Options
}

// NewDingNuts
// 构建一个新的dingnuts对象
func NewDingNuts(dir string) (*DingNuts, error) {
	_opt := nutsdb.DefaultOptions
	_opt.Dir = dir
	_db, err := nutsdb.Open(_opt)
	return &DingNuts{db: _db, opt: _opt}, err
}

// NewDingNutsFromOpt
// 从opt构建一个新的dingnuts对象
func NewDingNutsFromOpt(youropt nutsdb.Options) (*DingNuts, error) {
	_db, err := nutsdb.Open(youropt)
	return &DingNuts{db: _db, opt: youropt}, err
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
