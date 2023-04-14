package utils

import (
	"log"
)

func ExtractRoles(realm_access map[string]interface{}) []string {
	var roleList []string
	for k, v := range realm_access {
		if k == "roles" {
			for _, value := range v.([]interface{}) {
				roleList = append(roleList, value.(string))
			}
		}
	}

	return roleList
}

// Log is an interface for logging.
type Log interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Printf(format string, args ...interface{})
}

type Logger struct {
	Debug bool
}

// Log to stdout without timestamp
func (logger Logger) Printf(format string, args ...interface{}) {
	log.Fatalf("auth: "+format, args...)
}

// Log to stdout only if Debug is true.
func (logger Logger) Debugf(format string, args ...interface{}) {
	if logger.Debug {
		log.Printf(format+"\n", args...)
	}
}

// Log to stdout only if Debug is true.
func (logger Logger) Infof(format string, args ...interface{}) {
	if logger.Debug {
		log.Printf(format+"\n", args...)
	}
}

// Log to stdout always.
func (logger Logger) Warnf(format string, args ...interface{}) {
	log.Printf(format+"\n", args...)
}

// Log to stdout always.
func (logger Logger) Errorf(format string, args ...interface{}) {
	log.Printf(format+"\n", args...)
}
