DOCKER_TAG="goutils:tests"

build:
	docker build --tag=${DOCKER_TAG} .

test:
	docker run --rm -it --add-host="testxx:127.0.0.1" ${DOCKER_TAG}
