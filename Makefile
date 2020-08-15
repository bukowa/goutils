DOCKER_TAG="goutils:tests"
SUBMODULES := http kubernetes utils

gen-makefiles:
	"$(MAKE)" for-module-cmd CMD="./Makefile.generate.sh"

gofmt:
	go fmt ./...

gotest:
	go test -v  ./...

for-module-cmd:
	for module in ${SUBMODULES} ; do \
	$(CMD) $$module ; \
	done

for-makefile:
	for module in ${SUBMODULES}; do \
  	"$(MAKE)" -C $$module $(ARGS) || exit 1; \
  	done

#git-perm:
#	git update-index --chmod=+x \
#	./_tests/pkg/server_test.sh \
#	&& git commit -m fix file perm