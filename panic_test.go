package stringx

import (
	"log"
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestFail(t *testing.T) {
	failedMessage := "index out of range"
	result := String("Hello World")
	assert.Equal(t, Log, log.Println)
	fail(failedMessage, result, func(v ...any) {
		assert.Equal(t, failedMessage, v[0])
	})
}
