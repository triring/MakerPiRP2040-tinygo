// > tinygo build -o PWM_MotorCTL.uf2 -target=pico -size short .
//   code    data     bss |   flash     ram
//  10120     108    5240 |   10228    5348
//
// > tinygo flash -target=pico -size short -monitor .

package main

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/buzzer"
)

type note struct {
	tone     float64
	duration float64
}

func main() {
	Motor1A := machine.GP8 // M1A : GP8 モータ制御ポートA
	Motor1A.Configure(machine.PinConfig{Mode: machine.PinOutput})
	bzr := buzzer.New(Motor1A)

	Motor1B := machine.GP9 // M1A : GP9 モータ制御ポートB
	Motor1B.Configure(machine.PinConfig{Mode: machine.PinOutput})
	Motor1B.Low()

	for i := 0; i < 6; i++ {
		bzr.Tone(buzzer.G5, buzzer.Half)
		time.Sleep(10 * time.Millisecond)
		bzr.Tone(buzzer.A5, buzzer.Half)
		time.Sleep(10 * time.Millisecond)
		bzr.Tone(buzzer.F5, buzzer.Half)
		time.Sleep(10 * time.Millisecond)
		bzr.Tone(buzzer.F4, buzzer.Half)
		time.Sleep(10 * time.Millisecond)
		bzr.Tone(buzzer.C5, buzzer.Whole)
		time.Sleep(10 * time.Millisecond)
		time.Sleep(1000 * time.Millisecond)
	}
	Motor1A.Low()
	Motor1B.Low()
}
