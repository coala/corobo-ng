package server

import (
	"fmt"

	"github.com/coala/corobo-ng/config"
	"github.com/coala/corobo-ng/db"
)

func Init() {
	config := config.GetConfig()
	r := NewRouter(db.GetDB())
	r.Run(fmt.Sprintf("0.0.0.0:%s", config.Server.Port))
}
