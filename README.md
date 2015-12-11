# docker-go-tutorial
quick demonstration of using docker and go

## Docker 101

We use docker to simplify recreating isolated development environments. It also
allows us to precisely define the environment the application will run in
production.

**Docker** is the low level tool to manage these environments. While docker's
cli and api are simple, setting up a full environment with database and other
dependencies could take several steps. These steps are combined and simplified
by docker-compose, which is a tool on top of docker orchastrating the full
environment.

**Docker-compose** reads a config file (usually `docker-compose.yml`) and based
on this file downloads or builds docker containers, then sets up the environment
with network connection between these containers and envirnment variables set.

When you execute commands with docker-compose, you need to specify which
container you want to spin up. Example: `docker-compose run app bundle`.

Docker compose can start up the whole project with `docker-compose up`. You can
also start specific services with start command i.e. `docker-compose start db`.

**Docker machine** is the tool for managing virtual machines. It essentially
creates a very small linux box (only a kernel + docker) and sets up your local
shell environment with that box. Docker machine can also work with cloud
providers, so you can easily move your box to a cloud service. We recommend
using virtualbox on your dev machine.

Important commands:

- Creating a box: `docker-machine create <name of the box>`
- Start the box: `docker-machine start <box>`
- Set up your local shell environment: `eval $(docker-machine env <box>)`


Additional resources:

- docker-compose.yml reference: https://docs.docker.com/compose/compose-file/
- docker-compose overview: https://docs.docker.com/compose/reference/overview/
- docker-machine (also check out the subcommands section on the left side): https://docs.docker.com/machine/drivers/virtualbox/


## Installation

### Set up Docker Machine & Compose

```
brew install docker-compose docker-machine
docker-machine create dev --driver=virtualbox
docker-machine start dev
eval $(docker-machine env dev)
```

### Run project

```
docker-compose up
```

