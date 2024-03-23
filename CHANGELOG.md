## v1.7.3 (2024-03-23)

### Fix

- **report**: update report status if all required data has been gathered
- **types**: Make action type more flexible
- **types**: design more straight and simpler type for actions (messages) between services
- **report**: trigger gathering information and return update message
- **report**: change service structure, change mq exchange type
- **report**: make separate service

### Refactor

- **report**: default state in switch/case

## v1.7.2 (2024-03-17)

### Fix

- run updated architecture
- change ports env variables names
- update readme
- change application structure
- **redis**: get value from redis if exists and store in redis in debug mode
- **report**: change action types
- **router**: Changing action types #75

## v1.7.1 (2024-03-16)

### Fix

- **db**: Write main db methods and tests
- **db**: make database interface with postgres implementation
- **db**: use gorm instead of raw queries
- change application folder structure

## v1.7.0 (2024-03-02)

### Feat

- save datasource responses into redis store
- add redis pkg

### Fix

- remove alphavantage data provider
- **config**: use shell script and rabbitmqadmin command line to setup rabbitmq
- **mq**: unbind unused queues #66
- update golang version
- **db**: define simple reports table

### Refactor

- **db**: move into 'internal' folder
- **database**: user single db instead of couptle #64

## v1.6.0 (2024-02-11)

## v1.1.1 (2023-12-16)

## v1.0.0 (2023-12-16)

### Feat

- consume messages in report module #53
- add profile page #54
- add auth and logout link #54
- update templates #57
- update angular cli #57
- implement google signin #45
- check is account abble to make reports
- light/dark themes
- add financials chart
- add price chart
- remove alphavantage
- assembling all togather
- reorganizing application
- add new data to report request
- setting up frontend
- simplify frontend
- update signup ui + add new pages
- add auth service for root module
- add recovery component
- global configuration update
- remove useless services

### Fix

- updates
- **api**: change docker domain in oathkeeper config #53
- **kratos**: Remove separate kratos proxy service #53
- replace angular-highcharts with chart.js #58
- remove report id creatino in router module #53
- use combine reducer for dashboard
- add card component
- dashboard layout
- remove icons module
- redefine messages types
- refactor datasource
- udpate and fix consume data on frontend
- Configuration
- configuration
- move port variable to global declarations
- little tweaks
- dependencies
- release strategy

### Refactor

- providers
- start to organize everything in single source
- remove own css

## v1.5.0 (2023-11-18)

## v1.4.0 (2023-11-18)

## v1.3.0 (2023-11-16)

## v1.2.0 (2023-11-06)

## v1.1.0 (2023-11-02)

### Feat

- initial commit
