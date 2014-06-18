
###

```bash
$ heroku create -b https://github.com/kr/heroku-buildpack-go.git
Creating frozen-citadel-6244... done, stack is cedar
BUILDPACK_URL=https://github.com/kr/heroku-buildpack-go.git
http://frozen-citadel-6244.herokuapp.com/ | git@heroku.com:frozen-citadel-6244.git
Git remote heroku added

$ heroku addons:add heroku-postgresql
Adding heroku-postgresql on fierce-ocean-1703... done, v7 (free)
Attached as HEROKU_POSTGRESQL_BLACK_URL
Database has been created and is available
 ! This database is empty. If upgrading, you can transfer
 ! data from another database with pgbackups:restore.
Use `heroku addons:docs heroku-postgresql` to view documentation.

$ git push heroku master
Initializing repository, done.
Counting objects: 76, done.
Delta compression using up to 8 threads.
Compressing objects: 100% (64/64), done.
Writing objects: 100% (76/76), 57.95 KiB | 0 bytes/s, done.
Total 76 (delta 9), reused 0 (delta 0)

-----> Fetching custom git buildpack... done
-----> Go app detected
-----> Installing go1.2... done
-----> Running: godep go install -tags heroku ./...
-----> Discovering process types
       Procfile declares types -> web

-----> Compressing... done, 1.9MB
-----> Launching... done, v4
       http://frozen-citadel-6244.herokuapp.com/ deployed to Heroku

To git@heroku.com:frozen-citadel-6244.git
 * [new branch]      master -> master
```

pg_backend_pid comes from postgres
