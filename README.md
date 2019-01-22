# ports

## Running services: 
```
docker-compose build
docker-compose up
```

## APIService endpoints:
### Searching data 
`GET: http://localhost:8083/api/v1/port?portId={KEY}`


### Sending json file to save data: 
`POST: http://localhost:8083/api/v1/ports` \
Add json file to `form-data` with key `file`
