
# Go examples

In this folder, you'll find example escape pod extensions written in Go. Please see the readmes in each of the directories for configuration examples, notes, etc.

|Plugin| Description |
|--|--|
| pkg/server | The service responsible for GRPC interactions |
| pkg/battery | Ask Vector for battery status |
| pkg/genericresponses | An example that makes vector say a canned phrase |
| pkg/selfaware | See what happens when you try to make Vector self aware. |
| pkg/openweathermap | A quick-and-dirty openweathermap.org processor |

## Super quick start

This section describes how to run the current services built with no changes whatsoever.

1. Make sure you've added the following lines to your escape pods config (in /etc/escape-pod.conf)

```sh
ENABLE_EXTENSIONS=true
ESCAPEPOD_EXTENDER_TARGET=${host ip you plan on running this software on}
ESCAPEPOD_EXTENDER_DISABLE_TLS=true
```

  If you're just now adding this, reboot the escape pod or restart the service.

2. Edit the .env file in this directory, replacing "CHANGEME" with appropriate values.

3. run ```$ make build```

4. run ```docker-compose up --build```

## Usage

This section describes how to run outside of docker

1. Export the appropriate environment variables

|Variable| Description |
|--|--|
| DDL_RPC_PORT | The TCP port you'd like to run the service on |
| DDL_RPC_INSECURE | For TLS-free communication |

For a full list of additional variables, please see the [hugh grpc server library viper loader](https://github.com/digital-dream-labs/hugh/blob/main/grpc/server/viper.go)

2. Export any other required environment variables for your application.

For example, if you'd like to use the extensions provided as is, you'd have to export the variables from the following repositories


* [pkg/battery](pkg/battery)
* [pkg/genericresponses](pkg/genericresponses)
* [pkg/selfaware](pkg/selfaware)
* [pkg/openweathermap](pkg/openweathermap)

3. Start your application!

```sh
# go run cmd/main.go
```