# from golang 19 alpine
FROM golang:1.19-alpine as builder

# service http port
ARG HTTP_PORT=8080

# src directory
WORKDIR /app/src/

# copy all files
COPY . .

# build golang application
RUN go build -o ./runner

FROM alpine:latest

WORKDIR /app/

COPY --from=builder ./runner .

CMD ["./runner $HTTP_PORT"]