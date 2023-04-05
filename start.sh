sudo service mmk stop
sudo chmod -R 777 ./lambda
rm -R ./lambda

echo "INIT START"

mkdir ./lambda
mkdir ./lambda/models
mkdir ./lambda/graph
mkdir ./lambda/schemas
mkdir ./lambda/schemas/form
mkdir ./lambda/models/form
mkdir ./lambda/models/grid
chmod -R 777 ./lambda
go run ./bootstrap/init/init.go
chmod -R 755 ./lambda

echo "Ready"
go build main.go

sudo service mmk start
