//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	//var str = "hello wprd"
//	//var flag=make([]int32,8)
//	//for key, value := range str {
//	//	fmt.Println("====start=====")
//	//	fmt.Println("key:",key)
//	//	index:=value/32     //因为每一个数组元素是一个int32，也就是32位，这样可以将字符分配到8个int中的对应int位置，也就是flag数组的index
//	//	shift:=uint(value)%32  //偏移量，也就是距离自己那个32的距离，距离自己flag[index]的距离
//	//	fmt.Println("index:",index)
//	//	fmt.Println("shift:",shift)
//	//	fmt.Println("flag:",flag[index])
//	//	fmt.Println("1<<shift:",1<<shift)  //换一种写法，就是32bit只有高位1,2147483648-1右移>>shift
//	//	flag[index]|=1<<shift //1<<shift表示1从最右边向左偏移shift个位置
//	//	fmt.Println("flag:",flag[index])
//	//	fmt.Println("=====end======")
//	//	//fmt.Println(104/32)   //这个地方主要为了提示一个东西：golang整数相除结果会舍弃小数位
//	//}
//	fmt.Println(isDup("hello word"))
//	fmt.Println(isDup("helo word"))
//	fmt.Println(isDup("dd"))
//
//	fmt.Println(isDup("ds"))
//}
//
////func isDuplicate(str string)bool{
////	for i:=0;i<len(str);i++{
////		for j:=i+1;j<len(str);j++{
////			if str[i]==str[j]{
////				return true
////			}
////		}
////	}
////	return false
////}
//
//func isDup(str string) bool {
//	var flag = make([]int32, 8)
//	for _, v := range str {
//		index := v / 32
//		shift := uint(v) % 32
//		if (flag[index] & (1 << shift)) != 0 { //只有1&1才会==1，说明有一个位重复了
//			return true
//		}
//		flag[index] |= 1 << shift
//	}
//	return false
//}
