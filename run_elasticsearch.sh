#!/bin/bash

docker run --name es01 -p 9200:9200 -it -m 1GB docker.elastic.co/elasticsearch/elasticsearch:8.17.4
