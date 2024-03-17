#!/bin/bash
#https://stackoverflow.com/questions/31721086/docker-backup-and-restore-postgres

#First time

docker pull postgres:9.6
docker run -v /var/lib/postgresql/data --name postgresdata postgres:9.6 /bin/true

#Restart
docker stop postgres
docker rm postgres
docker run -d --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=12345678a --volumes-from=postgresdata postgres:9.6