#!bin/bash

echo "#### Generating a Private Key ####"
openssl genpkey -algorithm RSA -out private_key.pem -pkeyopt rsa_keygen_bits:2048

echo "#### Generating a Certificate Signing Request (CSR) ####"
openssl req -new -key private_key.pem -out csr.pem

echo "#### Generating a Self-Signed Certificate ####"
openssl x509 -req -days 365 -in csr.pem -signkey private_key.pem -out certificate.pem
