package singleton

import "sync"

var conf *config
var once sync.Once

type config struct {
	db string
}

func NewConfig() *config {
	once.Do(func() {
		conf = &config{db: "mysql db address"}
	})

	return conf
}
