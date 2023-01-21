# REGISTERSERVICE

About to listen on Port: 8442.

SUPPORTED REQUESTS:

GET:

Create Error on: http://127.0.0.1:8442/error

Goal is to shutdown the server.

POST:

Register on: http://127.0.0.1:8442/register requires a JSON Body with the following format:

    {
    "Username": "mmuster",
    "Password": "password",
    "Firstname": "Max",
    "Lastname": "Muster",
    "Housenumber": "1",
    "Street": "Musterstr.",
    "Zipcode": "01234",
    "City": "Musterstadt",
    "Email": "max.muster@mail.com",
    "Phone": "012345678910"
    }
