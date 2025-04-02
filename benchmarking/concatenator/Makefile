.PHONY: bench-prepare-env
bench-prepare-env:
	sudo systemctl stop snapd.service
	sudo systemctl stop unattended-upgrades.service
	sudo systemctl stop thermald
	sudo systemctl stop power-profiles-daemon
	sudo cpupower frequency-set -g performance
	echo 0 | sudo tee /sys/devices/system/cpu/intel_pstate/no_turbo
	sudo cpupower idle-set -D 0
	sudo perflock -governor none -daemon

.PHONY: bench-restore-env
bench-restore-env:
	sudo kill `pgrep perflock`
	sudo cpupower frequency-set -g powersave
	echo 1 | sudo tee /sys/devices/system/cpu/intel_pstate/no_turbo
	sudo cpupower idle-set -E
	sudo systemctl start thermald
	sudo systemctl start power-profiles-daemon
	sudo systemctl start snapd.service
	sudo systemctl start unattended-upgrades.service