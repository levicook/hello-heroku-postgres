
### GET /
```
BUILDPACK_URL=https://github.com/kr/heroku-buildpack-go.git
DYNO=web.1
PATH=/usr/local/bin:/usr/bin:/bin:/app/bin
PWD=/app
PS1=\[\033[01;34m\]\w\[\033[00m\] \[\033[01;32m\]$ \[\033[00m\]
SHLVL=1
HOME=/app
PORT=54407
_=/app/bin/hello-heroku-db
```

###

```bash
$ heroku addons:add heroku-postgresql
Adding heroku-postgresql on fierce-ocean-1703... done, v7 (free)
Attached as HEROKU_POSTGRESQL_BLACK_URL
Database has been created and is available
 ! This database is empty. If upgrading, you can transfer
 ! data from another database with pgbackups:restore.
Use `heroku addons:docs heroku-postgresql` to view documentation.
```

### GET /

```
BUILDPACK_URL=https://github.com/kr/heroku-buildpack-go.git
HEROKU_POSTGRESQL_ONYX_URL=postgres://qbgtcjqdvuhdlc:2TIRMmUbqlu739_Wz8yxoxOCyE@ec2-54-225-101-199.compute-1.amazonaws.com:5432/deb8jkskas8vj5
DYNO=web.1
PATH=/usr/local/bin:/usr/bin:/bin:/app/bin
PWD=/app
PS1=\[\033[01;34m\]\w\[\033[00m\] \[\033[01;32m\]$ \[\033[00m\]
SHLVL=1
HOME=/app
PORT=42813
DATABASE_URL=postgres://qbgtcjqdvuhdlc:2TIRMmUbqlu739_Wz8yxoxOCyE@ec2-54-225-101-199.compute-1.amazonaws.com:5432/deb8jkskas8vj5
_=/app/bin/hello-heroku-db
```
