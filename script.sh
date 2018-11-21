heroku container:login
heroku create
heroku config:set MONGO_USER=
heroku config:set MONGO_PASS=
heroku config:set MONGO_HOST=
heroku container:push web
heroku container:release web
heroku open
