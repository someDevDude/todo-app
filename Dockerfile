# STAGE 1: Build
from golang:latest AS build

# fetch dependencies
RUN go mod download

WORKDIR /go/src/github.com/someDevDude/todo-app

# Copy all sources in
COPY . .

RUN go build -o todo-server .

# STAGE 2: Runtime
FROM alpine

COPY --from=build /go/bin/todo-server /todo-server

CMD ["/todo-server"]
