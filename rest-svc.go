package abango

import (
	_ "github.com/go-sql-driver/mysql"
)

func RestSvcStandBy(RouterHandler func(*AbangoAsk)) {
	var v AbangoAsk
	RouterHandler(&v)
}
