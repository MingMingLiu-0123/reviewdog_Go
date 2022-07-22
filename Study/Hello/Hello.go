package main // 声明 main 包，表明当前是一个可执行程序

import (
	"fmt" // 导入内置 fmt 包
)

func main() { // main函数，是程序执行的入口
	fmt.Println("Hello World!" + "哈哈哈h!") // 在终端打印 Hello World!

	a1 := "111"
	fmt.Println("hello world", a1)
	/* 定义局部变量 */
	var a int = 100
	var b int = 200
	// var ret int

	/* 调用函数并返回最大值 */
	ret := max(a, b)

	fmt.Printf("最大值是 : %d\n", ret)

	c, d := swap("交换", "数据")
	fmt.Println(c, d)

	balance := [8]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	for i, x := range balance { //在数组上使用range将传入index和值两个变量。
		fmt.Printf("Element[%d] = %.2f\n", i, x)
	}

	myMap()

	var i int = 15
	fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(uint64(i)))
}

/* 函数返回两个数的最大值 */
func max(num1, num2 int) int {
	/* 定义局部变量 */
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

func swap(x, y string) (string, string) {
	return y, x
}

func myMap() {
	var countryCapitalMap map[string]string /*创建集合 */
	countryCapitalMap = make(map[string]string)

	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India "] = "新德里"

	/*使用键输出地图值 */
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

	/*查看元素在集合中是否存在 */
	capital, ok := countryCapitalMap["American"] /*如果确定是真实的,则存在,否则不存在 */
	/*fmt.Println(capital) */
	/*fmt.Println(ok) */
	if ok {
		fmt.Println("American 的首都是", capital)
	} else {
		fmt.Println("American 的首都不存在")
	}
}

func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}
