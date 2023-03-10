sudo rm -R ./lambda
echo "INIT START"
sudo mkdir ./lambda
sudo chmod -R 777 ./lambda
go run ./bootstrap/init/init.go
sudo chmod -R 755 ./lambda
echo "Ready"
go run main.go
