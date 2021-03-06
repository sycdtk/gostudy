package main

import (
	"fmt"
	"reflect"
)

type Mystruct struct {
	Name string
}

type MyStruct struct {
	name string `fname:"1111";lname:"2222"`
	Age  int
}

func (m MyStruct) GetName() string {
	return "getname:" + m.name
}

func (m *MyStruct) GetName1() string {
	return "getname1:" + m.name
}

func (m *MyStruct) GetName2(str string) string {
	return "getname2:" + m.name + str
}

func main() {

	//	fmt.Println("对象的引用操作--------------------------------------")
	//	m1 := new(MyStruct) //对象的引用  &MyStruct{}，需要使用Elem()
	//	m1.name = "m1"
	//	m1.Age = 19

	//	m2 := MyStruct{} //对象
	//	m2.name = "m2"
	//	m2.Age = 9

	//	m1t := reflect.TypeOf(m1)  //涉及字段定义、方法名称获取
	//	m1v := reflect.ValueOf(m1) //涉及对象方法调用、数值修改

	//	fmt.Println(m1t.Kind() == reflect.Ptr) //判断是对象还是对象的引用reflect.Struct

	//	fmt.Println(m1t.Elem().Field(0).Tag.Get("fname")) //读取tag值

	//	nf := m1t.Elem().NumField() //获取字段数量
	//	fmt.Println(nf)
	//	for i := 0; i < nf; i++ {
	//		fmt.Println(m1t.Elem().Field(i).Name) //输出字段名称
	//	}

	//	nm := m1t.Elem().NumMethod() //获取方法数量，关联对象
	//	fmt.Println(nm)
	//	for i := 0; i < nm; i++ {
	//		fmt.Println(m1t.Elem().Method(i).Name) //输出方法名称
	//	}

	//	nm = m1t.NumMethod() //获取方法数量，关联对象引用
	//	fmt.Println(nm)
	//	for i := 0; i < nm; i++ {
	//		fmt.Println(m1t.Method(i).Name) //输出方法名称
	//	}

	//	fmt.Println(m1v.Elem().FieldByName("name")) //获取字段值
	//	fmt.Println(m1v.Elem().FieldByName("Age"))

	//	fmt.Println(m1v.Elem().FieldByName("name").CanSet()) //能够修改值
	//	fmt.Println(m1v.Elem().FieldByName("Age").CanSet())

	//	m1v.Elem().FieldByName("Age").SetInt(99) //修改值
	//	fmt.Println(m1v.Elem().FieldByName("Age"))

	//	m1v.Elem().FieldByName("Age").Set(reflect.ValueOf(66)) //修改值
	//	fmt.Println(m1v.Elem().FieldByName("Age"))

	//	fmt.Println(m1v.MethodByName("GetName").Call(nil))                                      //调用对象关联的函数
	//	fmt.Println(m1v.MethodByName("GetName1").Call(nil))                                     //调用对象引用关联的函数
	//	fmt.Println(m1v.MethodByName("GetName2").Call([]reflect.Value{reflect.ValueOf("你好！")})) //调用带参数的函数

	//	r := m1v.MethodByName("GetName2").Call([]reflect.Value{reflect.ValueOf("你好！")})
	//	fmt.Println(r[0].String()) //处理返回值

	//	fmt.Println("对象操作--------------------------------------")

	//	m2t := reflect.TypeOf(m2)  //涉及字段定义、方法名称获取
	//	m2v := reflect.ValueOf(m2) //涉及对象方法调用、数值修改

	//	fmt.Println(m2t.Kind() == reflect.Struct) //判断是对象还是对象的引用reflect.Ptr

	//	fmt.Println(m2t.Field(0).Tag.Get("fname")) //读取tag值

	//	nf = m2t.NumField() //获取字段名称
	//	fmt.Println(nf)
	//	for i := 0; i < nf; i++ {
	//		fmt.Println(m2t.Field(i).Name) //输出字段名称
	//	}

	//	fmt.Println(m2v.FieldByName("name")) //获取字段值
	//	fmt.Println(m2v.FieldByName("Age"))

	//	fmt.Println(m2v.FieldByName("name").CanSet()) //能够修改值
	//	fmt.Println(m2v.FieldByName("Age").CanSet())

	//	fmt.Println(m2v.MethodByName("GetName").Call(nil)) //调用对象关联的函数
	//	//	fmt.Println(m2v.MethodByName("GetName1").Call(nil))                                     //调用对象引用关联的函数  对象调用不成功
	//	//	fmt.Println(m2v.MethodByName("GetName2").Call([]reflect.Value{reflect.ValueOf("你好！")})) //调用带参数的函数 对象调用不成功

	//	nm = m2t.NumMethod() //获取方法数量，关联对象，无法调用关联对象的引用
	//	fmt.Println(nm)
	//	for i := 0; i < nm; i++ {
	//		fmt.Println(m2t.Method(i).Name) //输出方法名称
	//	}

	fmt.Println("构建信对象--------------------------------------")

	m11 := &Mystruct{Name: "aa"}
	fmt.Println(m11)

	m22 := Mystruct{Name: "bb"}
	fmt.Println(m22)

	m33 := MyNew(m11, "Name", "11")
	fmt.Println(m33)

	m44 := MyNew(m22, "Name", "22")
	fmt.Println(m44)

	m55 := MyNew2(m44, set)

	fmt.Println(m55)

}

func MyNew(sample interface{}, field string, value string) interface{} {
	t := reflect.TypeOf(sample)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	fmt.Println(t)
	v := reflect.New(t).Elem().Interface()

	if rv, ok := v.(Mystruct); ok {
		fmt.Println("-->", reflect.TypeOf(rv))
		reflect.Indirect(reflect.ValueOf(&rv)).FieldByName(field).SetString(value)
		return &rv
	}

	//	reflect.ValueOf(v).Elem().FieldByName(field).SetString(value)

	return v
}

func MyNew2(stru interface{}, f func(o interface{}, f, v string)) interface{} {
	t := reflect.TypeOf(stru)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	//	if rv, ok := v.(Mystruct); ok {
	//		fmt.Println("-->", reflect.TypeOf(rv))
	//		reflect.Indirect(reflect.ValueOf(&rv)).FieldByName("Name").SetString("xx")
	//		f(&rv, "Name", "qqqq")
	//		return &rv
	//	}
	f(&v, "Name", "qqqq")

	return v
}

func set(v reflect.Type, field, value string) {
	v := reflect.New(t).Elem().Interface()
	v1, ok := v.(*Mystruct)
	fmt.Println("00-->", reflect.TypeOf(v1))
	fmt.Println(v1, "=", ok)
	if ok {
		fmt.Println("===-->", reflect.TypeOf(v1))
		reflect.Indirect(reflect.ValueOf(&v1).Elem()).FieldByName(field).SetString(value)
	}
}
