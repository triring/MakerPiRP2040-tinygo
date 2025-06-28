// > tinygo build -o ToneTest.uf2 -target=pico -size short .
//    code    data     bss |   flash     ram
//   16640     180    3184 |   16820    3364
// > tinygo flash -target=pico -size short .

package main

import ( // パッケージを読込む。
	"machine"
	"time"

	"tinygo.org/x/drivers/tone"
)

var ( // 出力ポート等の定義
	// Configuration for the Adafruit Circuit Playground Bluefruit.
	pwm = machine.PWM3 // PWM ブロック3
	pin = machine.GP22 // ブザー接続ポート
)

func main() {
	// ドライバの初期化
	speaker, err := tone.New(pwm, pin)
	if err != nil {
		println("failed to configure PWM")
		return
	}
	for { // Two tone siren.
		speaker.SetNote(tone.B5)    // 鳴らす
		time.Sleep(time.Second / 2) // 待つ
		speaker.SetNote(tone.A5)    // 鳴らす
		time.Sleep(time.Second / 2) // 待つ
	}
}
