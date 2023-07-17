# greenlight
Code of Let’s Go Further book by Alex Edwards

### Postgresql
Log in to database (first time):

```shell
docker exec -it postgresql-greenlight psql -U postgres
```

Create database:

```shell
CREATE DATABASE greenlight;
```

Connect to database:

```shell
\c greenlight
```

Create database user:

```shell
CREATE ROLE greenlight WITH LOGIN PASSWORD 'pa55word';
```

Add citext extension:
*Note*: This adds a case-insensitive character string type to PostgreSQL.

```shell
CREATE EXTENSION IF NOT EXISTS citext;
```

Log in with new user:

```shell
docker exec -it postgresql-greenlight psql --host=localhost --dbname=greenlight --username=greenlight
```

### Migrations

Install migration tool:

```shell
brew install golang-migrate
```

Grant privileges to user:

```shell
GRANT ALL ON SCHEMA public TO greenlight;
```
