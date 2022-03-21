# Snippetbox

Go application to store quotes

## Starting/Ending db
```
/snippetbox$ docker-compose up -d database
/snippetbox$ docker-compose down
```

## Opening mySQL shell
```
/snippetbox$ docker-compose exec database mysql -u root -p
```

Important: Inside mysql shell you may need to use `use snippetbox` command
to select `snippetbox` database, before any operation.
