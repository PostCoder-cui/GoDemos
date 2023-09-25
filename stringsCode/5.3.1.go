//package main
//
//import "fmt"
//
//func main() {
//	str:="hello word 你好 世界ismale"
//	fmt.Println(ReverseStr(str))
//}
//
//func ReverseStr(str string)string{
//	newStr:=[]rune(str)
//
//	for i,j:=0,len(newStr)-1;i<j;i,j=i+1,j-1{
//		newStr[i]=newStr[i]^newStr[j]  //此时的newStr[i]已经变了
//		newStr[j]=newStr[i]^newStr[j]  //此时的newStr[j]已经变了
//		newStr[i]=newStr[i]^newStr[j]  //新的newStr[i]与newStr[j]进行异或操作
//	}
//	return  string(newStr)
//}