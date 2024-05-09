
# FROM golang:latest

# RUN go get -u github.com/cosmtrek/air

# WORKDIR /app

# COPY . .

# RUN go build -o main .

# EXPOSE 3000

# # Command to run the executable
# # CMD ["./main"]
# ENTRYPOINT ["air"]

FROM golang:latest
RUN go install github.com/cosmtrek/air@latest
WORKDIR /app
EXPOSE 3000
ENTRYPOINT ["air"]