FROM golang:latest
MAINTAINER "shadowedge"


ENV GO111MOUDLE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GIN_MODE=release \
    GOPROXY="https://goproxy.cn,direct"


WORKDIR /webwork
COPY . .
RUN go get -d -v ./...
RUN go build -o app . 
WORKDIR /dist
RUN mkdir html
RUN mkdir css
RUN cp /webwork/app .
RUN cp /webwork/config.yaml .
RUN cp /webwork/html/* ./html
RUN cp /webwork/css/* ./css
RUN mkdir static 
WORKDIR /dist/static
RUN mkdir img
WORKDIR /dist
EXPOSE 9090
CMD ["/dist/app"]