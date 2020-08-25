package main

import "fmt"

// 面向对象，类，对象，继承
func main() {
	//oo()
	//extends()
	//memberOperate()
	//pointerExtends()
	//nestExtends()
	//multipleExtends()
	//memberMethod()
	//methodExtends()
	//methodOverride()
	//methodValueAndExpression()
	test()
}

// GO语言使用结构体定义类
type Car struct {
	horsepower int
	brand      string
	topSpeed   int
}

// 1、面向对象
func oo() {
	var c1 = Car{800, "Ferrari", 350}
	c2 := Car{horsepower: 1500, brand: "bugatti", topSpeed: 420}
	fmt.Println(c1)
	fmt.Println(c2)
}

type Person struct {
	name string
	age  int
}

type Student struct {
	// 匿名变量，实现继承
	Person
	score float64
}

type Teacher struct {
	// 匿名变量，实现继承
	Person
	salary float64
}

// 2、继承
func extends() {
	var s1 = Student{Person{"哈登", 30}, 96}
	var s2 = Student{score: 70}
	s3 := Student{Person: Person{name: "麦迪"}}
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

	t1 := Teacher{Person{name: "欧文"}, 8000}
	fmt.Println(t1)
}

// 3、成员操作
func memberOperate() {
	s1 := Student{Person{"科比", 39}, 100}
	fmt.Println(s1)
	fmt.Println("s1.Person.name =", s1.Person.name)
	fmt.Println("s1.Person.age =", s1.Person.age)
	fmt.Println("s1.name =", s1.name)
	fmt.Println("s1.age =", s1.age)
	fmt.Println("s1.score =", s1.score)

	s1.name = "科比二世"
	s1.Person.age = 20
	fmt.Println(s1)
}

// 4、指针类型匿名字段继承
func pointerExtends() {
	type Driver struct {
		// 指针类型匿名字段，实现继承
		*Person
		drivingYears int
	}

	// 需要使用取地址符
	d1 := Driver{&Person{"舒马赫", 50}, 20}
	fmt.Println(d1)
	fmt.Println("d1.Person.name =", d1.Person.name)
	fmt.Println("d1.Person.age =", d1.Person.age)

	fmt.Println("(*d1.Person).name =", (*d1.Person).name)
	fmt.Println("(*d1.Person).age =", (*d1.Person).age)

	fmt.Println("(*(d1.Person)).name =", (*(d1.Person)).name)
	fmt.Println("(*(d1.Person)).age =", (*(d1.Person)).age)

	fmt.Println("d1.name =", d1.name)
	fmt.Println("d1.age =", d1.age)

	fmt.Println("d1.drivingYears =", d1.drivingYears)
}

// 5、多重继承
func nestExtends() {
	type Vehicle struct {
		topSpeed   int
		brand      string
		horsepower int
	}

	type SaloonCar struct {
		Vehicle
		radio float64
	}

	type SuperCar struct {
		SaloonCar
		nitrogen int
	}

	v1 := SuperCar{SaloonCar{Vehicle{350, "Ferrari", 800}, 102.5}, 100}
	fmt.Println(v1)

	v2 := SuperCar{SaloonCar: SaloonCar{Vehicle: Vehicle{topSpeed: 420}}}
	fmt.Println(v2)
	fmt.Println("v2.SaloonCar.Vehicle.topSpeed =", v2.SaloonCar.Vehicle.topSpeed)
	fmt.Println("v2.topSpeed =", v2.topSpeed)
}

// 6、多继承
func multipleExtends() {
	type Children struct {
		name string
		age  int
	}

	type Worker struct {
		salary float64
	}

	type ChildLabourer struct {
		Children
		Worker
	}

	cl1 := ChildLabourer{Children{"DongWenBin", 18}, Worker{1000}}
	fmt.Println(cl1)
	fmt.Println(cl1.salary)
	fmt.Println(cl1.name)
	fmt.Printf("%T\n", cl1)
}

type Computer struct {
	cpu int
	gpu int
	ram int
}

func (c Computer) run() {
	fmt.Println("Computer running...")
	fmt.Println(c.cpu)
	fmt.Println(c.gpu)
}

func (c Computer) add(num1 int, num2 int) int {
	return num1 + num2
}

func (c Computer) setRam(ram int) {
	// 不使用指针无法修改
	c.ram = ram
}

func (c *Computer) setRam2(ram int) {
	// 使用指针才能修改
	c.ram = ram
}

// 7、结构体(类)的成员方法
func memberMethod() {
	c1 := Computer{10, 8, 1024}
	c1.run()
	fmt.Println("c1.add(10, 30) =", c1.add(10, 30))
	c1.setRam(200)
	fmt.Println(c1)
	c1.setRam2(5000)
	fmt.Println(c1)
	// 与c1.setRam2(5000)写法等价
	(&c1).setRam2(6000)
	fmt.Println(c1)
}

type Animal struct {
	weight int
	age    int
	height int
}

func (a *Animal) eat() {
	fmt.Println("eat")
}

type Tiger struct {
	Animal
}

// 8、方法继承
func methodExtends() {
	t1 := Tiger{Animal{100, 20, 2}}
	t1.eat()
}

type Parent struct {
}

func (p *Parent) PrintInfo() {
	fmt.Println("父类")
}

type Sub struct {
	Parent
}

func (s *Sub) PrintInfo() {
	fmt.Println("子类")
}

// 9、方法重写
func methodOverride() {
	s1 := Sub{}
	// 调用的是子类的
	s1.PrintInfo()
	// 如果要调用父类的，需要显示指定父类名
	s1.Parent.PrintInfo()
}

type Airplane struct {
	speed int
}

func (a *Airplane) fly() {
	fmt.Println(*a)
}

func (a Airplane) landing() {
	fmt.Println(a)
}

// 10、方法值与方法表达式
func methodValueAndExpression() {
	a1 := Airplane{700}
	// 方法值，将方法赋值给变量
	m1 := a1.fly
	// 调用
	m1()
	fmt.Printf("m1 = %T\n", m1)
	fmt.Println("----------------------------")

	// 方法表达式
	m2 := (*Airplane).fly
	m2(&a1)
	fmt.Printf("m2 = %T\n", m2)
	fmt.Println("----------------------------")

	m3 := Airplane.landing
	m3(a1)
	fmt.Printf("m3 = %T\n", m3)
}

// 为int类型取一个别名
type Int int

// 为Int类型添加方法
func (self Int) add(b Int) Int {
	return self + b
}

// 11、为go语言基本数据类型添加自定义方法
func test() {
	var num Int = 200
	fmt.Println(num.add(20))
}
