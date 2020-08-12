FROM golang:1.14

WORKDIR /app
COPY . .

RUN ls
ENTRYPOINT ["/bin/bash"]
CMD ["./_tests/pkg/server_test.sh"]