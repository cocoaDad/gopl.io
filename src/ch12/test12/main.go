package main

import (
	"fmt"
	"reflect"
	"os"
)

func main(){
	fmt.Fprintf(os.Stdout, "%T\n", 3) // output the type of "3" directly
	t := reflect.TypeOf(3)
	fmt.Fprintf(os.Stdout, "%s\n", t.String()) // output the type with the help of function named reflect.Typeof()
	fmt.Fprintf(os.Stdout, "%s\n", t) // same as the line above
	fmt.Println("----------------------------------------------------------------")
	v := reflect.ValueOf(3)
	fmt.Fprintf(os.Stdout, "%v\n", v) // output the reflect.Value
	fmt.Fprintf(os.Stdout, "%v\n", v.String()) // attention: this output is "<int Value>"
	t = v.Type() // output the type of the reflect.Value
	fmt.Fprintf(os.Stdout, "%s\n", t.String())

	x := v.Interface() // 将v转换为一个纯粹interface{}类型
	i := x.(int)  // 接口断言，若x作为一个接口满足int这个接口的各个方法，则i的类型为int，值为3
	fmt.Fprintf(os.Stdout, "%d\n", i)
	fmt.Println("----------------------------------------------------------------")
	v = reflect.ValueOf("String")
	fmt.Fprintf(os.Stdout, "%v\n", v) // same as the line 20 above
	fmt.Fprintf(os.Stdout, "%v\n", v.String()) // !+ only when the type of the v is string, function with Stringer can output String
	t = v.Type() // output the type of the reflect.Value
	fmt.Fprintf(os.Stdout, "%s\n", t.String())
}