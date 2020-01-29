FROM golang:1.13 as build
COPY . .
RUN go get -u github.com/go-delve/delve/cmd/dlv

RUN ls -r
ENV GO111MODULE on
RUN go build -o /service -gcflags "all=-N -l" *.go
CMD dlv --headless --listen=:2345 --api-version=2 --accept-multiclient exec /service