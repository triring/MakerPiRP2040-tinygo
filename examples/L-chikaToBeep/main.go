// > tinygo build -o L-chikaToBeep.uf2 -target=pico -size short .
//	 code    data     bss |   flash     ram
//	16640     180    3184 |   16820    3364
//
// > tinygo flash -target=pico -size short .

// L-chikaを改良
// LEDを1秒間隔で点滅させていたものを
// 周期を変えて、Buzzピンに出力する。

package main

import ( // 必要なパッケージの読み込み
	"fmt"
	"machine"
	"time"
)

func main() {
	fmt.Printf("L-chika\n")
	// led := machine.GP9 // LED が接続されているピン番号
	led := machine.GP22 // ブサー が接続されているピン番号
	// LEDが接続されているポートを指定し、出力モードに設定する。
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for { // 無限ループ
		led.High()                       // LED 点灯
		time.Sleep(time.Millisecond * 5) // 0.005秒 待つ
		led.Low()                        // LED 消灯
		time.Sleep(time.Millisecond * 5) // 0.005秒 待つ
	}
}

// オンオフのタイミングを速くすると音になる。
