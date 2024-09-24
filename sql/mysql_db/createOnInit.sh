#!/bin/bash

set -e

mysql -v -u root -p$MYSQL_ROOT_PASSWORD <<-EOSQL
    CREATE DATABASE IF NOT EXISTS dev_db;
    CREATE DATABASE IF NOT EXISTS prod_db;
EOSQL

#!/bin/bash

# set -e

# mysql -v -u root -p$MYSQL_ROOT_PASSWORD <<-EOSQL
#     CREATE DATABASE IF NOT EXISTS dev_db;
#     CREATE DATABASE IF NOT EXISTS prod_db;
#     GRANT ALL PRIVILEGES ON dev_db.* TO 'user'@'%' IDENTIFIED BY '$MYSQL_PASSWORD';
#     GRANT ALL PRIVILEGES ON prod_db.* TO 'user'@'%' IDENTIFIED BY '$MYSQL_PASSWORD';
#     FLUSH PRIVILEGES;
# EOSQL