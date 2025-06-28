// > tinygo build -o InterruptTest.uf2 -target=pico -size short .
//
//	 code    data     bss |   flash     ram
//	16640     180    3184 |   16820    3364
//
// > tinygo flash -target=pico -size short -monitor .
package main

import (
	"fmt"
	"machine"
	"time"
)

func main() {
	count := 0
	led := machine.GP2 // LED が接続されているピン番号
	// LEDが接続されているポートを指定し、出力モードに設定する。
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	sw20 := machine.GPIO20 // キー入力割込みの設定
	sw20.Configure(machine.PinConfig{Mode: machine.PinInput})
	sw20.SetInterrupt(machine.PinFalling, func(machine.Pin) {
		//	キーが押された時に呼び出される。
		if 0 != (count % 2) {
			led.High()
		} else { // 押される毎に、On/Offを切り替える
			led.Low()
		}
		count += 1 // 押された回数を記録する。
	})
	for {
		fmt.Printf("key count: %d\n", count)
		time.Sleep(time.Second * 4) // 4秒スリープ
	}
}
