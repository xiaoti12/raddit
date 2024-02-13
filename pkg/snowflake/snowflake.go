package snowflake

import (
	"github.com/bwmarrin/snowflake"
	"time"
)

var node *snowflake.Node

func Init(startTime string, machineID int) error {
	var stime time.Time
	// 这里需要正确设置layout的值，似乎有一定限制
	stime, err := time.Parse("2006-01-02", startTime)
	if err != nil {
		return err
	}
	snowflake.Epoch = stime.UnixNano() / 1000000
	node, err = snowflake.NewNode(int64(machineID))
	return nil
}
func GenID() int64 {
	return node.Generate().Int64()
}
