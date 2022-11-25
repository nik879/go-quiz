# go-quiz

Simple plain golang API for quizes.

### Features

- Authentication via jwt
- Creating and answering questions
- Implements migration functionalities

### Getting started

**Requirements**: `docker` and `docker-compose`.

Just type the command below on the root folder of the project. The `docker-compose`
will fire up a MySQL 8 server and run the migrations.

    $ docker-compose up -d

After, you can run your code locally using `localhost:3306`.

The default environment variables for connecting with the database are:

- DB_CONNECTION=mysql
- DB_HOST=tcp(mysqlsrv)
- DB_DATABASE=quizdb
- DB_USERNAME=quizusr
- DB_PASSWORD=quizpasswd



#### code scanning test

The commited code is also being tested with a code scanning tool