# Got Cep

A simple cep getter from a API

## Dependencies

* Docker 18.09.0 or newer
* Docker Compose 1.22.0 or newer

## How Dev

Run the commands bellow to build the project

1. `docker-compose build`
2. `docker-compose up`

The project will be accessible on [http://localhost:9000](http://localhost:9000).

After the fist build just run `docker-compose up`.

### Code Layout

It will use Revel classic struct.

The directory structure of a generated Revel application:

    conf/             Configuration directory
        app.conf      Main app configuration file
        routes        Routes definition file

    app/              App sources
        init.go       Interceptor registration
        controllers/  App controllers go here
        views/        Templates directory

    messages/         Message files

    public/           Public static assets
        css/          CSS files
        js/           Javascript files
        images/       Image files

    tests/            Test suites

