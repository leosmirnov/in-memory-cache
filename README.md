# go-in-memory REST API cache

go-in-memory is an simple in-memory key:value store/cache similar to redis and memcached. Cache is able to store raw strings as a value and do some simple actions like adding, deleting, getting and viewing all the keys. All the data persists in memory. This example was made for demo usage. 

## deploy and configure instructions 

You can build docker image and run container by `./contrib/scripts/build_and_run_container.sh` from the project root. Also you are able to run bundle by `./contrib/scripts/build_and_run_bundle.sh`. 

You can configure project by editing file `./contrib/config.yml`. 

API documentation you are able to find in `./contrib/swagger/openapi.yaml`.

## usage

By default im-memory cache is deployed on :8080 port. Below you are able to see examples of requests using curl. 

To create new key, you should use: 


`curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"key":"key","value":"string", "expiration": 2}' \
  http://localhost:8080/values` 
It is worth mentioning that that `expiration` field represents expiration time in minutes. 


To get value by key, you should use: 
`curl http://localhost:8080/values/:key`

To get all the existing keys, you should use: 
`curl http://localhost:8080/keys`

To delete value by key, you should use: 
`curl --request DELETE \
  http://localhost:8080/values/:key`
