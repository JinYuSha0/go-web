package models

import (
	"time"

	"github.com/kataras/iris"
)

type Config struct {
	Iris  iris.Configuration
	Jwt   jwt
	Mongo mongo
}

type jwt struct {
	Timeout time.Duration
	Secret  string
}

type mongo struct {
	Uri      string
	Database string
}
