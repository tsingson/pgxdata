# pgxdata

 a tool to generate a [jackc/pgx](https://github.com/jackc/pgx) database go package  tailored to exists postgresql database schema.



## Usage

### **STEP 1**  install pgxdata 

first install pgxdata

```
go get -u github.com/jackc/pgxdata
```

### **STEP 2**  setting path 

add GOPATH/bin to PATH, or move pgxdata to /usr/local/bin ..etc , 
make sure pgxdata runing anywhere


### **STEP 3** create schema 

 then create a postgresql schema like this

```
    createdb pgxdata
    psql pgxdata -f test/structure.sql
```


### **STEP 4**  create package 
run pgxdata once to create a new go package name "dbi"  or some other package name within $GOPATH/src
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


### **STEP 5**  edit config 
edit the config file dbi/config.toml
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


### **STEP 6**  run generate 
goto dbi directory and run generate
```
cd dbi
pgxdata generate
```

### **Final**  review generated code 
check $GOPATH/src/dbi to see all generated go package for postgres database pgxdata
```
ls $GOPATH/src/dbi

8596271720 drwxr-xr-x  12 qinshen  staff    384 Jul 22 01:17 ./
8596271719 drwxr-xr-x   5 qinshen  staff    160 Jul 22 01:40 ../
8596271721 -rw-r--r--   1 qinshen  staff    505 Jul 22 01:17 config.toml
8596271722 -rw-r--r--   1 qinshen  staff  13952 Jul 22 01:17 crud_test.go
8596271723 -rw-r--r--   1 qinshen  staff    909 Jul 22 01:17 integration_test.go
8596271724 -rw-r--r--   1 qinshen  staff   2929 Jul 22 01:17 pgxdata_blob.go
8596271725 -rw-r--r--   1 qinshen  staff   4203 Jul 22 01:17 pgxdata_customer.go
8596271726 -rw-r--r--   1 qinshen  staff   1817 Jul 22 01:17 pgxdata_db.go
8596271727 -rw-r--r--   1 qinshen  staff   3024 Jul 22 01:17 pgxdata_part.go
8596271728 -rw-r--r--   1 qinshen  staff   4491 Jul 22 01:17 pgxdata_renamed_field_customer.go
8596271729 -rw-r--r--   1 qinshen  staff   3651 Jul 22 01:17 pgxdata_semester.go
8596271730 -rw-r--r--   1 qinshen  staff   3283 Jul 22 01:17 pgxdata_widget.go

```


## Testing

Create a test database and populate it with the test schema.

    createdb pgxdata
    psql pgxdata -f test/structure.sql

Set PG* envvars when running tests to pass connection information to tests.

    PGHOST=/var/run/postgresql PGDATABASE=pgxdata rake

Regenerate the test app in test/dbi as needed (probably need to automate this).
