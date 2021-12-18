package utility

import (
	"fmt"
	"testing"
)

func TestAddJsonFormTag(t *testing.T) {
	rs := AddJsonFormGormTag(`
	type User struct{
		Id        uint64
		NameSpace  string
		DailyRoutine int
	}
	`)
	fmt.Println(rs)
}

func TestTrans(t *testing.T) {
	rs := AddJsonFormGormTag(`
		type Trans struct {
			From string
			To string
			MediaId uint64
			TxHash string
		}
	`)
	fmt.Println(rs)
}
