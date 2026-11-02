# TinyGo PIO Sample

TinyGo で Raspberry Pi Pico (RP2040) の PIO を使うサンプル集です。
「TinyGo Keebook vol.4」の補足として、本文に関連するサンプルコードを置いています。

## サンプル

| ディレクトリ | 内容 |
| --- | --- |
| [blink_MHz](blink_MHz/) | LED を点灯させる |
| [blink_1Hz](blink_1Hz/) | LED を点滅させる |
| [PWM](PWM/) | デューティ比を指定して LED の明るさを変える |
| [PWM_of_pico-examples](PWM_of_pico-examples/) | [pico-examples の pio/pwm](https://github.com/raspberrypi/pico-examples/tree/master/pio/pwm) を TinyGo に移植したもの |

## 必要なもの

- Raspberry Pi Pico
- [TinyGo](https://tinygo.org/)
- `pioasm` (Go 出力に対応したもの。[pico-sdk](https://github.com/raspberrypi/pico-sdk) の `tools/pioasm`)

## 使い方

```sh
pioasm -o go blink_1Hz/blink.pio blink_1Hz/blink_pio.go
tinygo flash --target=pico ./blink_1Hz
```

Makefile を用意しているため、make blink_MHz という形で実行することも可能です。

## ライセンス

[MIT License](LICENSE)

ただし [PWM_of_pico-examples/pwm.pio](PWM_of_pico-examples/pwm.pio) は [pico-examples](https://github.com/raspberrypi/pico-examples) 由来のため、BSD-3-Clause (Copyright (c) 2020 Raspberry Pi (Trading) Ltd.) に従います。
