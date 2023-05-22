package data_structure

//
////package main
//
//import (
//	"encoding/json"
//	"fmt"
//	"math/rand"
//	"time"
//)
//
//const MaxLevel = 32
//const p = 0.5
//
//type Node struct {
//	value  uint32
//	levels []*Level //索引节点，index=0是基础链表
//}
//
//type Level struct {
//	next *Node
//}
//
//type SkipList struct {
//	header *Node  //表头节点
//	length uint32 //原始链表的长度，表头节点不计入
//	height uint32 //最高的节点的层数
//}
//
//func NewSkipList() *SkipList {
//	return &SkipList{
//		header: NewNode(MaxLevel, 0),
//		length: 0,
//		height: 1,
//	}
//}
//
//func NewNode(level, value uint32) *Node {
//	node := new(Node)
//	node.value = value
//	node.levels = make([]*Level, level)
//
//	for i := 0; i < len(node.levels); i++ {
//		node.levels[i] = new(Level)
//	}
//	return node
//}
//
//// randomLevel 随机函数 生成新节点的层数
//func (sl *SkipList) randomLevel() int {
//	level := 1
//	r := rand.New(rand.NewSource(time.Now().UnixNano()))
//	for r.Float64() < p && level < MaxLevel {
//		level++
//	}
//	return level
//}
//
//// Add 插入元素
///**
//1 		 ->			10
//1 	->	  7   ->	10    ->    18
//1 -> 4 -> 7 -> 8 -> 10 -> 15 -> 18 -> 21
//*/
//func (sl *SkipList) Add(value uint32) bool {
//	if value <= 0 {
//		return false
//	}
//	update := make([]*Node, MaxLevel)
//	// 每一次循环都是一次寻找有序单链表的插入过程
//	tmp := sl.header
//	for i := int(sl.height) - 1; i >= 0; i-- {
//		// 每次循环不重置tmp，直接从上一层确认的节点开始向下一层查找
//		for tmp.levels[i].next != nil && tmp.levels[i].next.value < value {
//			tmp = tmp.levels[i].next
//		}
//
//		// 避免插入相同的元素
//		if tmp.levels[i].next != nil && tmp.levels[i].next.value == value {
//			return false
//		}
//
//		update[i] = tmp
//	}
//
//	level := sl.randomLevel()
//	node := NewNode(uint32(level), value)
//	fmt.Printf("新加入的节点: %v 所在的层数是 level: %v\n", value, level)
//
//	if uint32(level) > sl.height {
//		sl.height = uint32(level)
//	}
//
//	for i := 0; i < level; i++ {
//		//说明新节点层数超过了跳表当前的最高层数，此时将头节点对应层数的后继节点设置为新节点
//		if update[i] == nil {
//			sl.header.levels[i].next = node
//			continue
//		}
//
//		//普通的插入链表操作，修改新节点的后继节点为前一个节点的后继节点，修改前一个节点的后继节点为新节点
//		node.levels[i].next = update[i].levels[i].next
//		update[i].levels[i].next = node
//	}
//
//	sl.length++
//	return true
//}
//
//// Find 查找值对应的节点
///**
//1 		 ->			10
//1 	->	  7   ->	10    ->    18
//1 -> 4 -> 7 -> 8 -> 10 -> 15 -> 18 -> 21
//*/
//func (sl *SkipList) Find(value uint32) *Node {
//	var node *Node
//	tmp := sl.header
//	for i := int(sl.height) - 1; i >= 0; i-- {
//		for tmp.levels[i].next != nil && tmp.levels[i].next.value <= value {
//			tmp = tmp.levels[i].next
//		}
//
//		if tmp.value == value {
//			node = tmp
//			break
//		}
//	}
//	return node
//}
//
//func (sl *SkipList) Print() {
//	tmp := sl.header
//	for i := int(sl.height) - 1; i >= 0; i-- {
//		tmp = tmp.levels[i].next
//		var line string
//		line = fmt.Sprintf("%d", int(tmp.value))
//		for tmp.levels[i].next != nil {
//			tmp = tmp.levels[i].next
//			line = fmt.Sprintf("%s -> %d", line, int(tmp.value))
//		}
//		fmt.Println(line)
//	}
//}
//
///*
//*
//1 		 ->			10
//1 	->	  7   ->	10    ->    18
//1 -> 4 -> 7 -> 8 -> 10 -> 15 -> 18 -> 21
//*/
//func main() {
//	skipList := NewSkipList()
//	skipList.Add(1)
//	skipList.Add(4)
//	skipList.Add(7)
//	skipList.Add(8)
//	skipList.Add(10)
//	skipList.Add(15)
//	skipList.Add(18)
//	skipList.Add(21)
//
//	//fmt.Printf("%+v", skipList.Find(10))
//	//skipList.Print()
//	marshal, err := json.Marshal(skipList.header.levels)
//	if err != nil {
//		fmt.Println("出错")
//	} else {
//		fmt.Println(string(marshal))
//	}
//}
