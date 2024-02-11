# Get company information

### 1 
user send message into websocket:

```json
{"ticker":"IBM","action":"search"}
```

### 2 
`pkg/router` - is the main and single place where application consume new
messages from frontend. It consume this message and "decide" what to do with it.
It it contain `"action": "search"` that mean that user want to find information
about company by `"ticker"` key. And send message to RabbitMQ that this user 
want information about this company. User and client id stored in message
headers.

### 3.a
Than `pkg/datasource` catch this message. `pkg/datasource` is the one of the 
places that connect with providers API's where we get company data. 
Currently it's GuruFocus ( Alphavantage is outdated since it's to expensive ).

`pkg/datasource` contain list of provider methods that require to call in order to 
get all values for report. It walk through this list by goroutines and send 
results in RabbitMQ. Simultaneously `pkg/datasource` store results in redis in order
to reduce API calls in data providers. If we have some data about some company 
in redis - we do not need to get same data from providers. Every key in Redis
will store for 30 min ( it can be changed by settings ).

### 3.b
It's not implemented yet but there should be the module that consume messages about
user requests. If user send message:
```json
{"ticker":"IBM","action":"search"}
```
We should save this information to db, add unique request identifier, and 
created date.

### 4
`pkg/datasource` send data from providers in exchange called `ex.client.<client_id>`.
Since user can be connected from multiple devices we need to differentiate them
to send data only to the device that ask this data. Without it - we will send 
to all client devices. Other events like `notifications` should be send for all
devices...

Also user request_id should be updated with `status="success"`. If some error 
happend during generating report - status should be `"failed"`

### 5
Client get this data, build overview page. This page contain button called - 
"Generate pdf Report". By clicking this button client should get pdf version of 
this report. To do this, frontend websocket message like this:
```json
{"id": "<uuid request_id>", "action": "generate_report"}
```

>Why not ```{"ticker":"IBM","action":"generate_report"}```? Because we need to 
>understand what client, when and what report he want to generate. Using just 
>`ticker` we lose all this information.

### 6
`pkg/router` again catch message from client and "understand" that user want
to generate PDF version of his request. And send RabbitMQ message to the 
`pkg/report`.

### 7
`pkg/report` - is the module responsible for generating pdf reports. It consume
messages from RabbitMQ and "see" that this user want to generate data for this
report. Than `pkg/report` get this data from redis using request_id. If key is 
not exists ( because redis will store it only for limited amount of time ) it 
go to database, ask for ticker name, and ask this value from `pkg/datasource`.

>I do not think this scenario with db will be very often. Usualy users should 
>generate pdf after a couple of minutes exploring data and deciding is it worth 
>to generate pdf version. They will not wait an hour without reloading the page
>and after that ask for pdf report... It it can happend, that's I we have this
>scenario.

### 8
Finaly `pkg/report` get all required data and generate pdf file. Store it in 
( currently non existed ) CDN and send to client message like this:
```json
{"id": "<uuid request_id>", "link": "<pdf file cdn link>"}
```

### 9
Client consume that message and display ( or replace "Generate pdf Report" btn )
with new button with link for pdf file. 

Instead of one "Generate pdf Report" btn we can create 2

- "Generate pdf Report" to download from browser.
- "Send via Email" - to send that file via email. This step will follow all 

required steps in order to generate file, but instead of returning link into 
the browser - it will send message that email with generated report has been
send to email. In this case websocket message should contain additional key
with "recepient email" like this:

```json
{
    "id": "<uuid request_id>",
    "action": "generate_report",
    "email_recepient": "<email address of the report recepient>"
}
```

if `"email_recepient"` was not provided - client email will be used.