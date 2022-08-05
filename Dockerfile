FROM alpine:latest

WORKDIR /bin

RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

COPY ./account-srv /bin

LABEL app-name=account-srv