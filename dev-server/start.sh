sudo service goweb stop
sudo rm -R ./lambda

echo "INIT START"
sudo mkdir ./schemas/digitalhub/lambda/graph
sudo mkdir ./schemas/digitalhub/lambda/models/form
sudo mkdir ./schemas/digitalhub/lambda/models/grid
sudo chmod -R 777 ./schemas/digitalhub/lambda/graph
sudo chmod -R 777 ./schemas/digitalhub/lambda/models/form
sudo chmod -R 777 ./schemas/digitalhub/lambda/models/grid
go run ./bootstrap/init/init.go
sudo chmod -R 777 ./schemas/digitalhub/lambda/graph
sudo chmod -R 777 ./schemas/digitalhub/lambda/models/form
sudo chmod -R 777 ./schemas/digitalhub/lambda/models/grid
sudo chmod -R 755 ./schemas/digitalhub/lambda/graph
sudo chmod -R 755 ./schemas/digitalhub/lambda/models/form
sudo chmod -R 755 ./schemas/digitalhub/lambda/models/grid
sudo mv ./schemas/digitalhub/lambda ./
echo "Ready"
go build main.go


sudo service goweb start
