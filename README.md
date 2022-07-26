# todo-go
A simply activity


### Database
* Installing PostgreSQL
```bash
$ sudo apt update

Then, install the Postgres package along with a -contrib package that adds some additional utilities and functionality:

$ sudo apt install postgresql postgresql-contrib
```
* Ensure that the server is running using the systemctl start command:
```bash
sudo systemctl start postgresql.service
```
* Switch over to the postgres account on your server by typing:
```bash
sudo -i -u postgres
```
* You can now access the PostgreSQL prompt immediately by typing:
```bash
psql
```
* `Exit out` of the PostgreSQL prompt by typing:
```bash
\q
```
###  Creating a New Database
```bash
create database <dbname>
```
* check database list
```bash
\list
```
* how to delete database
```bash
drop database <dbname>
```
