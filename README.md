# Dynamic UI Server
Dynamic UI server is the server for Dynamic UI configuration
CMS can be set up using : https://github.com/aviralbansal29/duis_cms

### Service Description
- `Variant` defines what variant user belongs to. For example, if this is attached to investments app, variant can be based on user activity like how much risk user takes or for how long user has made investments.
- `Schema` stores json of data to return to user. For example if schema is of type home_page, we define what all components to be visible, like in case of investment app, we can say, first display current investments, then recommended investments, then current trends and so on
- `User` table basically stores user info. If this is deployed as separate microservice, there will be no user table in this service.

### API Doc:
- Start the server
- Go to : http://0.0.0.0:9000/swagger/index.html#/

### Creating Environment
- Update env as needed in env.yaml
- Environment variables can be set to overwrite env.yaml, for eg. `aws.client_id` will be overwritten by environment variable `AWS_CLIENT`
- (Required before postgres volume is created) Create empty file `.env.docker` to store variables to be kept in secret manager

## Server:
- Build server  : `docker-compose build`
- Run server    : `docker-compose up -d`
- Live Logs     : `docker logs duis_app_1 -f`
- Stop server   : `docker-compose down`

## Database config
- Update postgres env varaibles inside .env.docker
- Varaibles to Set : 
  - POSTGRES_DB=<DB name>
  - POSTGRES_PORT=<DB Port>
  - POSTGRES_USER=<DB default username>
  - POSTGRES_PASSWORD=<Default password>
  - POSTGRES_HOST=database

