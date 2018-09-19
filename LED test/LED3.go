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

func (l LED) on_bright(LED_brightness int) {

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
	delay = flag.Int("delay", 2, "how long to delay single color")
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
		time.Sleep(time.Duration(*delay) * time.Second)

	case "green":

		green.on()
		fmt.Printf("case green on\n")
		time.Sleep(time.Duration(*delay) * time.Second)

	case "blue":

		blue.on()
		fmt.Printf("case blue on\n")
		time.Sleep(time.Duration(*delay) * time.Second)

	case "white":

		red.on()
		green.on()
		blue.on()
		fmt.Printf("case white on\n")
		time.Sleep(time.Duration(*delay) * time.Second)

	case "turquoise":

		green.on()
		blue.on()
		fmt.Printf("case turquoise on\n")
		time.Sleep(time.Duration(*delay) * time.Second)

	case "violet":

		red.on()
		blue.on()
		fmt.Printf("case violet on\n")
		time.Sleep(time.Duration(*delay) * time.Second)

	case "yellow":

		red.on()
		green.on()
		fmt.Printf("case yellow on\n")
		time.Sleep(time.Duration(*delay) * time.Second)
	}

}

func dim_controling(color_choose string, dim_time int, in_out bool) {

	flag.Parse()

	red := LED{"RED", 0, 1}

	green := LED{"GREEN", 0, 2}

	blue := LED{"BLUE", 0, 4}

	//defer red.off()

	//defer green.off()

	//defer blue.off()
	fmt.Printf("dim_controling color=%s\n", *color)

	var loop_times int
	var a_delay int
	var res_delay int
	var delta_time int
	var LED_brightness int
	//var a int

	delta_time = 30
	loop_times = dim_time / delta_time

	fmt.Printf("loop= %d\n", loop_times)

	for a := 0; a < loop_times; a++ {
		a_delay = 30
		res_delay = 30
		LED_brightness = (255 * a / loop_times)

		//fmt.Printf("a= %d   a_delay =%d    res_delay =%d \n", a, a_delay, res_delay)

		switch strings.ToLower(*color) {

		case "":
			fmt.Printf(" on\n")

			play(red, green, blue)

		case "red":
			//red.on()
			red.on_bright(LED_brightness)

		case "green":
			//green.on()
			green.on_bright(LED_brightness)

		case "blue":
			//blue.on()
			blue.on_bright(LED_brightness)

		case "white":
			red.on_bright(LED_brightness)
			green.on_bright(LED_brightness)
			blue.on_bright(LED_brightness)

		case "turquoise":
			green.on_bright(LED_brightness)
			blue.on_bright(LED_brightness)

		case "violet":
			red.on_bright(LED_brightness)
			blue.on_bright(LED_brightness)

		case "yellow":
			red.on_bright(LED_brightness)
			green.on_bright(LED_brightness)
		}

		if in_out {

			time.Sleep(time.Duration(a_delay) * time.Millisecond)

		} else {

			time.Sleep(time.Duration(res_delay) * time.Millisecond)

		}

		red.off()
		green.off()
		blue.off()

		if in_out {

			time.Sleep(time.Duration(res_delay) * time.Millisecond)

		} else {

			time.Sleep(time.Duration(a_delay) * time.Millisecond)

		}

	}

}

func dim_in_out_with_led(color_choose string, dim_time int, in_out bool) {

	flag.Parse()
	*color = color_choose
	/*

			red := LED{"RED", 0, 1}

			green := LED{"GREEN", 0, 2}

			blue := LED{"BLUE", 0, 4}

			defer red.off()
		ls

			defer green.off()

			defer blue.off()
	*/

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

func main() {
	/*
		fmt.Printf("Flow colors\n")
		//light on test, for seconds
		play_with_led("red", 1)
		play_with_led("green", 1)
		play_with_led("blue", 1)
		play_with_led("white", 1)
		play_with_led("turquoise", 1)
		play_with_led("violet", 1)
		play_with_led("yellow", 1)
	*/

	fmt.Printf("DIM controlling\n\n\n\n")
	//dim in test for ms, needs to
	time_checking()
	dim_in_out_with_led("red", 5000, true)
	time_checking()
	dim_in_out_with_led("green", 5000, true)
	time_checking()
	dim_in_out_with_led("blue", 5000, true)
	time_checking()
	dim_in_out_with_led("white", 5000, true)
	time_checking()
	dim_in_out_with_led("turquoise", 5000, false)
	time_checking()
	dim_in_out_with_led("violet", 5000, false)
	time_checking()
	dim_in_out_with_led("yellow", 5000, false)
	time_checking()

}
