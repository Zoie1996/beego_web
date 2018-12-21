docker pull mysql:5.7.23
docker run -p 3306:3306 --name mymysql -e MYSQL_ROOT_PASSWORD=root -d mysql:5.7.23