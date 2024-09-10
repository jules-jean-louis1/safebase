Monter un volume pour la persistance des données :

```bash
docker run --name safebase-db \
-e DB_USER=$DB_USER \
-e DB_PASSWORD=$DB_PASSWORD \
-e DB_NAME=$DB_NAME \
-v /my/own/datadir:/var/lib/postgresql/data \
-p 5432:5432 \
-d postgres
```
Se connecter a Mysql : 
```bash
http://localhost:8080/test-connection?host=mysql_db&port=3306&username=root&password=mysql_root_password&dbName=bibliotheque&dbType=mysql
```
Se connecter a Postgres : 
```bash
http://localhost:8080/test-connection?host=postgres_db&port=5432&username=postgres&password=password&dbName=postgresdb&dbType=postgres ou dev ou prod
```
