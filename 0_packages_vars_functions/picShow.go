package main
import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	
	var pic [][] uint8
	for i := 0; i < dy; i++ {
		var slice [] uint8
		for j := 0; j < dx; j++ {
			//slice = append(slice, uint8((i+j)/2))
			slice = append(slice, uint8(i^j))
		}
		pic = append(pic, slice)
	}
	return pic

}

func main() {
	pic.Show(Pic)
}
