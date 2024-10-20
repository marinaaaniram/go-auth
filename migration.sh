#!/bin/bash
source .env

goose -dir "${MIGRATION_DIR}" postgres "${MIGRATION_DSN}" up -v