golang standard app
-------------------

A set of standards with an example of how to create a well designed golang app.

tools
-----

* [Dep](https://github.com/golang/dep) the future official standard for golang package management

* [MetaLinter](https://github.com/alecthomas/gometalinter) combine and run golang code check tools

packages
--------

* [DocOpt](https://github.com/docopt/docopt.go) for command line parsing. It has documentation and command line parsing all in one!

* [Echo](https://github.com/labstack/echo) A minimalist web framework with good documentation

* [Viper](https://github.com/spf13/viper) Define your config and use multi-format config


project layout
--------------

* cmd (code that interacts with the command line goes here)
    * binary-name
        * main.go
* pkg (all other application code goes here)
    * config (all config related code here)
    * api (put together all the pieces for running the api)
        * api.go
    * controllers (any code that interacts directly with the api)
        * foo.go
    * models (any code that represents an object)
        * foo.go
    * services (business logic code that sits between the controllers and other libraries)
        * foo.go
    * db (any code that interacts with outside databases)
        * rethinkdb
            * foo.go
        *  etcd
            * foo.go
        * foo.go (interface for foo database code)

    * etc... (any other logical grouping of code that you can come up with)
* vendor (All of your dependencies go here. This should never be populated manually. Use dep to generate it)
* static (Web assets go here HTML/JS/CSS)
* .gitignore
* Dockerfile (use this to run the project in a container)
* lock.json (used by dep to keep dependcies locked until you update them)
* manifest.json (Define your vendor dependencies here)
* README.md (you should always have one of these)


Run rethinkdb in a container on localhost

    docker run -d --name rethinkdb -p 8080:8080 -p 28015:28015 rethinkdb
