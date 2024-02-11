# Generate pdf

>This document is a copy of [issue](https://github.com/shaninalex/financial-analyzer/issues/53)
>And describe how report process is ( or will be ) implemented.

## Description

Creating report has multiple stages and we can't just increment user report count 
just by 1 request. The stages are:

- check if user able to request report
- collect data from api's
- calculate custom values
- show data on frontend

For example if for some reason we do not deliver requested data ( or some part 
of it ) - it should not be count as complete report and reports count should not 
be incremented. Instead we update this request_id in database with `issue` message.
That will help us to investigate what happend in report process.

So only if we deliver to client ALL requested data - we can increment user report 
count.

We can use "flow" approach [like in kratos](https://www.ory.sh/docs/kratos/self-service). 
Report flow should be initialized when user submit "ticker" form. Next this flow 
will be updated by report id on every step. The flow can be like this:

- initialize report
- get data from providerA.dataPart1 => save in redis
- get data from providerA.dataPart2 => save in redis
- get data from providerB.dataPart1 => save in redis
- get data from providerB.dataPart2 => save in redis
- ---- // ----
- calculate custom values => save in redis
- ensure frontend receive all data
- report is done

## Implementation

Every time user request company details - we save request in database with unique 
identifier. This id is spreaded for all modules that taking action in generating
report process. Table should looks like this:

```
id: UUID
user_id: UUID
ticker: string
link: string (optional)
date_created: string
success: bool
issue: ForeignKey for issues.id ( optional )
```

Here `issue` is the foreign key for another table where I log all problems app
facced with. If something goes wrong with generating report process - we will 
save error in this table. This event can be shown to client and on the backoffice
admin panel ( currently not implemented ).
