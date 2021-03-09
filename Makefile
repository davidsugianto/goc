BINARY=goc
BIN_DIR=./bin

install:
	go install ${BINARY}

build:
	go build -o ${BIN_DIR}/${BINARY} ./*.go

clean:
	if [ -f ${BIN_DIR}/${BINARY} ] ; then rm -rf ${BIN_DIR} ; fi

docker:
	docker build -t ${BINARY} .

run:
	docker-compose up --build -d

stop:
	docker-compose down

.PHONY: install clean unittest build docker run stop build
