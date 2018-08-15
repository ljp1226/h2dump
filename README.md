# h2dump
this app can dump mysql table struct to h2 supported gramma

# usage
## help:
./h2dump -h</br>
flag needs an argument: -h</br>
Usage of ./h2dump:</br>
  -P string</br>
    	mysql port (default "3306")</br>
  -d string</br>
    	mysql database (default "mysql")</br>
  -f string</br>
    	script dir (default "db.sql")</br>
  -h string</br>
    	mysql host (default "127.0.0.1")</br>
  -p string</br>
    	mysql password (default "root")</br>
  -u string</br>
    	mysql user (default "root")</br>
## dump:
./h2dump -h 127.0.0.1 -P 3306 -u root -p root -d example_db -f db.sql

## dump with the simple param if local env, and other param used default value
./h2dump -d example_db

## build on three platform
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go</br>
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go</br>
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go</br>
