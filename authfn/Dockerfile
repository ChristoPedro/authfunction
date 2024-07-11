FROM golang:alpine as build
WORKDIR /function
ADD . /function
RUN go build -o func
FROM alpine
ENV GO111MODULE=on
RUN addgroup --gid 1000 fn &&\
    adduser -u 1000 -G fn fn -D
WORKDIR /function
COPY --from=build /function/func /function/
ENTRYPOINT ["./func"]