package suid

import (
	"log"

	"github.com/bwmarrin/snowflake"
)

// Node 是雪花节点
var Node *snowflake.Node

func init() {
	//雪花算法init
	err := InitIDGen(1)
	if err != nil {
		panic(err)
	}
}

// InitIDGen 初始化雪花 ID 生成器
func InitIDGen(nodeID int64) error {
	var err error
	Node, err = snowflake.NewNode(nodeID)
	if err != nil {
		return err
	}
	return nil
}

// GenerateID 生成一个新的唯一 ID
func GenerateID() string {
	if Node == nil {
		log.Fatal("Snowflake Node is not initialized")
	}
	return Node.Generate().String()
}
