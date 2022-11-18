#sudo rm -R ./lambda
#echo "INIT START"
#sudo mkdir ./lambda
#sudo chmod -R 777 ./lambda
#go run ./bootstrap/init/init.go
#sudo chmod -R 755 ./lambda
#echo "Ready"
#go run main.go

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
echo "Ready"
go run main.go
