NotFound
=========
[![](https://images.microbadger.com/badges/image/ushios/not-found.svg)](https://microbadger.com/images/ushios/not-found "Get your own image badge on microbadger.com")
[![](https://images.microbadger.com/badges/version/ushios/not-found.svg)](https://microbadger.com/images/ushios/not-found "Get your own version badge on microbadger.com")

Return 404 status everywhere.

### Run on docker

```bash
$ docker run -d -p 80:8080 ushios/not-found
$ curl localhost:8080
404 page not found
$ curl localhost:8080/hc
OK
```
