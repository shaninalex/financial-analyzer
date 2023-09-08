### Channels and Queues

### consumers
- `isReportReady` => Map to websocket for frontend. To disable static timeout on frontend.
- `saveReportInfo` => just save report information

### Providers
- `p.report` - send report name, user information, ticker etc, creation time, report file info

### Queue
- `q.document` - queue holds evens about new reports

Communication between components on report generating.
```
p.report (new event)
    ||
    ||
    \/
q.document
    |=====> isPDFReady
    |=====> saveReportInfo
```

