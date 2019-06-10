package main

import "fmt"

func main() {
	//oo()
	//extends()
	//memberOperate()
	//pointerExtends()
	//nestExtends()
	//multipleExtends()
	memberFunc()
}

// GO语言使用结构体定义类
type Car struct {
	horsepower int
	brand      string
	topSpeed   int
}

// 面向对象
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

// 继承
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

// 成员操作
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

// 指针类型匿名字段继承
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

// 多重继承
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

// 多继承
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

// 结构体或类的成员方法
func memberFunc() {
	c1 := Computer{10, 8, 1024}
	c1.run()
}
