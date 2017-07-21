# pgxdata

 a tool to generate a jackc/pgx database go package  tailored to exists postgresql database schema.

## Usage
**STEP 1** first install pgxdata
```
go get -u github.com/jackc/pgxdata
```

**STEP 2** add GOPATH/bin to PATH , make sure pgxdata runing anywhere


**STEP 3** then create a postgresql schema like this
```
    createdb pgxdata
    psql pgxdata -f test/structure.sql
```


**STEP 4** run pgxdata once to create a new go package name "data" within $GOPATH/src
```
   cd $GOPATH/src
   pgxdata init dbi
```

this command will create a directory name "dbi" and two file inside "dbi"

```
ls ./dbi
config.toml
pgxdata_db.go
```

cat the ./dbi/config.toml like this
```
cat ./dbi/config.toml
package = "dbi"

# Database connection information can be specified here or in PG* environment variables
#
# [database]
# host = "127.0.0.1"
# port = 5432
# database = "myapp_development"
# user = "myuser"
# password = "secret"

[[tables]]
table_name = "customer"
# struct_name = "Customer"
```


**STEP 5** edit the config file dbi/config.toml
```
vi dbi/config.toml
```

with this look like
```
package = "dbi"
[database]
host = "127.0.0.1"
port = 5432
database = "pgxdata"
user = "myuser"
password = "secret"

[[tables]]
table_name = "customer"
struct_name = "Customer"

[[tables]]
table_name = "widget"
struct_name = "Widget"

[[tables]]
table_name = "part"
struct_name = "Part"
primary_key = ["code"]

[[tables]]
table_name = "semester"
struct_name = "Semester"
primary_key = ["year", "season"]

[[tables]]
table_name = "customer"
struct_name = "RenamedFieldCustomer"

  [[tables.columns]]
  column_name = "first_name"
  field_name = "FName"

[[tables]]
table_name = "blob"
struct_name = "Blob"

```


**STEP 6** goto dbi directory and run generate
```
cd dbi
pgxdata generate
```

## Testing

Create a test database and populate it with the test schema.

    createdb pgxdata
    psql pgxdata -f test/structure.sql

Set PG* envvars when running tests to pass connection information to tests.

    PGHOST=/var/run/postgresql PGDATABASE=pgxdata rake

Regenerate the test app in test/dbi as needed (probably need to automate this).
