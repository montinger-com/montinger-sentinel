package config

import (
	"github.com/montinger-com/montinger-sentinel/lib/conf"
	"github.com/rashintha/logger"
)

var API_URL string
var UID string
var SECRET string

func init() {
	logger.Defaultln("initializing configuration variables")

	API_URL = conf.DATA["api_url"].(string)
	UID = conf.DATA["uid"].(string)
	SECRET = conf.DATA["secret"].(string)
}
