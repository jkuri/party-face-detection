BUILD_DIR = build

all: build

build:
	@mkdir -p build
	@go build -o ${BUILD_DIR}/party-face-detection main.go
	@echo "=> Build done."

clean:
	rm -rf ${BUILD_DIR}
