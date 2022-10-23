package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"unicode/utf8"
)

func main() {
	fmt.Println("hello world")

	var a, b, c, d = 3, 4, true, "abc"
	fmt.Println(a, b, c, d)
	fmt.Println(1<<10, 1<<20, 1<<30, 1<<40)

	fmt.Println(strconv.Itoa(a))

	switch a {
	case 3:
		fmt.Println("print equal 3")
	case 4:
		fmt.Println("print equal 4")
	default:
		fmt.Println("any else")
	}
	switch {
	case a > 4:
		fmt.Println("greater than 3")
	case a > 2:
		fmt.Println("greater than 4")
	}

	file, err := os.ReadFile("test1.txt")
	if err != nil {
		panic("file not exist")
	}

	fmt.Println(file)
	fmt.Println(string(file))

	file2, err := os.Open("test2.txt")
	scanner := bufio.NewScanner(file2)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	//for {
	//	fmt.Println("forever")
	//}
	apply(func(i int, i2 int) int {
		return i + i2
	}, 1, 2)
	fmt.Println(pow(2, 3))
	fmt.Println(sum(1, 2, 3))
	pointer()
	a, b = 3, 4
	swap(&a, &b)
	fmt.Printf("%d, %d\n", a, b)
	m, n := 3, 4
	m, n = swap2(m, n)
	fmt.Printf("%d, %d\n", m, n)
	arrays()
	slice()
	printMap()
	printRune()
	fmt.Println(lengthOfNonRepeatingSubStr("abcabcbb")) //3
	fmt.Println(lengthOfNonRepeatingSubStr("bbbbbbb"))  //1
	fmt.Println(lengthOfNonRepeatingSubStr("abdevbac")) //6
	printMultipleString()
}

func apply(op func(int, int) int, a, b int) int {
	pointer := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(pointer).Name()
	fmt.Printf("Calling function %s with args (%d, %d)\n", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func pointer() {
	var a int = 1
	var i *int = &a
	*i = 2
	fmt.Println(a)
}

func swap(a, b *int) {
	*b, *a = *a, *b
}
func swap2(a, b int) (int, int) {
	return b, a
}
func arrays() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 7, 10}
	var grid [4][5]int

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}
	for idx := range arr3 {
		fmt.Println(arr3[idx])
	}
	for i, val := range arr3 {
		fmt.Printf("%d, %d\n", i, val)
	}
}

// view of array, 引用传递
func slice() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr = ", arr)
	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Println("arr[:6] = ", arr[:6])
	fmt.Println("arr[2:] = ", arr[2:])
	fmt.Println("arr[:] = ", arr[:])

	var s []int
	for i := 0; i < 10; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}

	s1 := []int{2, 4, 6, 8}
	printSlice(s1)

	s2 := make([]int, 16)     //slice len
	s3 := make([]int, 10, 32) //slice len, cap
	printSlice(s2)
	printSlice(s3)

	fmt.Println("Copying slice")
	copy(s2, s1)
	printSlice(s2)
	fmt.Println("Deleting elements from slice")
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)
}

func printSlice(s []int) {
	fmt.Printf("%v, len=%d, cap=%d\n", s, len(s), cap(s))
}

func printMap() {
	m := map[string]string{
		"name":    "ds",
		"age":     "18",
		"country": "china",
		"class":   "aaa",
	}
	m2 := make(map[string]int) // m2 == empty map
	var m3 map[string]int      // m3==nil
	fmt.Println(m, m2, m3)
	fmt.Printf("map=%v\n", m)

	for k := range m {
		fmt.Println(k)
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
	clz, ok := m["class"]
	clz1, ok1 := m["class1"]
	fmt.Println(clz, ok)
	fmt.Println(clz1, ok1)

	delete(m, "class")
	fmt.Println(m)
}

func printRune() {
	s := "Yes我爱中国!"
	fmt.Println(s)
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()
	for i, ch := range s {
		fmt.Printf("(%d %X) ", i, ch)
	}
	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c %X) ", i, ch, ch)
	}
	fmt.Println()
}

// abdevbac
func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0

	for i, ch := range []byte(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func printMultipleString() {
	s := `
		hello "world"
		hello "world1"
	`
	fmt.Println(s)
}

func calcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}
