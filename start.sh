#!/bin/bash


if [ -x /app/lambda/ ]; then
 # Check if main exists and is an executable file
 if [ -x main ]; then

   rm -R /app/lambda/
   mkdir /app/lambda/
   chmod -R 700 /app/lambda/
   ./main generate

 fi

else
    if [ $OSTYPE == "msys" ]; then
        rm -R -f ./lambda
        mkdir ./lambda
        ./main generate
    else
           sudo rm -R lambda/
           mkdir lambda/
           chmod -R 700 lambda/
           ./main generate
    fi
fi


# Rebuild the Go app
go build -o main .



# Run the new binary
exec ./main
