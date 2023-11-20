package test

import (
	"testing"
	"time"
)

const ttl = "20m"

func TestTime(t *testing.T) {
	t.Run("test time", func(t *testing.T) {
		duration, _ := time.ParseDuration(ttl)
		now := time.Now().Add(duration)
		t.Log(now)
	})
}
