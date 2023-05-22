package main

import "fmt"

func ArraySlice() {
	// 声明切片
	var s1 []int
	if s1 == nil {
		fmt.Println("切片为空")
	} else {
		fmt.Println("切片不为空")
	}

	// 发现一个问题没有，上面只做了声明的切片的值是nil，但是这里还是可以获取其长度和容量
	// 其实切片的本质是一个结构体，有一个指针是指向一个数组的，还有len、cap属性
	fmt.Printf("只做了声明的切片的长度:%d和容量:%d\n", len(s1), cap(s1))

	s2 := []int{}
	var s3 []int = make([]int, 0)
	fmt.Println(s1, s2, s3)

	// 初始化赋值
	var s4 []int = make([]int, 0, 0)
	fmt.Println(s4)

	s5 := []int{1, 2, 3}
	fmt.Println(s5)

	// 切片赋值为数组的引用
	arr := [5]int{1, 2, 3, 4, 5}
	var s6 []int
	// 切片引用数组从下标1到4的片段，包括1、不包括4
	s6 = arr[1:4]
	fmt.Println(s6)

	// 改变数组下标为1的数，切片的值也变了
	arr[1] = 7
	fmt.Println(s6)

	fmt.Println(nil)

	arrData := [...]int{0, 1, 2, 3, 4, 10: 0}
	s7 := arrData[2:3]  // 元素是2，长度是3-2=1，容量是从下标为2开始，一直到11
	s8 := arrData[:2:3] // 元素是0、1，长度是2，容量是3
	fmt.Println(s7, len(s7), cap(s7))
	fmt.Println(s8, len(s8), cap(s8))

	// 当向切片中添加数据超过切片的cap限制，就会重新分配底层数组，那么切片就会指向新分配的数组，和原来的数组没关系了
	arrData2 := [...]int{0, 1, 2, 3, 4, 10: 0}
	s9 := arrData2[:2:3]      // 切片的长度是2，容量是3，底层数组指向的是arrData2的第一个和第二个元素
	s9 = append(s9, 100, 200) // 一次append两个值，超出s.cap限制
	fmt.Println(s9, arrData2) // 重新分配底层数组，与原数组无关
	fmt.Println(&s9[0], &arrData2[0])

	var arr3 = []int{1, 3, 4, 5}
	fmt.Printf("slice arr3: %v, len(arr3) : %d, cap(arr3) : %d\n", arr3, len(arr3), cap(arr3))
	s10 := arr3[1:2]
	fmt.Printf("slice s10: %v, len(s10) : %d, cap(s10) : %d\n", s10, len(s10), cap(s10))
	s11 := s10[0:3]
	fmt.Printf("slice s11: %v, len(s11) : %d, cap(s11) : %d\n", s11, len(s11), cap(s11))

	var p *string
	fmt.Printf("p的值是:%v, 类型是:%T", p, p)

	//var a *int
	//*a = 100
	//fmt.Println(*a)

	var b map[string]int
	b["测试"] = 100
	fmt.Println(b)

	arrA := [5]int{1, 3, 4}
	arrB := arrA
	fmt.Printf("数组arrA的第一个元素的地址：%p", &arrA[0])
	fmt.Printf("数组arrB的第一个元素的地址：%p", &arrB[0])

	sliceA := []int{1, 3, 4}
	sliceB := sliceA
	fmt.Printf("切片sliceA的第一个元素的地址：%x", &sliceA[0])
	fmt.Printf("切片sliceB的第一个元素的地址：%x", &sliceB[0])
}

// SliceReference
// 切片在函数传递中，是值传递还是引用传递
// 如果在函数中，切片没有发生扩容，修改了切片，外面的切片受影响了吗。如果切片发生了扩容，修改里面的切片，外面的切片受影响了吗
func SliceReference() {
	outerSlice := make([]int, 0, 6)
	fmt.Printf("outerSlice地址：%p\n", &outerSlice)

	func(sli []int) {
		fmt.Printf("outerSlice地址：%p\n", &sli)
		for i := 0; i < len(sli)+6; i++ {
			sli = append(sli, i)
		}
		fmt.Printf("outerSlice地址：%p\n", &sli)
	}(outerSlice)

	fmt.Printf("outerSlice地址：%p\n", &outerSlice)

}

func SliceInitiation() {
	var sli []int
	fmt.Printf("切片：%v，长度：%d，容量：%d\n", sli, len(sli), cap(sli))
	for i := 0; i < 6; i++ {
		sli = append(sli, i)
	}
	fmt.Printf("切片：%v，长度：%d，容量：%d\n", sli, len(sli), cap(sli))

	makeSli := make([]int, 0, 0)
	fmt.Printf("make创建的切片：%v，长度：%d，容量：%d\n", makeSli, len(makeSli), cap(makeSli))
	for i := 0; i < 6; i++ {
		makeSli = append(makeSli, i)
	}
	fmt.Printf("make创建的切片：%v，长度：%d，容量：%d\n", makeSli, len(makeSli), cap(makeSli))
}
