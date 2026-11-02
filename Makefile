TARGET := pico
DIRS   := blink_MHz blink_1Hz PWM PWM_of_pico-examples

# .pio のファイル名（拡張子なし）。デフォルトは blink
PIO := blink
PWM_of_pico-examples: PIO := pwm

.PHONY: help $(DIRS)

help:
	@echo Usage: make [$(DIRS)]

# 例: make blink_1Hz
#   blink_1Hz/blink.pio から blink_1Hz/blink_pio.go を生成し、./blink_1Hz を書き込む
$(DIRS):
	pioasm -o go $@/$(PIO).pio $@/$(PIO)_pio.go
	tinygo flash --target=$(TARGET) ./$@
