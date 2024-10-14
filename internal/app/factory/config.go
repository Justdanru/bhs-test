package factory

import (
	"github.com/Justdanru/bhs-test/config"
	"github.com/google/wire"
)

var configSet = wire.NewSet(
	config.NewConfig,
)
