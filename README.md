# h2dump
this app can dump mysql table struct to h2 supported gramma

# usage
## help
./h2dump -h
flag needs an argument: -h
Usage of ./h2dump:
  -P string
    	mysql port (default "3306")
  -d string
    	mysql database (default "mysql")
  -f string
    	script dir (default "db.sql")
  -h string
    	mysql host (default "127.0.0.1")
  -p string
    	mysql password (default "root")
  -u string
    	mysql user (default "root")
## dump
./h2dump -h 127.0.0.1 -P 3306 -u root -p root -d example_db -f db.sql

## dump on local mysql, the simple param seems good, and other param used default value
./h2dump -d example_db
