FROM golang:latest


LABEL Owner="r.ryan.rivera@outlook.com"
LABEL Description="Boogeyman GO Application"

WORKDIR /app 

COPY go.mod go.sum ./

RUN go mod download 
RUN openssl genpkey -algorithm RSA -out /var/local/private1_key.pem -pkeyopt rsa_keygen_bits:2048
RUN chmod 755 /var/local/private1_key.pem 
RUN openssl rsa -pubout -in /var/local/private1_key.pem -out /var/local/public1_key.pem
RUN chmod 755 /var/local/public1_key.pem 

COPY . .

RUN go build -o main .

EXPOSE 8082

ENTRYPOINT ["./main"]