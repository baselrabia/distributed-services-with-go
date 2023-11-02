#  Let's Go (Http Server)

### Test Your API
You now have a functioning JSON/HTTP commit log service you can run and
test by hitting the endpoints with curl. Run the following snippet to start the server: </br>
```
$ go run main.go
```
Open another tab in your terminal and run the following commands to add
some records to your log:
```
$ curl -X POST localhost:8080 -d \
'{"record": {"value": "TGV0J3MgR28gIzEK"}}'
$ curl -X POST localhost:8080 -d \
'{"record": {"value": "TGV0J3MgR28gIzIK"}}'
$ curl -X POST localhost:8080 -d \
'{"record": {"value": "TGV0J3MgR28gIzMK"}}'
```
Go’s encoding/json package encodes []byte as a base64-encoding string. The
record’s value is a []byte, so that’s why our requests have the base64 encoded
forms of Let’s Go #1–3. You can read the records back by running the following
commands and verifying that you get the associated records back from the
server:
```
$ curl -X GET localhost:8080 -d '{"offset": 0}'
$ curl -X GET localhost:8080 -d '{"offset": 1}'
$ curl -X GET localhost:8080 -d '{"offset": 2}'
```
Congratulations—you have built a simple JSON/HTTP service and confirmed
it works!