#!/usr/bin/env bash

sudo -S service postgresql start
cd $HOME/bootdev/bloggator/sql/schema
goose postgres "postgres://postgres:postgres@localhost:5432/bloggator" up
cd -
