package lib

import (
	"database/sql/driver"

	post "github.com/lib/pq"
)

func Con() driver.Conn {

	s, _ := post.Open("gttg")

	return s
}
