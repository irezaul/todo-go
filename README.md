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
###  Creating a New Role
* If you are logged in as the postgres account, you can create a new user by typing:
```bash
createuser --interactive
```
#### Important- ( Currently, you just have the postgres role configured within the database. You can create new roles from the command line with the `createrole` command. The `--interactive` flag will prompt you for the name of the new role and also ask whether it should have superuser permissions. )
