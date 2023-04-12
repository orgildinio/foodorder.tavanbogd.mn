#!/bin/bash


if [ $OSTYPE == "msys" ]; then
    rm -R -f ./lambda
    echo "INIT START"
    mkdir ./lambda
else
    sudo rm -R ./lambda
    echo "INIT START"
    mkdir ./lambda
fi

go run ./bootstrap/init/init.go


# Rebuild the Go app
go build -o main .



# Run the new binary
exec ./main
