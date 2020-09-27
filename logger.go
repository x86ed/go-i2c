package i2c

import (
	"os"
	"strconv"

	logger "github.com/d2r2/go-logger"
)

// You can manage verbosity of log output
// in the package by changing last parameter value
// (comment/uncomment corresponding lines).
var logLevel, err = strconv.Atoi(os.Getenv("GO_I2C_LOG"))
var lg = logger.NewPackageLogger("i2c", logLevel)
