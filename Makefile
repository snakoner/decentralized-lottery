.PHONY: build
.default := build

CONTRACT_NAME = DecentralizedLotteryV2

build:
	docker-compose build

generate_go_bindings:
	solc --base-path . --include-path node_modules --overwrite --abi --bin -o build contracts/$(CONTRACT_NAME).sol
	abigen --bin=build/$(CONTRACT_NAME).bin --abi=build/$(CONTRACT_NAME).abi --pkg=lottery --out=build/lottery.go
	cp ./build/lottery.go ./backend/internal/lottery

deploy:
	rm -rf ./ignition/deployments
	npx hardhat ignition deploy ./ignition/modules/Lottery.ts --show-stack-traces --network sepolia

up:
	docker-compose up

stop:
	docker-compose down

rmi_last_image:
	docker rmi $(shell docker images -q | head -n 1)

rm_last_container:
	docker rm -f $(shell docker ps -a -q | head -n 1)