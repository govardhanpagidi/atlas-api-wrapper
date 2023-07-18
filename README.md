# mongodb atlas-api-helper
### Provides REST apis for atlas resources

### Req. Flow 
    Client ---------> atlas-api-helper -----> atlas-go-sdk---> MongoDB Atlas Cloud/INFRA

#####        1. APIs uses atlas-go-sdk to make api calls to mongodb atlas cloud and provision resources.
#####        2. There is a middleware which authenticates API requests expecting ATLAS API keys.


### Major libraries used:
    1.gorilla-MUX (gorilla/mux) for HTTP handlers and routes.
    2.atlas-go-sdk (mongodbatlas) for atlas api 


## How to test

#### API runs on port 8080
#### example:  
            curl --location --request GET 'localhost:8080/api/project?Id=64b6d746ec88e0087ddfg10a' \
            --header 'Content-Type: application/json' \
            --header 'Authorization: Basic Z2hld3ZuasdfZ3k6ZTA3MDJkNmItYjA2Mi00sdfasdYTcwLWJiZDAtNzA0NGM0ZjUwZjc1' \
            --data '{
                "Name": "$ProjectName",
                "OrgId": "$OrgID",
                "ProjectSettings": {
                    "IsCollectDatabaseSpecificsStatisticsEnabled": false,
                    "IsDataExplorerEnabled": false
                }
            }'
