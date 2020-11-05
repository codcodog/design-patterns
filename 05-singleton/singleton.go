package singleton

import "sync"

var Config *config
var once sync.Once

type config struct {
	db string
}

func NewConfig() *config {
	once.Do(func() {
		Config = &config{db: "mysql db address"}
	})

	return Config
}
