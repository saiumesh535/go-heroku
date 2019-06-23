# go-heroku
Deploying Go Application on Heroku

# Config files

* **heroku.yml** file Helps heroku to identify project Language (GoLang) In our example

* **Dockerfile** and **Makefile** once heroku identifies Application Language (GoLang) it runs **Makefile** which includes **Dockerfile** to build our application

* **Procfile** Once heroku builds our application, **Procfile** tells heroku to run which command to start application.

## Important

This repo Go project name is **go-heroku** which will be in several files, to rename project name, find and replace **go-heroku** with your desired name.

For more information please visit [here](https://devcenter.heroku.com/articles/getting-started-with-go)

**Deploy commands** [here](https://devcenter.heroku.com/articles/getting-started-with-go#deploy-the-app)

## Restart server on save.

We will be using [fswatch](https://github.com/codeskyblue/fswatch) awesome library for restarting our server on save.

```cmd
fswatch
```
above command will start and restart the server on every save.