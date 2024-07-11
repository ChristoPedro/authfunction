FROM oraclelinux:9-slim

RUN microdnf -y install go-toolset && \
    microdnf clean all
ENV GO111MODULE=on
RUN groupadd --gid 1000 fn && \
    adduser --uid 1000 --gid fn fn
WORKDIR /function
ADD . /function/
RUN go build -o func

ENTRYPOINT ["./func"]