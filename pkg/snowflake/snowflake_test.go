package snowflake

import "testing"

func TestGenID(t *testing.T) {
	Init("2024-01-01", 1)
	t.Log(GenID())
}
