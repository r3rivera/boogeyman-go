FROM golang:latest


LABEL Owner="r.ryan.rivera@outlook.com"
LABEL Description="Boogeyman GO Application"

WORKDIR /app 

COPY go.mod go.sum ./

RUN go mod download 

COPY . .

RUN go build -o main .

EXPOSE 8082

ENTRYPOINT ["./main"]