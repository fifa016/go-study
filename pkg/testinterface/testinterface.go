/**
 * 值类型拥有值方法，指针类型拥有全部方法
 */
package main

import "fmt"

type RssMatcher struct{}

type Matcher interface {
	pMatch()
	valueMatch()
}

func (r *RssMatcher) pMatch() {
	fmt.Println("p match")
}

func (r RssMatcher) valueMatch() {
	fmt.Println("value match")
}

func main() {
	fmt.Println("=========start run=========")

	//part1 值类型和指针类型会互相转
	rssMatcher := RssMatcher{}
	rssMatcher2 := &RssMatcher{}

	fmt.Println("test value and point method")
	rssMatcher.pMatch()
	rssMatcher.valueMatch()
	rssMatcher2.pMatch()
	rssMatcher2.valueMatch()
	fmt.Println()

	//part2 指针
	var rssMatcher3 RssMatcher  //值
	var rssMatcher4 *RssMatcher //指针
	rssMatcher3 = rssMatcher
	//rssMatcher3 = &rssMatcher //值类型变量不能指向指针 done
	//rssMatcher3 = rssMatcher2 //指针类型不能指向值类型 done
	rssMatcher3 = *rssMatcher2
	rssMatcher4 = &rssMatcher
	rssMatcher4 = rssMatcher2
	fmt.Println("test pointer")
	fmt.Println(rssMatcher3)
	fmt.Println(rssMatcher4)
	fmt.Println()

	//part3 接口
	var matcher Matcher
	//matcher = rssMatcher //rssmatcher为值类型 没有pmatch方法
	matcher = &rssMatcher
	matcher = rssMatcher2
	//matcher = *rssMatcher2 //*rssmatcher2为值类型 没有pmatch方法
	fmt.Println("test interface")
	matcher.pMatch()
	matcher.valueMatch()
	fmt.Println()

    // 没有这么做的
	//var matcher2 *Matcher
	//matcher2 = rssMatcher
	//matcher2 = &rssMatcher
	//matcher2 = rssMatcher2
	//matcher2 = &rssMatcher2

}
