package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Giangvien struct {
	Name   string `json:"hoten"`
	Gender int    `json:"gioitinh"`
	Phone  string `json:"dienthoai"`
}

func (gv *Giangvien) hienthithongtin() {
	fmt.Printf("Ho ten: %s \n", gv.Name)
	fmt.Printf("Gioi tinh: %d \n", gv.Gender)
	fmt.Printf("Dien thoai: %s \n", gv.Phone)
}

func (gv *Giangvien) clear() {
	gv.Name = ""
	gv.Gender = 0
	gv.Phone = ""
}

func main() {

	minhbao := Giangvien{
		Name:   "Vo Minh Bao",
		Gender: 1,
		Phone:  "0901234567",
	}

	// minhbao.hienthithongtin()

	// fmt.Println()

	// aravo := Giangvien{
	// 	Name:   "ara Vo",
	// 	Gender: 1,
	// 	Phone:  "0907654321",
	// }

	// aravo.hienthithongtin()

	// fmt.Println()

	// johny := Giangvien{
	// 	"Johny",
	// 	1,
	// 	"0903456789",
	// }

	// johny.hienthithongtin()

	// johny.clear()

	// fmt.Println()

	// johny.hienthithongtin()

	output, err := json.Marshal(minhbao)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}
