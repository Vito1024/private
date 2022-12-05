package main

import "fmt"

var m = map[string]string{
"请开启前照灯": "近光灯",
"通过路口": "近光灯",
"夜间与对方会车，距离对方来车150米": "近光灯",
"夜间同方向跟车行驶时": "近光灯",
"进入有路灯，照明良好路段": "近光灯",
"夜间没有路灯，照明不良条件下行驶": "远光灯",
"夜间通过急弯，拱桥，坡路，人行横道或没有信号灯控制的路口":"交替远近灯",
"前方开始超车":"交替远近灯",
"在路边临时停车": "应急灯",
}

func main() {
	test()
}

func test() {
	correctCount := 0
	for k, v := range m {
		fmt.Println(k)
		var input string
		fmt.Scanf("%s", &input)
		if input == v {
			fmt.Println("correct")
			correctCount++
			continue
		}
		fmt.Println("not correct,", v)
	}
	fmt.Println("your score：", correctCount * 10)
}
