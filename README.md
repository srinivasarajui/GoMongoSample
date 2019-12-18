# GoMongoSample

##Database usage
A docker compose file is provided to start mongo.

To start
```Console
docker-compose -f "docker-compose.yml" up -d  
```
To access docker container use the following command,  this will login.
```Console
docker exec -it testing-mongo-container bash
```

To access the mongo shell use the following when you are in docker console.

```Console
mongo -u admin1 -p password1 --authenticationDatabase testing
```


