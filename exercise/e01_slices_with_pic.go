package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	// Tạo slice 2D: có dy hàng
	picture := make([][]uint8, dy)

	// Với mỗi hàng y
	for y := 0; y < dy; y++ {
		// Tạo một slice cho hàng y, có dx cột
		picture[y] = make([]uint8, dx)

		// Gán giá trị cho từng cột x
		for x := 0; x < dx; x++ {
			// Chọn một trong các hàm sau:
			// 1. (x + y) / 2
			// 2. x * y
			// 3. x ^ y  (xor)

			// Ví dụ dùng: x ^ y
			value := x ^ y
			picture[y][x] = uint8(value)
		}
	}

	return picture
}

func main() {
	pic.Show(Pic)
}
