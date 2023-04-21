FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main

FROM alpine:3 
COPY --from=builder /app/main /bin
COPY --from=builder /app/static /static
EXPOSE 8080
ENTRYPOINT [ "/bin/main" ]