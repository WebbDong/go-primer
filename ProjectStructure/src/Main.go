package main

// 直接根据$GOPATH/src目录导入，自己的包和go文件，需要在IDE中配置GOPATH路径。
import (
	"fmt"
	"model"
)

// 相对路径导入，不建议使用这种导入方式
// 起别名，biz为新起的别名
import biz "./service"

// 使用.省略包名，直接使用
import . "controller"

// 该操作其实只是引入该包。当导入一个包时，它所有的init()函数就会被执行，但有些时候并非真的需要使用这些包，
// 仅仅是希望它的init()函数被执行而已。这个时候就可以使用_操作引用该包。即使用_操作引用包是无法通过包名来调用包中的导出函数，
// 而是只是为了简单的调用其init函数()。往往这些init函数里面是注册自己包里面的引擎，让外部可以方便的使用，
// 例如实现database/sql的包，在init函数里面都是调用了sql.Register(name string, driver driver.Driver)注册自己，
// 然后外部就可以使用了。
import _ "dao"

// ide默认运行时，只会编辑当前的go文件，如果要编译整个src目录下的go文件，需要进行配置
func main() {
	fmt.Println(Add(10, 20))
	p1 := model.CreatePerson(10, "Ferrari", 90)
	p1.Walk()
	p1.Run()
	p1.Speak()
	fmt.Println("p1.GetName() =", p1.GetName())
	model.Func2()
	fmt.Println("------------------------------------")

	t1 := model.CreateTeacher(11, "Lamborghini", 99)
	t1.Speak()
	t1.Teach()
	fmt.Println("------------------------------------")

	biz.Login()

	Register()
	fmt.Println("------------------------------------")

	var pi model.Personer = p1
	pi.Speak()
}
