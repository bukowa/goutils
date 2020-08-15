DOCKER_TAG="goutils:tests"
SUBMODULES := http kubernetes utils

gen-makefiles:
	"$(MAKE)" for-module ARGS="./Makefile.generate.sh"

for-module:
	for module in ${SUBMODULES} ; do \
	$(ARGS) $$module ; \
	done

for-makefile:
	for module in ${SUBMODULES}; do \
  	"$(MAKE)" -C $$module $(ARGS) || exit 1; \
  	done

gotest:
	"$(MAKE)" for-makefile ARGS="gotest"

build:
	docker build --tag=${DOCKER_TAG} .

test_scripts: build
	docker run --rm --add-host="testxx:127.0.0.1" ${DOCKER_TAG}

test_pkg: build
	docker run --rm --entrypoint="" --add-host="testxx:127.0.0.1" ${DOCKER_TAG} /bin/bash -c "go test -v  ./..."

git-perm:
	git update-index --chmod=+x \
	./_tests/pkg/server_test.sh \
	&& git commit -m fix file perm