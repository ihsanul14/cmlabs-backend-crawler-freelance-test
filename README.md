# cmlabs-backend-crawler-freelance-test

This project goal is to handle web crawling with simple handler. 

The source code layout were separated into 4 group based on separation of concern. Those groups are:
- Framework 
  - This layer is consist of infrastucture of the project such as databases, logger, error, utils function, .etc

- Application 
  - This layer is focusing as presentation layer to the client such as http, gRPC, websocket, .etc. In terms of http all of the handler function will be define in this layer

- Usecase
  - This layer is focusing on business logic of the application like retrieve data, transform it, and manipulating the data output  

- Entity
  - This layer is consist of all data sources of the projects like query to the databases, fetching data from external resources through API, .etc. 


# How to run (Locally)

- clone the project and cd to the project
```
git clone https://github.com/ihsanul14/cmlabs-backend-crawler-freelance-test.git
cd /path/to/project
```

- install additional dependency. Follow this link :
  - chrome-driver: https://chromedriver.chromium.org/getting-started/

- running the application
```
go run .
```

- access the application to http://localhost:30001

# Run Unit Test

```
go test ./... --coverprofile=coverage.out
```

# Test the Application

- [API Documentation](./web-crawler.postman_collection.json)
- [HTML output](./framework/output/)