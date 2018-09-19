package main

import (
	"flag"
	"fmt"
	"io/ioutil"
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
	delay = flag.Int("delay", 2, "how long to delay single color")
)

func main() {

	flag.Parse()

	red := LED{"RED", 0, 1}

	green := LED{"GREEN", 0, 2}

	blue := LED{"BLUE", 0, 4}

	defer red.off()

	defer green.off()

	defer blue.off()

	switch strings.ToLower(*color) {

	case "":

		play(red, green, blue)

	case "red":

		red.on()

		time.Sleep(time.Duration(*delay) * time.Second)

	case "green":

		green.on()

		time.Sleep(time.Duration(*delay) * time.Second)

	case "blue":

		blue.on()
		fmt.Printf("blue.on\n")
		time.Sleep(time.Duration(*delay) * time.Second)

	}

}
