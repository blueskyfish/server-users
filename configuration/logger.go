package configuration

import "github.com/sirupsen/logrus"

// WithGroup create the logger with the given group
func (conf Configuration) WithGroup(value interface{}) *logrus.Entry {
	return conf.log.WithField("group", value)
}
