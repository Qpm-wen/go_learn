package main

func main() {
	m := mkmap()
	println(m["a"])

	s := mkslice()
	println(s[0])

}

func mkslice() []int  {
	s := make ([]int, 0, 100)
	s = append(s, 100)
	return s
}

func mkmap() map[string]int  {
	m := make(map[string]int)
	m["a"] = 1
	return m
}
