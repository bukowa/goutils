DOCKER_TAG="goutils:tests"

build:
	docker build --tag=${DOCKER_TAG} .

test_scripts:
	docker run --rm --add-host="testxx:127.0.0.1" ${DOCKER_TAG}

test_pkg:
	docker run --rm --entrypoint="" --add-host="testxx:127.0.0.1" ${DOCKER_TAG} /bin/bash -c "go test -v ./..."

test: test_scripts test_pkg
