FROM golang:latest

ENV SRC_DIR=/go/src/github.com/ioannisGiak89/foodHygiene_IG

COPY . $SRC_DIR

RUN cd $SRC_DIR \
    && go build -o main \
    && cp docker-cmd.sh /usr/bin \
    && chmod +x /usr/bin/docker-cmd.sh

CMD ["docker-cmd.sh"]
