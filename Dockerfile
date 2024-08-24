FROM golang:1.22
WORKDIR /app

COPY . /app
RUN go build -o go-live-streaming main.go

ENTRYPOINT [ "./go-live-streaming" ]

EXPOSE 8000