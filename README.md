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

## Inserting dummy data

### Creating a new UTF-8 `snippetbox` database.
```
CREATE DATABASE snippetbox CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

### Switch to using the `snippetbox` database.
```
USE snippetbox;
```

### Create a `snippets` table.
```
CREATE TABLE snippets (
  id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
  title VARCHAR(100) NOT NULL,
  content TEXT NOT NULL,
  created DATETIME NOT NULL,
  expires DATETIME NOT NULL
);
```

### Add index
```
CREATE INDEX idx_snippets_created ON snippets(created);
```

### Adding dummy data
```
INSERT INTO snippets (title, content, created, expires) VALUES (
  'An old silent pond',
  'An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.',
  UTC_TIMESTAMP(),
  DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);

INSERT INTO snippets (title, content, created, expires) VALUES (
  'Over the wintry forest',
  'Over the wintry\nforest, winds howl in rage\nwith no leaves to blow.',
  UTC_TIMESTAMP(),
  DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);

INSERT INTO snippets (title, content, created, expires) VALUES (
  'First autumn morning',
  'First autumn morning\nthe mirror I stare into\nshows my father''s face.\n\
   UTC_TIMESTAMP(),
  DATE_ADD(UTC_TIMESTAMP(), INTERVAL 7 DAY)
);
```

### Grant permissions to user
```
GRANT SELECT, INSERT ON snippetbox.* TO 'user';
```

Important: Inside mysql shell you may need to use `use snippetbox` command
to select `snippetbox` database, before any operation.
