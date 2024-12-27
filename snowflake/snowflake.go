package main

import (
	"fmt"
	"strconv"

	"github.com/bwmarrin/snowflake"
)

const (
	idLength = 8
)

const (
	MaxNodeId = 1024
)

var (
	nodeMap = make(map[int]*snowflake.Node, 0)
)

func getNodeByRef(ref string) *snowflake.Node {
	// if err,nodeId = 0
	nodeId, _ := strconv.Atoi(ref)
	nodeId = nodeId % MaxNodeId

	if node, ok := nodeMap[nodeId]; ok {
		return node
	}
	node, _ := snowflake.NewNode(int64(nodeId))
	nodeMap[nodeId] = node
	return node
}

// GenerateUniqueID 生成一个唯一的 Snowflake ID，并将其转换为 Base62 编码，限制为 8 个字符
func GenerateUniqueID(ref string) string {
	node := getNodeByRef(ref)
	id := node.Generate().Base32()
	if len(id) > idLength {
		// fmt.Printf("Original ID: %s\n", id)
		// use last 8 characters
		id = id[len(id)-idLength:]
	}
	return id
}

func main() {
	recordMap := make(map[string]bool, 0)
	ref := 101568720036448
	for i := 0; i < 999999999; i++ {
		ref++
		id := GenerateUniqueID(strconv.Itoa(ref))
		// fmt.Printf("ID: %s\n", id)
		if _, ok := recordMap[id]; ok {
			fmt.Printf("Duplicate ID, tried %d times\n", i)
			panic("Duplicate ID")
		} else {
			recordMap[id] = true
		}
	}
	fmt.Println("All IDs are unique")
}
