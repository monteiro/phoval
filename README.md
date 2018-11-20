# Phone validation API using Go

Still in early stage.

## Motivation

Every business has the need to verify phone numbers in order to guarantee that they exist or in case we need to contact a specific customer or send a marketing SMS.
The goal with this tool is to allow all businesses to deploy this solution in-house.

## API

### Create a new phone number verification

POST `/phone/verification&phone_number=919999999&country_code=351`

#### Responses

`204`: Verification was created with success. `verification_id` is in the header.
`400`: There is validation error with the arguments that were passed.

### Phone number validation

PUT `/phone/verification&phone_number=919999999&country_code=351&code=768782`

`204`: Verification was validated with success
`409`: Verification does not exist or Verification exists and it was already validated
