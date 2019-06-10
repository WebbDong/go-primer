package main

import (
	"fmt"
	"strings"
)

type Person struct {
	name   string
	age    int
	addr   string
	mobile string
}

var persons []Person

// 通讯录案例练习
func main() {
	var operate string
	for {
		printSelectOperate()
		fmt.Scan(&operate)
		if "Q" == strings.ToUpper(operate) {
			break
		}
		choseOperate(operate)
	}
}

func choseOperate(o string) {
	switch o {
	case "1":
		// 添加联系人
		addOperate()
	case "2":
		// 删除联系人
		deleteOperate()
	case "3":
		// 查询联系人
		queryOperate()
	case "4":
		// 编辑联系人
		editOperate()
	case "5":
		// 列出所有联系人
		listAllPersons()
	default:
		fmt.Println("未知操作")
	}
}

// 添加联系人
func addOperate() {
	newPerson := Person{}
	fmt.Println("添加联系人")
	fmt.Println("输入姓名：")
	fmt.Scan(&newPerson.name)
	fmt.Println("输入年龄：")
	fmt.Scan(&newPerson.age)
	fmt.Println("输入手机号：")
	fmt.Scan(&newPerson.mobile)
	fmt.Println("输入地址：")
	fmt.Scan(&newPerson.addr)
	persons = append(persons, newPerson)
	fmt.Println("添加成功!!!")
	pressAny()
}

// 删除联系人
func deleteOperate() {
	fmt.Println("输入要删除的姓名：")
	var name string
	fmt.Scan(&name)
	for i, v := range persons {
		if v.name == name {
			persons = append(persons[:i], persons[i+1:]...)
			break
		}
	}
}

// 查询联系人
func queryOperate() {
	fmt.Println("输入要查询的姓名：")
	var name string
	fmt.Scan(&name)
	for _, v := range persons {
		if v.name == name {
			fmt.Println("-------------- 查询联系人 --------------")
			fmt.Printf("姓名：%s，年龄：%d，手机号：%s，地址：%s\n", v.name, v.age, v.mobile, v.addr)
			fmt.Println("----------------------------------------")
			break
		}
	}
}

// 编辑联系人
func editOperate() {
	fmt.Println("输入要编辑的姓名：")
	var name string
	fmt.Scan(&name)
	for i, v := range persons {
		if v.name == name {
			editPerson(i)
			break
		}
	}
}

// 列出所有联系人
func listAllPersons() {
	if len(persons) == 0 {
		fmt.Println("没有联系人")
		return
	}
	fmt.Println("-------------- 所有联系人 --------------")
	for _, v := range persons {
		fmt.Printf("姓名：%s，年龄：%d，手机号：%s，地址：%s\n", v.name, v.age, v.mobile, v.addr)
	}
	fmt.Println("----------------------------------------")
	pressAny()
}

// 修改联系人
func editPerson(index int) {
	printEditOperate()
	var editOperate string
	fmt.Scan(&editOperate)
	fmt.Println("输入新值：")
	switch editOperate {
	case "1":
		// 修改姓名
		fmt.Scan(&persons[index].name)
		fmt.Println("修改成功")
	case "2":
		// 修改年龄
		fmt.Scan(&persons[index].age)
		fmt.Println("修改成功")
	case "3":
		// 修改手机号
		fmt.Scan(&persons[index].mobile)
		fmt.Println("修改成功")
	case "4":
		// 修改地址
		fmt.Scan(&persons[index].addr)
		fmt.Println("修改成功")
	default:
		fmt.Println("退出编辑")
		return
	}
}

func printSelectOperate() {
	fmt.Println("添加联系人，请按1")
	fmt.Println("删除联系人，请按2")
	fmt.Println("查询联系人，请按3")
	fmt.Println("编辑联系人，请按4")
	fmt.Println("列出所有联系人，请按5")
	fmt.Println("退出程序请按Q")
	fmt.Println("-----------------------------")
}

func printEditOperate() {
	fmt.Println()
	fmt.Println("修改姓名，请按1")
	fmt.Println("修改年龄，请按2")
	fmt.Println("修改手机号，请按3")
	fmt.Println("修改地址，请按4")
	fmt.Println("其他退出编辑")
	fmt.Println("-----------------------------")
}

func pressAny() {
	//fmt.Println("按下任意按键继续")
}
