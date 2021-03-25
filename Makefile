TARGET_PATH = bin
GOARCH = GOARCH=amd64
VERSION = 1.0.4
GOMODULE = github.com/kubeopsskills/acr-mirror/cmd/acr

buildWindows:
	env GOOS=windows $(GOARCH) go build -o ./$(TARGET_PATH)/windows/acr-mirror.exe $(GOMODULE)
	cd $(TARGET_PATH) && zip acr-mirror-Windows-$(VERSION).zip ./windows/acr-mirror.exe

buildMacOS:
	env GOOS=darwin $(GOARCH) go build  -o ./$(TARGET_PATH)/macos/acr-mirror $(GOMODULE)
	cd $(TARGET_PATH) && tar -zcvf acr-mirror-MacOS-$(VERSION).tar.gz ./macos/acr-mirror

buildLinux:
	env GOOS=linux $(GOARCH) go build -o ./$(TARGET_PATH)/linux/acr-mirror $(GOMODULE)
	cd $(TARGET_PATH) && tar -zcvf acr-mirror-Linux-$(VERSION).tar.gz ./linux/acr-mirror

buildARM:
	env GOOS=linux GOARCH=arm64 go build -o ./$(TARGET_PATH)/arm/acr-mirror $(GOMODULE)
	cd $(TARGET_PATH) && tar -zcvf acr-mirror-ARM-$(VERSION).tar.gz ./arm/acr-mirror

build: buildWindows buildMacOS buildLinux buildARM

clean:
	rm -rf bin

all: clean build
