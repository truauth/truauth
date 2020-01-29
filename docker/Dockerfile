FROM golang:1.13 as build
WORKDIR /app
COPY . .
ENV GO111MODULE on
RUN go build -v -o /service *.go
# Now copy it into our base image.
FROM gcr.io/distroless/base
COPY --from=build /service /service
CMD ["/service"]
