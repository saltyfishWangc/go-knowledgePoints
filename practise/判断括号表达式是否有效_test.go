package practise

import (
	"testing"
)

// 1、（（ 结果是false
// 2、（（））（） 是true
// 3、（（（）））（））是false  用go实现

const (
	startByte = '('
	endByte   = ')'
)

func TestStackUsage(t *testing.T) {
	//arr := []byte{'(', '('}
	arr := []byte{'(', '(', ')', ')', ' ', '(', ')'}
	//arr := []byte{'(', '(', '(', ')', ')', ')', '(', ')', ')'}
	t.Logf("切片：%v 是否有效：%t", arr, arrayIsValid(arr))
}

func arrayIsValid(arr []byte) bool {
	stack := NewStack()
	for _, v := range arr {
		if stack.isEmpty() && endByte == v {
			// 在栈为空的情况下遇到了)，那就是直接返回false
			return false
		}
		// 判断栈是否是空栈，是的话直接入栈
		if stack.isEmpty() {
			stack.push(v)
		} else {
			// 遇到(就要入栈
			if startByte == v {
				stack.push(v)
			} else if endByte == v {
				// 遇到)就要弹出栈顶元素
				data := stack.pop()
				// 因为数组中是用interface接收的，用于直接用==，一定是不相等的，因为startByte和data类型不同
				if startByte != data.(byte) {
					return false
				}
			}
		}
	}
	// 最后判断栈是否清空了，没有的话说明还有(没消除，直接返回false
	if stack.isEmpty() {
		return true
	} else {
		return false
	}
}

// 思路：设计一个栈的数据结构出来，遇到(做入栈操作，遇到)则弹出栈顶元素判断是否是(
type Stack struct {
	Data []interface{}
	Size int
}

func NewStack() *Stack {
	return &Stack{Size: 0}
}

// push 入栈
func (stack *Stack) push(data interface{}) {
	stack.Data = append(stack.Data, data)
	stack.Size += 1
}

// pop 出栈
//
//	func (stack *Stack) pop() (interface{}, error) {
//		if stack.Size == 0 {
//			return 0, errors.New("空栈")
//		}
//		data := stack.Data[stack.Size-1]
//		stack.Size -= 1
//		return data, nil
//	}
func (stack *Stack) pop() interface{} {
	if stack.isEmpty() {
		return nil
	}
	data := stack.Data[stack.Size-1]
	stack.Size -= 1
	return data
}

func (stack *Stack) isEmpty() bool {
	return stack.Size == 0
}
