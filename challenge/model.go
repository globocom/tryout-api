package challenge

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type Path struct {
	Path       string `json:"path" xorm:"notnull"`
	HttpStatus int    `json:"http_status" xorm:"notnull"`
	Throughput int    `json:"throughput"`
	Input      string `json:"input" xorm:"notnull"`
	Output     string `json:"output"`
}

type Challenge struct {
	Name  string `json:"name" xorm:"notnull"`
	Paths []Path `json:"paths" xorm:"notnull"`
}

func Create(c Challenge) error {
	eng, err := xorm.NewEngine("mysql", "root@/tryout")
	if err != nil {
		return err
	}
	eng.Sync2(new(Challenge))
	_, err = eng.Insert(c)
	return err
}
