FROM golang:1.14

WORKDIR /app/goutils
COPY . .

RUN ls
ENTRYPOINT ["/bin/bash", "-c"]
CMD ["./_tests/pkg/server_test.sh"]