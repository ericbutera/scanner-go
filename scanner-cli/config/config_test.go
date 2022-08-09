package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/mock"
)

// type MockedConfig struct {
// }
// func (m *MockedConfig) GetString(key string) string {
// 	return "test"
// }

func TestConfig(t *testing.T) {
	assert.Equal(t, "one", "one")
}
