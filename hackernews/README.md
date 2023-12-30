docker run -p 3306:3306 --name mysql -e MYSQL_ROOT_PASSWORD=dbpass -e MYSQL_DATABASE=hackernews -d mysql:latest

go get -u github.com/go-sql-driver/mysql

go install -tags 'postgres' -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

cd internal/pkg/db/migrations/
~/go/bin/migrate create -ext sql -dir mysql -seq create_users_table
~/go/bin/migrate create -ext sql -dir mysql -seq create_links_table

~/go/bin/migrate -database mysql://root:dbpass@/hackernews -path internal/pkg/db/migrations/mysql force 2

```sql
CREATE TABLE Users (ID INT NOT NULL UNIQUE AUTO_INCREMENT, Username VARCHAR (127) NOT NULL UNIQUE, Password VARCHAR (127) NOT NULL, PRIMARY KEY (ID));
CREATE TABLE Links (ID INT NOT NULL UNIQUE AUTO_INCREMENT, Title VARCHAR (255), Address VARCHAR (255), UserID INT, FOREIGN KEY (UserID) REFERENCES Users(ID), PRIMARY KEY (ID));
```