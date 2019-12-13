FROM golang:latest
RUN mkdir /app
ADD . /app
WORKDIR /app

COPY . .

RUN go mod download

ENV PORT=8080

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE $PORT

# Command to run the executable
CMD ["/app/main"]