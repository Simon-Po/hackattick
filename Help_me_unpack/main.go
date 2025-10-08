package main

import (
	"fmt"

	"hackattic/Help-me-unpack/hmu"
)

func main() {
	b := hmu.FetchAndParse("https://hackattic.com/challenges/help_me_unpack/problem?access_token=be8466c39c975877")
	i, b := hmu.GetInt(b)
	ui, b := hmu.GetUInt(b)
	short, b := hmu.GetShort(b)
	b = hmu.Skip(b)
	float, b := hmu.GetFloat(b)
	double, b := hmu.GetDouble(b)
	doubleBe, _ := hmu.GetDoubleBE(b)

	fmt.Println("Int32: ", fmt.Sprint(i))
	fmt.Println("UInt32: ", fmt.Sprint(ui))
	fmt.Println("Short: ", fmt.Sprint(short))
	fmt.Println("float: ", fmt.Sprint(float))
	fmt.Println("double: ", fmt.Sprint(double))
	fmt.Println("Double be: ", fmt.Sprint(doubleBe))

	hmu.SendResult(hmu.Response{I: i, UI: ui, Short: short, Float: float, Double: double, DoubleBe: doubleBe})
}
