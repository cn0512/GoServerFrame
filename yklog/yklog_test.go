package yklog

import "testing"

func TestLogout(t *testing.T) {
	Logout("test", 1, "2")
}

func BenchmarkLogout(b *testing.B) {
	Logout("benchmark", 1, "2")
}
