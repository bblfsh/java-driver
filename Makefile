-include .sdk/Makefile

$(if $(filter true,$(sdkloaded)),,$(error You must install bblfsh-sdk))

MVN_CMD := mvn
JAR := native-jar-with-dependencies.jar

test-native-internal:
	cd native; \
	$(MVN_CMD) test

build-native-internal:
	cd native; \
	$(MVN_CMD) package
	cp native/target/$(JAR) $(BUILD_PATH); \
	cp native/src/main/sh/native.sh $(BUILD_PATH)/native; \
	chmod +x $(BUILD_PATH)/native

