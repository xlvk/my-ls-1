FROM golang:1.20
COPY . /ghost-ls
WORKDIR /ghost-ls
RUN [ "make" ]
ENTRYPOINT [ "./myLs" ]
