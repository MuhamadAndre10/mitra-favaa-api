package environment

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadEnv(t *testing.T) {

	cfg, err := LoadEnv()

	assert.Nil(t, err)

	assert.Equal(t, cfg.GetString("DB_HOST"), "localhost")

}
