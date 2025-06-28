// > tinygo build -o L-chika.uf2 -target=pico -size short .
//
//	 code    data     bss |   flash     ram
//	16640     180    3184 |   16820    3364
//
// > tinygo flash -target=pico -size short .
package main

import ( // 必要なパッケージの読み込み
	"fmt"
	"machine"
	"time"
)

func main() {
	fmt.Printf("L-chika\n")
	led := machine.GP2 // LED が接続されているピン番号
	// LEDが接続されているポートを指定し、出力モードに設定する。
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for { // 無限ループ
		led.High()                  // LED 点灯
		time.Sleep(time.Second * 2) // 2秒 待つ
		led.Low()                   // LED 消灯
		time.Sleep(time.Second * 1) // 1秒 待つ
	}
}

// ゲンジボタルの発光パターン, 西日本と九州[点灯2秒,消灯1秒],
// 東日本[点灯4秒,消灯1秒], 長崎県五島列島[点灯1秒,消灯1秒]
