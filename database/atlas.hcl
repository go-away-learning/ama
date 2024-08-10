locals {
    db_user     = getenv("DATABASE_USER")
    db_password = urlescape(getenv("DATABASE_PASSWORD"))
    db_host     = getenv("DATABASE_HOST")
    db_port     = getenv("DATABASE_PORT")
    db_name     = getenv("DATABASE_NAME")
}

env "local" {
    src = "database/schema.sql"
    url = "postgres://${local.db_user}:${local.db_password}@${local.db_host}:${local.db_port}/${local.db_name}?search_path=public&sslmode=disable"
    dev = "docker://postgres?search_path=public"
}
