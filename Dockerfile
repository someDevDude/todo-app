FROM golang:latest
RUN mkdir /app
ADD . /app
WORKDIR /app

# needed when dependencies are added
COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV PORT=8080

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE $PORT

# Command to run the executable
CMD ["/app/main"]