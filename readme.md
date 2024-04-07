# HTML Template Emailer built in GO

# Command to jump into docker container

docker exec -it gomail_db bash

# Enter mysql inside of container

mysql -h 127.0.0.1 -P 3306 -u root -p

# Change to snippetbox database

use gomail

# Change to snippets |



### Build command for windows
❯ GOOS=windows GOARCH=amd64 go build -o dist/gomail.exe mail.go
