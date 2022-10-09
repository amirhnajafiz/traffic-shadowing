<center>
<img src="assets/logo.png" width="500" />
</center>

<h1 align="center">
    Traffic Shadowing
</h1>

Using **Goreplay** to shadow http service traffic. (with Docker container example)

**Gor** is an open-source tool for 
capturing and replaying live HTTP traffic into a 
test environment in order to continuously test your 
system with real data. 

It can be used to increase confidence in code 
deployments, configuration changes and infrastructure changes.

- [Documentation for Gor](https://github.com/buger/goreplay/wiki)
- [Official Website](https://goreplay.org/shadowing.html)

In this project we are going to build two http services with
_golang echo library_ and then shadow our http requests between these
services with help of **Gor**.

## Requirements

Make sure to have **Golang** ```version 19+``` installed on your system.

After that you need to have **Gor** installed on your system. You can
install it on MacOS with _brew_:

```shell
brew install gor
```

Or you can clone the original repository and build it:
```shell
# installing gor on your local system
curl https://github.com/buger/goreplay/archive/refs/tags/v2.0.0-rc2.tar.gz -O -J -L

# unzip the downloaded file
tar -xf goreplay-2.0.0-rc2.tar.gz
```

Use the following script:
```shell
chmod +x setup/raw/install.sh
./setup/raw/install.sh
```

Now test the installation:
```shell
gor
```

```shell
2022/10/09 10:09:23 [PPID 2128 and PID 3093] Version:1.3.0
2022/10/09 10:09:23 Required at least 1 input and 1 output
```

## Start servers
You can start echo server with the following command:
```shell
go run internal/cmd/main.go [port]
```

```shell
go run internal/cmd/main.go 8080
go run internal/cmd/main.go 8081
```

Or you can use docker compose to create two services on docker:
```shell
docker compose up -d
```

### Shadow command
Now with the following command you can shadow http requests
that are sent to ```localhost:8080``` to ```localhost:8081```.

```shell
sudo gor --input-raw :8080 --output-http http://localhost:8081
```

Now try to make http requests on ```localhost:8080```.

```shell
curl localhost:8080
```

There are two log files named ```log8080.txt``` and ```log8081.txt```. You can
see that all requests that we send to first service is also being sent to other service.
