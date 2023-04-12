## Docker commands
sudo docker logs traefik
sudo docker restart traefik
sudo docker logs superset_app
sudo docker ps
sudo docker compose down
sudo docker compose up -d
sudo docker logs food-prod
sudo docker images
sudo docker rmi 3dc0e49b0ef5
## Show container IP
sudo docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' postgres-server
sudo docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' superset_app
sudo docker inspect postgres-server
sudo docker inspect superset_app

## super set allow DB connection
sudo docker compose -f docker-compose-non-dev.yml up -d

sudo apt-get install nmap
sudo ufw allow from 172.21.0.0/16 to any port 5432 proto tcp

ifconfig
eth0 ip use

## Delete or show status 
sudo ufw status numbered
sudo ufw delete 6


## Create Docker network: UNUSED !!!
sudo docker network ls
sudo docker network rm web
sudo docker network create --subnet=172.21.0.0/16 web

## Connect to docker Container: UNUSED
sudo docker exec -it food-prod sh
printenv | grep DB_HOST