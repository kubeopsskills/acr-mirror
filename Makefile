TARGET_PATH = bin
GOARCH = GOARCH=amd64
VERSION = 1.0.1

buildWindows:
	env GOOS=windows $(GOARCH) go build -o ./$(TARGET_PATH)/windows/acr-mirror.exe github.com/kubeopsskills/acr-mirror/cmd/acr
	cd $(TARGET_PATH) && zip acr-mirror-Windows-$(VERSION).zip ./windows/acr-mirror.exe

buildMacOS:
	env GOOS=darwin $(GOARCH) go build  -o ./$(TARGET_PATH)/macos/acr-mirror github.com/kubeopsskills/acr-mirror/cmd/acr
	cd $(TARGET_PATH) && tar -zcvf acr-mirror-MacOS-$(VERSION).tar.gz ./macos/acr-mirror

buildLinux:
	env GOOS=linux $(GOARCH) go build -o ./$(TARGET_PATH)/linux/acr-mirror github.com/kubeopsskills/acr-mirror/cmd/acr
	cd $(TARGET_PATH) && tar -zcvf acr-mirror-Linux-$(VERSION).tar.gz ./linux/acr-mirror

build: buildWindows buildMacOS buildLinux

clean:
	rm -rf bin

all: clean build