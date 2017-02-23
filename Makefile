MVN_CMD := mvn

test-native:
	cd native; \
	$(MVN_CMD) test

build-native:
	cd native; \
	$(MVN_CMD) package

include .sdk/Makefile
