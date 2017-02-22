# Config
LANGUAGE := java
NATIVE_RUNTIME_VERSION := 8.121.13-r0

# Java
MVN_CMD := mvn

test-native:
	cd native; \
	$(MVN_CMD) test

build-native:
	cd native; \
	$(MVN_CMD) package

include .sdk/Makefile