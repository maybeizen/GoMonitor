package outputs

import (
	"monitor/models"
)

type OutputHandler interface {
	Write(info *models.SystemInfo) error

	Name() string
}

func NewOutputHandler(config models.OutputConfig) (OutputHandler, error) {
	switch config.Type {
	case models.FileOutput:
		return NewFileOutputHandler(config.FilePath)
	case models.APIOutput:
		return NewAPIOutputHandler(config.APIURL, config.APIKey, config.APIMethod)
	default:
		return NewFileOutputHandler(config.FilePath)
	}
} 