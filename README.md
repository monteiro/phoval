# Phone validation API using Go

[![Build Status](https://travis-ci.org/monteiro/phoval.svg?branch=master)](https://travis-ci.org/monteiro/phoval)

## Motivation

Every business has the need to verify phone numbers in order to guarantee that they exist or in case we need to contact a specific customer or send a marketing SMS.
The goal with this tool is to allow all businesses to deploy this solution in-house.

## Install

Testing locally:

```
make docker-up 
make migrate
```

You can now call the API using curl or any other tool.

It will instantiate 2 containers (one _MySQL_ and one executing the binary with the tool in development mode).
In **Development mode** you can test the API and see the results in the output of `docker-compose`:

```
phovalapp_1_d00e1ca2feb7 | 2018/12/07 09:56:58 172.21.0.1:36848 - "HTTP/1.1 POST /phone/verification?phone_number=963695658&country_code=351"
phovalapp_1_d00e1ca2feb7 | 2018/12/07 09:56:58 SMS was sent: '{351 9611111111 This is your code: 799184
phovalapp_1_d00e1ca2feb7 |  phoval}'
``` 

## Run the http server on your machine or in production

Define the following environment variables for _AWS SES_ and _MySQL_ Database configuration:

```
export AWS_SDK_LOAD_CONFIG=1
export AWS_ACCESS_KEY_ID=XXXXXX
export AWS_SECRET_ACCESS_KEY=XXXXXX
export AWS_DEFAULT_REGION=XXXXXX
export DB_USER=XXXXXX
export DB_PASSWORD=XXXXXX
export DB_HOST=XXXXXX
export DB_PORT=3306
export DB_NAME=XXXXXX
```

```
make build
make migrate
./bin/phoval-{linux,mac} or ./bin/phoval-windows.exe {flags}
```

### Usage

- `addr`: Http network address and port to bind (e.g. 192.168.0.1:4000)
- `userdb`: database user
- `passworddb`: database user password
- `hostdb`: database host
- `namedb`: database name
- `env`: environment ("prod", "dev" or testing) - it will use a different SMS implementation according to the environment value.
- `brand`: brand name used in the SMS to specify the where it comes (e.g. phoval-brand)
- `apikey`: api key to protect the service. By default it's `changeme`. Better to change it in production.
- `template-folder`: template folder where are the SMS templates. Currently it's in `messages` folder.  

Example:

```
bin/phoval-linux -addr=phoval-app.com:4000 -userdb=myDbUser -passworddb=myDbPassword -hostdb=phovaldb.com -namedb=phoval -env=prod -brand=phoval -apikey=secret -template-folder=/usr/messages/
```

## API usage

Arguments:
- phone_number (string): phone number 
- country code (string): needs to exist (it does not include the 00 or + as prefix)
- locale (optional, string): locale used to fetch the right message translation inside `pkg/phoval/messages` (only en locale supported at this moment)

### Create a new phone number verification

POST `/phone/verification&phone_number=919999999&country_code=351&locale=en`

#### Responses

`204`: Verification was created with success. `verification_id` is in the header.

`400`: There is validation error with the arguments that were passed.

### Phone number validation

PUT `/phone/verification&phone_number=919999999&country_code=351&code=768782`

`204`: Verification was validated with success

`409`: Verification does not exist or Verification exists and it was already validated

### Contribution

How can you help and contribute to this tool:

- [Creating issues](https://github.com/monteiro/phoval/issues/new) with ideas
- Creating PRs of those ideas

#### Development environment: 

```
make docker-up
make migrate
make deps
make run
```
