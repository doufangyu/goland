//程序所属包
package main
//导入依赖包
import "fmt"
//定义常量
const name string = "imooc"
//全局变量的声明和赋值
var a string  =  "窦芳玉"
//一般数据类型声明
type imoocInt int
//结构声明
type Learn struct {

}
//声明接口
type Ilearn interface {

}
//函数定义
func learImooc()  {
	fmt.Println("learImooc")

}

func main() {
	fmt.Println(a)
	learImooc()
	fmt.Println("const is ,",name)
}
