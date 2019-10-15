package main

import (
	"./model"
	"encoding/json"
	"fmt"
)

func main() {
	//objToJson()
	//jsonToObj()
	//jsonToMap()
	jsonToSlice()
}

// 对象转json字符串
func objToJson() {
	d1 := model.Doctor{Name: "d1", Age: 40, Address: "shanghai"}
	bytes, err := json.Marshal(d1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bytes))
}

// json字符串转对象
func jsonToObj() {
	jsonStr := "{\"name\":\"Harden\",\"age\":50,\"address\":\"Houston\"}"
	var d1 model.Doctor
	if err := json.Unmarshal([]byte(jsonStr), &d1); err != nil {
		fmt.Println(err)
		return
	}
	d1.PrintAllFields()

	p1 := model.Patients{Name: "病人1", Num: 100, Symptoms: []string{"头痛", "恶心", "发热"}, Doctor: d1}
	bytes, err := json.Marshal(p1)
	if err != nil {
		fmt.Println(err)
		return
	}
	jsonStr = string(bytes)
	var p2 model.Patients
	err = json.Unmarshal(bytes, &p2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(p2.Doctor.Name)
}

// json字符串转map
func jsonToMap() {
	d1 := model.Doctor{Name: "医生1", Age: 50, Address: "上海"}
	p1 := model.Patients{Name: "病人1", Num: 100, Symptoms: []string{"头痛", "恶心", "发热"}, Doctor: d1}
	bytes, err := json.Marshal(p1)
	if err != nil {
		fmt.Println(err)
		return
	}
	jsonStr := string(bytes)
	fmt.Println(jsonStr)

	var m map[string]interface{}
	if err := json.Unmarshal(bytes, &m); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(m)

	symptoms, ok := m["symptoms"].([]interface{})
	if ok {
		fmt.Println(symptoms[0])
	}
	doctorMap, ok := m["doctor"].(map[string]interface{})
	if ok {
		fmt.Println(doctorMap["name"])
	}
}

// json字符串转切片
func jsonToSlice() {
	arr1 := []string{"str1", "str2", "str3", "str4"}
	bytes, err := json.Marshal(arr1)
	if err != nil {
		fmt.Println(err)
		return
	}
	jsonStr := string(bytes)
	fmt.Println(jsonStr)

	var arr2 []string
	if err := json.Unmarshal([]byte(jsonStr), &arr2); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(arr2[0])
}
