#!/bin/bash

echo "restarting..."
docker-compose restart web
docker-compose logs -f web
