FROM golang:latest

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN apk add --no-cache postgresql-client

ENV POSTGRES_HOST=localhost
ENV POSTGRES_PORT=5432
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=260702
ENV POSTGRES_DB=School

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 3030

CMD ["./main"]