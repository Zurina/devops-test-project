FROM golang:alpine as builder

WORKDIR /app 

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o devops main.go

FROM scratch

WORKDIR /app

COPY --from=builder /app/devops /usr/bin/

EXPOSE 8080

ENTRYPOINT ["devops"]