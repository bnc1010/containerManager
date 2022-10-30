BIN_FILE=containerManager

hello:
	echo "Hello"

build:
	go build -o output/${BIN_FILE}
	cp -r conf output/

run:
	output/${BIN_FILE}

start:
	nohup make run > output/start_`date +%Y-%m-%d`.txt 2> output/run_`date +%Y-%m-%d`.txt &

stop:
	pidof ./${BIN_FILE} | xargs kill -9

restart: stop start

build_run: build run