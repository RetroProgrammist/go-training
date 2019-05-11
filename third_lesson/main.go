package main

import (
	"fmt"
	"runtime"
)

var c, python, java bool //Объявление переменной вне тела функции, все типа bool 
const k float32 = 273.15 // declaration constants

func main() {
	var pizza, pizza1 string = "Margnerita", "Dont Margnerita" //declaration + initialisation

	var pizza2 string //declaration
	pizza2 = "Margnerita" //initialization (or initialisation)

	pizza3 := "Margnerita"//type inference - declaration + initialisation
	
	fmt.Println(pizza, pizza1)
	fmt.Println(pizza2)
	fmt.Println(pizza3)
	//! В будущем сменить тип переменной можно только явным преобразованием типов (type conversion)
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)	
	
	fmt.Println(u)

	/*------------------Constants------------------*/
	/*
		Во время выполнения программы не может меняться ее значение 
	*/
	const e = 2.718 // указывать тип не обязательно
	const (
		g = 9.8
		Hello = "Hellow World" 
	)
	//iota auto implemented properties
	const (
		Fantesy = iota //0
		Social = iota //1
		Nature = iota //2
	)
	/*Или
	const (
		Fantesy = iota //0
		Social 	//1
		Nature  //2
	)*/


	someFunc(1, 2, "3")
	Hi, _ := notUseThis("Hello ", "World")
	fmt.Println(Hi)

	/*------------------Arrays && Slices------------------*/
	var ar [2]string
	ar[0] = "element 1"
	ar[1] = "element 2"
	//ar[3] = "element 3" //Лишнее 

	arInt := [6]int{1, 2, 3, 53, 4, 22}
	fmt.Println(arInt)

	var firstSlice []int = arInt[1:4] // Сделали из arInt 
	fmt.Println(firstSlice)

	//https://tour.golang.org/moretypes/1
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[0:5]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)


	/*------------------if statement initialization------------------*/
	lim := 10
	if v := 3; v < lim {
		fmt.Printf("%g >= %g\n", v, lim)
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}

	/*------------------Switch------------------*/
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

/*------------------Function Return && params/arguments function------------------*/
/*
	Parameter is variable in the declaration of function. 
	Argument is the actual value of this variable that gets passed to function.
*/
func someFunc(i, t int, s string) (int, string) {
	var someNum int = 22;
	fmt.Println(i , t , s) //1, 2, "3"
	return someNum, s
}

func notUseThis(a, b string) (c,d string) {
	c = a + b 
	d = b + c	
	return
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}