#  Pms 

Service pms is used for project management. 

## Configuration

The service is configured using the environment variables presented in the
following table. Note that any unset variables will be replaced with their
default values.

| Variable                   | Description                                                          | Default               |
|----------------------------|----------------------------------------------------------------------|-----------------------|
| MF_PMS_LOG_LEVEL         | Log level for twin service (debug, info, warn, error)                | error                 |
| MF_PMS_HTTP_PORT         | Twins service HTTP port                                              | 9021                  |
| MF_PMS_SERVER_CERT       | Path to server certificate in PEM format                             |                       |
| MF_PMS_SERVER_KEY        | Path to server key in PEM format                                     |                       |
| MF_JAEGER_URL              | Jaeger server URL                                                    |                       |
| MF_PMS_DB_NAME           | Database name                                                        | mainflux              |
| MF_PMS_DB_HOST           | Database host address                                                | localhost             |
| MF_PMS_DB_PORT           | Database host port                                                   | 27017                 |
| MF_PMS_SINGLE_USER_EMAIL | User email for single user mode (no gRPC communication with users)   |                       |
| MF_PMS_SINGLE_USER_TOKEN | User token for single user mode that should be passed in auth header |                       |
| MF_PMS_CLIENT_TLS        | Flag that indicates if TLS should be turned on                       | false                 |
| MF_PMS_CA_CERTS          | Path to trusted CAs in PEM format                                    |                       |
| MF_PMS_MQTT_URL          | Mqtt broker URL for twin CRUD and states update notifications        | tcp://localhost:1883  |
| MF_PMS_THING_ID          | ID of thing representing v2ms service & mqtt user                   |                       |
| MF_PMS_THING_KEY         | Key of thing representing v2ms service & mqtt pass                  |                       |
| MF_PMS_CHANNEL_ID        | Mqtt notifications topic                                             |                       |
| MF_NATS_URL                | Mainflux NATS broker URL                                             | nats://127.0.0.1:4222 |
| MF_AUTHN_GRPC_PORT         | Authn service gRPC port                                              | 8181                  |
| MF_AUTHN_TIMEOUT           | Authn gRPC request timeout in seconds                                | 1                     |
| MF_AUTHN_URL               | Authn service URL                                                    | localhost:8181        |

## Deployment

The service itself is distributed as Docker container. The following snippet
provides a compose file template that can be used to deploy the service container
locally:

```yaml
version: "3"
services:
  v2ms:
    image: v2ms:[version]
    container_name: [instance name]
    ports:
      - [host machine port]:[configured HTTP port]
    environment:
      MF_PMS_LOG_LEVEL: [Twins log level]
      MF_PMS_HTTP_PORT: [Service HTTP port]
      MF_PMS_SERVER_CERT: [String path to server cert in pem format]
      MF_PMS_SERVER_KEY: [String path to server key in pem format]
      MF_JAEGER_URL: [Jaeger server URL]
      MF_PMS_DB_NAME: [Database name]
      MF_PMS_DB_HOST: [Database host address]
      MF_PMS_DB_PORT: [Database host port]
      MF_PMS_SINGLE_USER_EMAIL: [User email for single user mode]
      MF_PMS_SINGLE_USER_TOKEN: [User token for single user mode]
      MF_PMS_CLIENT_TLS: [Flag that indicates if TLS should be turned on]
      MF_PMS_CA_CERTS: [Path to trusted CAs in PEM format]
      MF_PMS_MQTT_URL: [Mqtt broker URL for twin CRUD and states]
      MF_PMS_THING_ID: [ID of thing representing v2ms service]
      MF_PMS_THING_KEY: [Key of thing representing v2ms service]
      MF_PMS_CHANNEL_ID: [Mqtt notifications topic]
      MF_NATS_URL: [Mainflux NATS broker URL]
      MF_AUTHN_GRPC_PORT: [Authn service gRPC port]
      MF_AUTHN_TIMEOUT: [Authn gRPC request timeout in seconds]
      MF_AUTHN_URL: [Authn service URL]
```

To start the service outside of the container, execute the following shell script:

```bash
# download the latest version of the service
go get github.com/cloustone/pandas/mainflux

cd $GOPATH/src/github.com/cloustone/pandas/mainflux

# compile the v2ms
make v2ms

# copy binary to bin
make install

# set the environment variables and run the service
MF_PMS_LOG_LEVEL: [Twins log level]
MF_PMS_HTTP_PORT: [Service HTTP port] 
MF_PMS_SERVER_CERT: [String path to server cert in pem format] 
MF_PMS_SERVER_KEY: [String path to server key in pem format]
MF_JAEGER_URL: [Jaeger server URL]
MF_PMS_DB_NAME: [Database name] 
MF_PMS_DB_HOST: [Database host address] 
MF_PMS_DB_PORT: [Database host port] 
MF_PMS_SINGLE_USER_EMAIL: [User email for single user mode] 
MF_PMS_SINGLE_USER_TOKEN: [User token for single user mode] 
MF_PMS_CLIENT_TLS: [Flag that indicates if TLS should be turned on] 
MF_PMS_CA_CERTS: [Path to trusted CAs in PEM format] 
MF_PMS_MQTT_URL: [Mqtt broker URL for twin CRUD and states] 
MF_PMS_THING_ID: [ID of thing representing v2ms service] 
MF_PMS_THING_KEY: [Key of thing representing v2ms service]
MF_PMS_CHANNEL_ID: [Mqtt notifications topic]
MF_NATS_URL: [Mainflux NATS broker URL] 
MF_AUTHN_GRPC_PORT: [Authn service gRPC port] 
MF_AUTHN_TIMEOUT: [Authn gRPC request timeout in seconds]
MF_AUTHN_URL: [Authn service URL] $GOBIN/mainflux-v2ms
```

## Usage

### Starting v2ms service

The v2ms service publishes notifications on an mqtt topic of the format
`channels/<MF_PMS_CHANNEL_ID>/messages/<twinID>/<crudOp>`, where `crudOp`
stands for the crud operation done on twin - create, update, delete or
retrieve - or state - save state. In order to use twin service, one must
inform it - via environment variables - about the Mainflux thing and
channel used for mqtt notification publishing. You can use an already existing
thing and channel - thing has to be connected to channel - or create new ones.

To set the environment variables, please go to `.env` file and set the following
variables:

```
MF_PMS_THING_ID=
MF_PMS_THING_KEY=
MF_PMS_CHANNEL_ID=
```

with the corresponding values of the desired thing and channel. If you are
running mainflux natively, than do the same thing in the corresponding console
environment.

For more information about service capabilities and its usage, please check out
the [API documentation](swagger.yaml).

[doc]: http://mainflux.readthedocs.io