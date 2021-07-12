  **Prepare**
  - Run ```bindatafs bindatafs```       
  - Make ```go run main.go``` without params with Compile() function after SetAssetFS() in myAdmin.go
  - ```go build -tags bindatafs main.go``` for build project
  - ```go run```

  **Build**
  ```bigquery
  docker buildx build --platform linux/amd64 -t registry.ru/handh-school/back-amd:latest --push .
  ```
  
  **Deploy**
  ```
  backend:
    image: registry.ru/handh-school/back:latest
    networks:
      - default
    ports:
      - "7773:7771"        
    environment:
      MYSQL_HOST: "database:3306"
      MYSQL_USER: "user"
      MYSQL_PASSWORD: "password"
      MYSQL_DATABASE: "handh-school"
      START_DELAY: 10s
      AUTH_USER: "admin"
      AUTH_PASSWORD: "password"
    depends_on:
      - database
  ```    
      
  **PS**

  After first deploy run ```ALTER TABLE requests CONVERT TO CHARACTER SET utf8;``` in database container for correct support russian symbols
