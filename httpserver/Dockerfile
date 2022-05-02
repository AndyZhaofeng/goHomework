FROM golang:1.18-alpine
WORKDIR /app
COPY ./src/server/go.mod /app/
RUN go mod download
COPY ./src/server/main.go /app/
RUN go build /app/main.go
EXPOSE 8080
CMD [ "/app/main" ]
