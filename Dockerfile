FROM golang:latest AS build
RUN mkdir /app
ADD . /app
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o ./build/todo-server ./cmd/todo-server


FROM alpine
COPY --from=build /app/build/todo-server ./
EXPOSE 8080
RUN ls
RUN ls todo-server
CMD ["./todo-server"]