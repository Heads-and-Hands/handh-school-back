  
  
  ````
  backend:
    image: registry./handh-school/back:latest
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
      