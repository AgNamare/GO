package main

import "golang.org/x/tour/pic"


func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	
	for i:=0; i<dy; i++ {
		pic[i] = make([]uint8, dx)
		for t := 0; t<dx; t++{
			pic[i][t] = uint8((i+t)/2)
		}
	}
	return pic
}

func main() {
	pic.Show(Pic)
}
