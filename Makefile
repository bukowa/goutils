DOCKER_TAG="goutils:tests"
SUBMODULES := http kubernetes utils storage/bolt

gen-makefiles:
	"$(MAKE)" for-module-cmd CMD="./Makefile.generate.sh"

#gofmt:
#	"$(MAKE)" for-module-cmd CMD="./gofmt.sh"

gofmt:
	"$(MAKE)" for-module-cmd CMD="./gofmt.sh"

gotest:
	"$(MAKE)" for-module-cmd CMD="./gotest.sh"

for-module-cmd:
	for module in ${SUBMODULES} ; do \
	$(CMD) $$module $(ARGS) || exit 1 ; \
	done

git-perm:
	git update-index --chmod=+x \
	./gofmt.sh \
	./gotest.sh \
	&& git commit -m fix file perm