package conf

import (
	"encoding/json"
	"os"

	"github.com/rashintha/logger"
)

var DATA map[string]interface{}

func init() {
	logger.Defaultln("Loading configuration variables")

	file, err := os.Open(".conf")
	if err != nil {
		logger.Errorf("Error in reading configuration file: %v", err.Error())
	}
	defer file.Close()

	DATA = make(map[string]interface{})

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&DATA)
	if err != nil {
		logger.Errorf("Error in decoding configuration file: %v", err.Error())
	}
}
