TARGET := pico
DIRS   := blink_2MHz blink_1Hz PWM PWM_of_pico-examples

.PHONY: help $(DIRS)

help:
	@echo Usage: make [$(DIRS)]

# 例: make blink_1Hz
#   blink_1Hz/blink.pio から blink_1Hz/blink_pio.go を生成し、./blink_1Hz を書き込む
$(DIRS):
	pioasm -o go $@/blink.pio $@/blink_pio.go
	tinygo flash --target=$(TARGET) ./$@
