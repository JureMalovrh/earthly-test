FROM golang:1.16.3-buster as builder
RUN mkdir /app
WORKDIR /app
ADD ./main.go /app/.
ADD ./go.mod /app/.
ADD ./go.sum /app/.
RUN CGO_ENABLED=0 go build -o main main.go

FROM alpine:latest as prod
RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/main .
ENTRYPOINT [ "/app/main" ]