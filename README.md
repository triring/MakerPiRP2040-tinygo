# MakerPiRP2040-tinygo
## TinyGo で Cytron Maker Pi RP2040 を動かす。  
ある勉強会でマイコンの入出力について解説をすることになったので、その説明資料として作成した簡単なサンプル集です。   
開発環境は、[TinyGo](https://tinygo.org/) を用い、[Cytron社](https://www.cytron.io/) の[Maker Pi RP2040](https://www.cytron.io/p-maker-pi-rp2040-simplifying-robotics-with-raspberry-pi-rp2040)
 をターゲットとして作成しました。　

特別な入出力機能は使用していないので、入出力のピン番号をお使いのボードに合わせていただければ、他のボードでも利用できると思います。  
広く公開されているサンプルに、解説し易いように少し手を加えただけなので、野良コードとし、特にライセンスは設定いたしません。ご自由にお使い下さい。  

## Hardware:

[Maker Pi RP2040](https://www.cytron.io/p-maker-pi-rp2040-simplifying-robotics-with-raspberry-pi-rp2040)

このボードは、Raspberry Pi 財団が開発したマイクロコントローラー(マイコン)　RP2040を使用したロボット制御組込ボードで、その仕様は以下のとおりです。  

* CPU：RP2040,ROM：2MB,RAM：264KB
* GPIO：18
* DCモータドライバ	2個
* サーボモータ端子	4個
* Groveコネクタ	7個
* テスト用スイッチ	2個
* 確認用LED
    * RGB LED （Neopixel）	2個
    * インジケータに使える直列に配置されたLED	13個
* ブザーが1個(On/Offスイッチ付き)

## Software:

[TinyGo](https://tinygo.org/)

### 開発環境のインストール  

ここでは、Windows11上での開発環境構築について解説します。他のOSについては、本家サイトの解説をお読み下さい。

1. パッケージ管理ツールscoopのサイトを開き、導入スクリプトを入手する。

	[scoop](https://github.com/ScoopInstaller/Scoop)

2. Powershellを開いて、以下のスクリプトを実行する。

> \> Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser  
> \> Invoke-RestMethod -Uri https://get.scoop.sh | Invoke-Expression  


3. 以下のコマンドを実行して、環境構築は終わり。

> \>scoop install go tinygo

4. 以下のコマンドを実行できればOK

> \>tinygo version  
tinygo version 0.38.0 windows/amd64 
(using go version go1.24.4 and LLVM version 19.1.2)

## サンプルリスト:

```
|   go.mod
|   go.sum
|   README.md
|
\---examples
    +---AnalogInput_light
    |       main.go
    +---AnalogInput_light_waveform
    |       main.go
    +---hand-cranked_Motor
    |       main.go
    +---InterruptTest
    |       main.go
    +---L-chika
    |       main.go
    +---L-chikaToBeep
    |       main.go
    +---PWM_MotorCTL
    |       main.go
    +---PWM_Motor_Music
    |       main.go
    +---SwitchToConsole
    |       main.go
    +---SwitchToLED
    |       main.go
    \---ToneTest
            main.go
```

### [L-chika](examples/L-chika/main.go)
	LEDを秒間単位の間隔で点滅させる。

### [L-chikaToBeep](examples/L-chikaToBeep/main.go)
    L-chikaを改良
	出力先をLEDから、ブザーに変更し、OnOffの間隔を短くしたもの

### [ToneTest](examples/ToneTest/main.go)
	toneパッケージを使い、設定した2種類の周波数の音を交互に鳴らす。
	出力先をLEDから、ブザーに変更し、OnOffの間隔を短くしたもの

### [hand-cranked_Motor](examples/hand-cranked_Motor/main.go)
    ボタンが押されたら、モータを駆動する。手動PWM制御

### [PWM_MotorCTL](examples/PWM_MotorCTL/main.go)
	3段階にパルス幅を変更してモータを回す

### [PWM_Motor_Music](examples/PWM_Motor_Music/main.go)
    PWMの設定を音階と同じ周波数にしてモータに送ると音楽が聞こえる。
    "tinygo.org/x/drivers/buzzer"　パッケージを使用

### [SwitchToLED](examples/SwitchToLED/main.go)
	スイッチが押されたら、LEDを点灯する。

### [SwitchToConsole](examples/SwitchToConsole/main.go)
	スイッチが押されたら、LEDを点灯する。また、その状態をコンソールに表示する。

### [AnalogInput_light](examples/AnalogInput_light/main.go)
	アナログ信号の読み込み。
	LEDを接続し、光起電力効果による起電力を測定する。

### [AnalogInput_light_waveform](examples/AnalogInput_light_waveform/main.go)
	アナログ信号の読み込み。
	LEDを接続し、光起電力効果による起電力をコンソールに表示する。

### [InterruptTest](examples/InterruptTest/main.go)
	割り込みプログラムのサンプル。
    スイッチが押されたら、割り込みでLEDを点灯する。  
