## Messages types

Base message:
```json
{
    "action": "<notification | data_result>",
    "payload": "<IResponseData | INotification>"
}
```

---

Return from datasource to frontend (payload.IResponseData)
```json
{
    "action": "data_result",
    "payload": {
        "data": "any",
        "ticker": "string",
        "type": "string"
    }
}
```

payload.type defined here: [internal/typedefs/datasource.go](internal/typedefs/datasource.go)

---

Notifications (payload.INotification):
```json
{
    "action": "notification",
    "payload": {
        "level": "<success | error | warning | info>",
        "message": "<string>",
        "read": "<boolean>",
        "datetime": "<Date>"
    }
}
```
