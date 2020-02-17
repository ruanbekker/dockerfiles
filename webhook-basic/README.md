Server:

```
> docker run -it -p 5000:5000 ruanbekker/python-flask-webhook
```

Client:

```
> curl -XPOST http://localhost:5000 -d "foobar" -d "name" -d "foo"
event: {'event': b'foobar&name&foo'}
```

```
> curl -H "Content-Type: application/json" -XPOST http://localhost:5000 -d '{"name": "james"}'
event: {'event': {'name': 'james'}}
```
