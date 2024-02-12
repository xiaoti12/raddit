package snowflake

import (
	"github.com/bwmarrin/snowflake"
	"time"
)

var node *snowflake.Node

func Init(startTime string, machineID int) {
	var stime time.Time
	stime, err := time.Parse("2000-01-01", startTime)
	if err != nil {
		return
	}
	snowflake.Epoch = stime.UnixNano() / 1000000
	node, err = snowflake.NewNode(int64(machineID))
	return
}
func GenID() int64 {
	return node.Generate().Int64()
}
