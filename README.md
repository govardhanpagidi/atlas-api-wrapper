# atlas-api-helper
## Provides REST apis for atlas resources

### Req. Flow 
    Client ---------> atlas-api-helper -----> atlas-go-sdk---> ATLAS CLOUD/INFRA

#####        1. APIs uses atlas-go-sdk to make api calls atlas cloud and provision resources.
#####        2. There is a middleware which authenticates API requests expecting ATLAS API keys.


### Major libraries used:
    1.gorilla-MUX (gorilla/mux)
    2.atlas-go-sdk (mongodbatlas)
