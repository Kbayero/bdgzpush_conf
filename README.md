# BitdefenderGravityZonePushConfiguration
### Version 1.2.0

*BitdefenderGravityZonePushConfiguration is a set of tools that allow you to correctly configure the Bitdefender GravityZone Cloud PUSH API*

## Getting started
_Before you can successfully configure your API, you must generate an APIKEY in your Bitdefender control account. For it:_
* Login to ***https://cloud.gravityzone.bitdefender.com/***
* Click on your username in the top right corner of the console and choose MY ACCOUNT
* Select the APIs you want to use, in this case it will be just Event Push Service API 
* Click on save, an ApiKey will be generated for the selected API

_This tool uses some configurations that you must define before starting, for this you must create an .env file in the root of this project, and define the following environment variables:_
* BDGZ_API_KEY= "APIKEY obtained in the previous process"
* BDGZ_ACCESS_URL= "You can find this url in the My Account/Control Center API/Access URL tab of your account, and add /v1.0/jsonrpc/push to it, in most cases it is ***https://cloud.gravityzone.bitdefender .com/api/v1.0/jsonrpc/push***"
* BDGZ_URL= "url of the server or log collector that will receive the Bitdefender logs, example: https://example.com:8000"

### Requirements
go 1.19
github.com/joho/godotenv

## Deploy

_First you must build your executable, for this:_

```
go build
```

_Then, to send your configuration to Bitdefender's Push API, you must run your executable passing it the sentConfig parameter._

```
.\bdgzpush_conf.exe sentConfig
```

_If you want to check your Bitdefender Push API configuration, you should run its executable passing it the getConfig parameter._

```
.\bdgzpush_conf.exe getConfig
```

## Running tests

_Then, to send a test log to Bitdefender's Push API, you must run its executable passing it the logTest parameter._

```
.\bdgzpush_conf.exe logTest
```

_At this time, the test logs should be arriving at your server or log collector_

## Licence
_This project is under MIT license_
