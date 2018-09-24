//reference https://golang.org/pkg/flag/

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
	"time"
)

type LightCombo struct {
	red bool

	green bool

	blue bool
}

type LED struct {
	name string

	brightness int

	index uint8
}

func (l LED) write() {

	fn := fmt.Sprintf("/sys/devices/soc0/leds/leds/%s/brightness", l.name)

	vl := fmt.Sprintf("%d", l.brightness)

	ioutil.WriteFile(fn, []byte(vl), 0644)

}

func (l LED) off() {

	l.brightness = 0

	l.write()

}

func (l LED) on() {

	l.brightness = 1

	l.write()

}

func (l LED) set(onoff bool) {

	if onoff {

		l.on()

	} else {

		l.off()

	}

}

func play(red, green, blue LED) {

	fmt.Printf("play func\n")

	colors := []LightCombo{

		LightCombo{red: true, green: true, blue: true}, // white

		// Make it easy on the testers' eyes

		LightCombo{red: true},

		LightCombo{green: true, blue: true}, // turquoise

		//

		LightCombo{green: true},

		LightCombo{red: true, blue: true}, // violet

		//

		LightCombo{blue: true},

		LightCombo{red: true, green: true}, // yellow

	}

	for _, combo := range colors {

		red.set(combo.red)

		green.set(combo.green)

		blue.set(combo.blue)

		time.Sleep(500 * time.Millisecond)
	}
}

var (
	color = flag.String("color", "", "color you want to turn on")
	delay = flag.Int("delay", 1000, "how long to delay single color")
)

func play_with_led(color_choose string, delay_time int) {

	*color = color_choose
	*delay = delay_time

	flag.Parse()

	red := LED{"RED", 0, 1}

	green := LED{"GREEN", 0, 2}

	blue := LED{"BLUE", 0, 4}

	defer red.off()

	defer green.off()

	defer blue.off()

	fmt.Printf("case color=%s\n", *color)

	switch strings.ToLower(*color) {
	case "":

		fmt.Printf("case empty on\n")

		play(red, green, blue)

	case "red":

		red.on()
		fmt.Printf("case red on\n")
		time.Sleep(time.Duration(*delay) * time.Millisecond)

	case "green":

		green.on()
		fmt.Printf("case green on\n")
		time.Sleep(time.Duration(*delay) * time.Millisecond)

	case "blue":

		blue.on()
		fmt.Printf("case blue on\n")
		time.Sleep(time.Duration(*delay) * time.Millisecond)

	case "white":

		red.on()
		green.on()
		blue.on()
		fmt.Printf("case white on\n")
		time.Sleep(time.Duration(*delay) * time.Millisecond)

	case "turquoise":

		green.on()
		blue.on()
		fmt.Printf("case turquoise on\n")
		time.Sleep(time.Duration(*delay) * time.Millisecond)

	case "violet":

		red.on()
		blue.on()
		fmt.Printf("case violet on\n")
		time.Sleep(time.Duration(*delay) * time.Millisecond)

	case "yellow":

		red.on()
		green.on()
		fmt.Printf("case yellow on\n")
		time.Sleep(time.Duration(*delay) * time.Millisecond)
	}

}

func dim_controling(color_choose string, dim_time int, in_out bool) {

	flag.Parse()

	red := LED{"RED", 0, 1}

	green := LED{"GREEN", 0, 2}

	blue := LED{"BLUE", 0, 4}
	fmt.Printf("dim_controling color=%s\n", *color)

	var loop_times int
	var a_delay int
	var res_delay int
	var delta_time int

	delta_time = 10
	loop_times = dim_time / delta_time

	fmt.Printf("loop= %d\n", loop_times)

	for a := 0; a < loop_times; a++ {
		a_delay = delta_time * a / loop_times
		res_delay = delta_time * (loop_times - a) / loop_times

		switch strings.ToLower(*color) {

		case "":
			fmt.Printf(" on\n")

			//play(red, green, blue)

		case "red":
			red.on()

		case "green":
			green.on()

		case "blue":
			blue.on()

		case "white":
			red.on()
			green.on()
			blue.on()

		case "turquoise":
			green.on()
			blue.on()

		case "violet":
			red.on()
			blue.on()

		case "yellow":
			red.on()
			green.on()
		}

		if in_out {

			time.Sleep(time.Duration(a_delay) * time.Millisecond)

		} else {

			time.Sleep(time.Duration(res_delay) * time.Millisecond)
			red.off()
			green.off()
			blue.off()

		}

		if in_out {

			time.Sleep(time.Duration(res_delay) * time.Millisecond)
			red.off()
			green.off()
			blue.off()

		} else {

			time.Sleep(time.Duration(a_delay) * time.Millisecond)

		}

	}

}

func dim_in_out_with_led(color_choose string, dim_time int, in_out bool) {

	flag.Parse()
	*color = color_choose

	red := LED{"RED", 0, 1}

	green := LED{"GREEN", 0, 2}

	blue := LED{"BLUE", 0, 4}

	defer red.off()

	defer green.off()

	defer blue.off()

	fmt.Printf("dim color=%s\n", *color)

	dim_controling(color_choose, dim_time, in_out)

}

func exeSysCommand(cmdStr string) string {
	cmd := exec.Command("sh", "-c", cmdStr)
	opBytes, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(opBytes)
}

func GetLocaltime() string {
	tmp := exeSysCommand("date")
	if len(tmp) == 0 {
		fmt.Println("Get local time Failed")
		return ""
	}

	localtime := strings.Trim(tmp, "\n")
	return localtime
}

func time_checking() {
	//t := time.Now()
	//fmt.Printf("%s\n", t)
	fmt.Printf("%s\n", GetLocaltime())

}

func flash_serval_times(color_choose string, delay_time int, times int) {

	for a := 0; a < times; a++ {
		play_with_led(color_choose, delay_time)
		time.Sleep(time.Duration(delay_time) * time.Millisecond)
	}

}

func main() {

	var duration_flash int
	var duration_flash_times int
	var duration_flow int
	var duration_dim int

	duration_flash = 100
	duration_flash_times = 10

	/**/
	duration_flow = 200

	duration_dim = 2000

	fmt.Printf("Flow colors\n\n\n")
	//light on test, for seconds
	play_with_led("red", duration_flow)
	play_with_led("green", duration_flow)
	play_with_led("blue", duration_flow)
	play_with_led("white", duration_flow)
	play_with_led("turquoise", duration_flow)
	play_with_led("violet", duration_flow)
	play_with_led("yellow", duration_flow)
	/**/

	fmt.Printf("Flash colors\n\n\n")
	flash_serval_times("red", duration_flash, duration_flash_times)
	flash_serval_times("green", duration_flash, duration_flash_times)
	flash_serval_times("blue", duration_flash, duration_flash_times)
	/*
		fmt.Printf("Flash yellow 200s 5 times\n\n\n")
		time.Sleep(time.Duration(1000) * time.Millisecond)
		duration_flash = 200
		flash_serval_times("yellow", duration_flash, duration_flash_times)

		fmt.Printf("Flash yellow 200s 5 times\n\n\n")
		time.Sleep(time.Duration(1000) * time.Millisecond)
		duration_flash = 400
		flash_serval_times("yellow", duration_flash, duration_flash_times)

		fmt.Printf("Flash yellow 200s 5 times\n\n\n")
		time.Sleep(time.Duration(1000) * time.Millisecond)
		duration_flash = 600
		flash_serval_times("yellow", duration_flash, duration_flash_times)
	*/

	fmt.Printf("DIM controlling\n\n\n")
	//dim in test for ms, needs to
	time_checking()
	dim_in_out_with_led("red", duration_dim, true)
	time_checking()
	dim_in_out_with_led("green", duration_dim, true)
	time_checking()
	dim_in_out_with_led("blue", duration_dim, true)
	time_checking()
	dim_in_out_with_led("white", duration_dim, true)
	time_checking()
	dim_in_out_with_led("turquoise", duration_dim, false)
	time_checking()
	dim_in_out_with_led("violet", duration_dim, false)
	time_checking()
	dim_in_out_with_led("yellow", duration_dim, false)
	time_checking()

}
