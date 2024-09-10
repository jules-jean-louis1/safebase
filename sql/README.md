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
