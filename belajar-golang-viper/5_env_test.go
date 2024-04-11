package belajar_golang_viper

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestENV(t *testing.T) {
	var config *viper.Viper = viper.New()
	config.AutomaticEnv()

	assert.Equal(t, "Hello", config.GetString("FROM_ENV"))
}
