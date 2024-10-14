# Cadet statistics

Service for obtaining statistics of cadets' academic performance.
Getting data as json or xlsx-file

### Run service (Pipeline)
Service will be built and started automatically using **.gitlab-ci.yml** file


### Run service mannually (REQ_LOG - optionally)
It is necessary to set the exported variables to the correct data: set connection to database, session service, key for connection to graphql
```bash
export STATS_DATABASE_URL="postgres://login:password@192.168.100.5:5432/session_manager?search_path=academ_stats&sslmode=disable&pool_max_conns=20"
export ZERO_TOKEN="token_for_zero_one_graphql"
export EXCEL_DATABASE_URL="postgres://login:password@192.168.100.5:5432/excel_table?search_path=excel_table&sslmode=disable&pool_max_conns=20"
export EXCEL_SRV_ADDR="excel_table_container"
export EXCEL_SRV_PORT="4444"
export SESSION_SRV_ADDR="192.168.100.5"
export SESSION_SRV_PORT="9191"
export REQ_LOG="false"
```
after that you can start the service
```bash
make run
```

### APIs
#### Get statistic in json format
```http
GET http://localhost:9595/api/academ-stats/journey/top-cadets?id=66
```

#### Get statistic in excel format
```http
GET http://localhost:9595/api/academ-stats/journey/top-cadets-file?id=66
```

#### Get module list
```http
GET http://localhost:9595/api/academ-stats/module-list
```