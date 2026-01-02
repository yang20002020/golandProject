package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
)

type Person struct {
	Name string
	Age  int
	Sex  string
}

// 序列换
func writeJson(fileName string) (err error) {
	var persons []*Person
	for i := 0; i < 10; i++ {
		p := &Person{
			Name: fmt.Sprintf("name%d", i),
			Age:  rand.Intn(100),
			Sex:  fmt.Sprintf("Man"),
		}
		persons = append(persons, p)

	}
	data, err := json.Marshal(persons)
	if err != nil {
		fmt.Printf("marshal failed,err:%\v\n", err)
		return err
	}
	err = ioutil.WriteFile(fileName, data, 0755)
	if err != nil {
		fmt.Printf("write file failed,err:%v\n", err)
		return err
	}
	return

}

// 反序列化
func readJson(fileName string) (err error) {
	var persons []*Person
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &persons)
	if err != nil {
		return
	}
	fmt.Printf("%#v\n", persons)
	for _, v := range persons {
		fmt.Printf("%#v\n", v)
	}

	return
}

func main() {
	filename := "E:\\GO\\GOWORK\\src\\golandProject\\035json\\tmp\\person.txt"

	err := writeJson(filename)
	if err != nil {
		fmt.Printf("wrire json failed,err:%v\n", err)
		return
	}

	err = readJson(filename)
	if err != nil {
		fmt.Printf("read json failed,err:%v\n", err)
		return
	}

}
