# framework-microwave
A backend* web framework for people with different tastes  

## General technology.
- Server lang: golang
- Database: mongodb
- queries: graphql (kinda)

## A very short explination what my idea is
- Bind a database collection to webserver routes.  
- Use graphql to filter/query inside the database and change the data that will be send back.  
- DB collection settings/skeletons are defined by yaml files  
- Events for everything  
- Being a good api
- Make building complex websites easy

## Why i want to make this
There is not realy a all in one backend framework like laravel/buffalo that combines golang, mongodb and graphql without the serverless, google, amazon, and so on  

## Why mongodb
Because i'm not a fan of sql and i also don't like cloud providers like google's firebase  

## Why golang
I like go... Not much else :)
Okey maybe one thing types are handy

## kinda graphql???
I like the way graphql works on the client side but i don't like the way it works on the server.  
Graphql gets complicated when you have so much data that you would need an extra filter to the database like only filter out 1000 items and then run the Graphql query over that data. When the data that you want is outside that range you need to re-run the query on a list of extra items.  
Because of that reason i want to inplement the filtering from graphql directly on the database that would make it so much better.  
After the data is reciefed from the data filter out the un needed object items.

## Docs
More detailed docs can be found in the docs folder: https://github.com/mjarkk/framework-microwave/tree/master/docs

## Some of my ideas for this backend frameworks
NOTE: take this as just an idea everthing can still be changed
See: https://github.com/mjarkk/ideas/blob/master/server-framework/README.md
