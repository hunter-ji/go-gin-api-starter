// @Title snow_flake.go
// @Description
// @Author Hunter 2024/9/3 17:23

package snowFlake

import "github.com/bwmarrin/snowflake"

var (
	Node *snowflake.Node
)

func init() {
	node, err := snowflake.NewNode(0)
	if err != nil {
		panic(err)
	}

	Node = node
}

func GenStringID() string {
	return Node.Generate().String()
}
