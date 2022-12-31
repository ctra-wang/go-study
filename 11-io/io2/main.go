package main

import (
	"fmt"
)

/*
将err定义为string类型，然后把EOF设置为常量
*/
type err string
type code int

func (e err) AAA2() string {
	fmt.Println("____+++++___")

	return string(e)
}
func (e err) Err() string {
	fmt.Println("_______")
	return string(e)
}

const EOF = err("EOF")
const Code = code(1)

//const Err = err.Error()

func main() {
	//fmt.Println(EOF)
	//fmt.Println(Code)
	//
	//money := 0.12
	//test := "ABSIB T"
	//arr := strings.Split(test, " ")
	//if len(arr) == 1 {
	//	fmt.Println(len(arr[0]))
	//}
	//
	//for i := len(arr) - 1; i >= 0; i++ {
	//	if arr[i] != "" {
	//		fmt.Println(len(arr))
	//	}
	//}
	//reader := bufio.NewReader(os.Stdin)
	//data, _, _ := reader.ReadLine()
	//fmt.Println(string(data))

	//data1 := bufio.NewReader(os.Stdin)
	//data2 := bufio.NewReader(os.Stdin)
	//num1, _, _ := data1.ReadLine()
	//num2, _, _ := data2.ReadLine()
	//lower1 := strings.ToLower(string(num1))
	//lower2 := strings.ToLower(string(num2))
	//index := 0
	//for temp := range lower1 {
	//	if string(lower1[temp]) == lower2 {
	//		index++
	//	}
	//}
	//fmt.Print(index)

	//reader := bufio.NewReader(os.Stdin)
	//res, _, _ := reader.ReadLine()
	//
	//str := string(res)
	//
	//strSplit := strings.Split(str, "")
	//
	//if len(strSplit) < 8 {
	//	len := 8 - len(strSplit)
	//	temp := strings.Join(strSplit, "")
	//	for i := 0; i < len; i++ {
	//		temp = temp + "0"
	//	}
	//	fmt.Println(temp)
	//	return
	//}
	//
	//index := 0
	//for {
	//	if index+8 > len(strSplit) && len(strSplit[index:len(strSplit)]) != 0 {
	//		length := 8 - len(strSplit[index:len(strSplit)])
	//		temp1 := strings.Join(strSplit[index:len(strSplit)], "")
	//		for i := 0; i < length; i++ {
	//			temp1 = temp1 + "0"
	//		}
	//		fmt.Println(temp1)
	//		return
	//	}
	//	str := strSplit[index : index+8]
	//	temp := strings.Join(str, "")
	//	index += 8
	//	fmt.Println(temp)
	//}

	//for {
	//	var ret int
	//	_, err := fmt.Scanf("%p", &ret)
	//	if err != nil {
	//		break
	//	}
	//	fmt.Println(ret)
	//}
	//reader := bufio.NewReader(os.Stdin)
	//byte, _, _ := reader.ReadLine()
	//um := string(byte)
	//num, _ := strconv.Atoi(um)
	//if num < 0 {
	//	num += 4294967296
	//}
	//if num == 0 {
	//	fmt.Println("0")
	//}
	//res := []int32{}
	//hash := []int32{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}
	//for num != 0 {
	//	temp := num % 16
	//	num = num / 16
	//	fmt.Println("temp------------", temp)
	//	fmt.Println("num162------------", num)
	//	res = append([]int32{hash[temp]}, res...)
	//}
	var scan int64
	fmt.Scan(&scan)
	var i int64
	for i = 2; i*i <= scan; i++ {
		fmt.Println("----------", scan%i)
		for scan%i == 0 {
			fmt.Printf("%d \n", i)
			scan /= i
		}

	}

	//fmt.Println(str[0:1], str[1:2])
	//atoi, err := strconv.Atoi(str[1:2] )
	//if err != nil {
	//	return
	//}
	//res := strconv.FormatFloat(money, 'f', 2, 64)
	//fmt.Println(res)
}
