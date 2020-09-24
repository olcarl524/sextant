package sextant

import (
	"github.com/cockroachdb/pebble"
	"go.uber.org/zap"
)

/*
Definition has a logger field for uber's Zap logger object and db as the db path for pebble.
*/
type Definition struct {
	logger *zap.Logger
	db     string
}

// New Returns a new Definition object
func New() *Definition {
	l, _ := zap.NewProduction()
	return &Definition{logger: l}
}

// WithLogContext Appends the log context of upper contexts to sextant context
func (c *Definition) WithLogContext(logContext *zap.Logger) *Definition {
	c.logger = logContext
	return &*c
}

// WithDatabase Defines the pebble database path to be opened
func (c *Definition) WithDatabase(database string) *Definition {
	c.db = database
	return &*c
}

func (c *Definition) logErrorIfAny(err error) {
	if err != nil {
		c.logger.Error(err.Error())
	}
}

/*
Set will create a databse if it doesn't exist or open an existing one
set the key k to value v and close the database
TODO: refactor arguments to interface and reflect it's type so it doesn't use
string as types
*/
func (c *Definition) Set(k string, v string) error {
	db, err := pebble.Open(c.db, &pebble.Options{})
	c.logErrorIfAny(err)
	err = db.Set([]byte(k), []byte(v), pebble.Sync)
	c.logErrorIfAny(err)
	err = db.Close()
	c.logErrorIfAny(err)
	return err
}

/*
Get will create a databse if it doesn't exist or open an existing one
get the key k and return the value as string
TODO: refactor arguments to interface and reflect it's type so it doesn't use
string as types
*/
func (c *Definition) Get(k string) (string, error) {
	db, errOpen := pebble.Open(c.db, &pebble.Options{})
	c.logErrorIfAny(errOpen)
	downstreamResponse := ""
	response, closer, err := db.Get([]byte(k))
	if response != nil {
		downstreamResponse = string(response)
	}
	c.logErrorIfAny(err)
	err = closer.Close()
	c.logErrorIfAny(err)
	defer db.Close()
	return downstreamResponse, err
}
