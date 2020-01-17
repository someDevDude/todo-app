# build
FROM golang:1.13-alpine AS build
RUN mkdir /app
ADD . /app
WORKDIR /app
COPY . .
RUN go mod download
ENV GOOS=linux GOARCH=386
RUN go build -o ./build/todo-server ./cmd/todo-server

# run
FROM alpine
COPY --from=build /app/build/todo-server ./
EXPOSE 8080
CMD ["./todo-server"]
