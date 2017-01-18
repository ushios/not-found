NotFound
=========
Return 404 status everywhere.

### Run on docker

```bash
$ docker run -d -p 80:8080 ushios/not-found
$ curl localhost:8080
404 page not found
$ curl localhost:8080/hc
OK
```
