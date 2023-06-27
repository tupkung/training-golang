FROM --platform=linux/amd64 golang:1.20.5-alpine3.17

WORKDIR /app


COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/front-end-golang .

EXPOSE 3000

CMD ["/app/bin/front-end-golang"]
