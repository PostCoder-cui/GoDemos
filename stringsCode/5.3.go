//反转
//package main
//
//import "fmt"
//
//func main() {
//	str:="你好吗朋友 is good ja"
//	fmt.Println(ReseverStr(str))
//}
//
//func ReseverStr(str string)string{
//	newStr:=[]rune(str)
//	fmt.Println("length:",len(newStr))
//	for i,j:=0,len(newStr)-1;i<j;i,j=i+1,j-1{
//
//		fmt.Println("i:",i)
//		fmt.Println("j:",j)
//		temp:=newStr[i]
//		newStr[i]=newStr[j]
//		newStr[j]=temp
//	}
// fmt.Println("[]rune:",newStr)
//	return string(newStr)
//}