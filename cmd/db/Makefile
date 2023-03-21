.PHONY: static
SHELL := /bin/bash

DB_URL := mongodb+srv://sandbox:hzZQpPSUS6ROJd0L@cluster0.3h50i.mongodb.net/sandbox
COLL_NAME := usecase

db-backup:
	mongodump --uri $(DB_URL) --collection $(COLL_NAME)

db-restore:
	mongorestore --uri $(DB_URL) $(COLL_NAME).json
