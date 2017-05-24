# Welcome to go-api-test with Revel

[Revel](https://revel.github.io/) is a high-productivity web framework for the [Go language](http://www.golang.org/).


### Start the web server:
```sh
$ export $GOPATH=$REVEL_HOME
$ cd $GOPATH

$ revel run github.com/Junbong/go-api-test
```

### Go to http://localhost:9000/ and you'll see:

    "It works"

## Routes
### Shop Funnels
[http://localhost:9000/shops/:shop_id/funnels](http://localhost:9000/shops/:shop_id/funnels)

* URL params
  * `shop_id`: Single shop ID, comma separated multiple shop IDs, or tag
* Query params
  * `from`: Date formatted YYYY-MM-DD
  * `to`: Date formatted YYYY-MM-DD
  * `profile`: Comma separated multiple profiles

### Shop Funnels Average
[http://localhost:9000/shops/:shop_id/funnels/average](http://localhost:9000/shops/:shop_id/funnels/average)

* URL params
  * `shop_id`: Single shop ID, comma separated multiple shop IDs, or tag
* Query params
  * `from`: Date formatted YYYY-MM-DD
  * `to`: Date formatted YYYY-MM-DD
  * `profile`: Comma separated multiple profiles

## Code Layout

The directory structure of a generated Revel application:

    conf/             Configuration directory
        app.conf      Main app configuration file
        routes        Routes definition file

    app/              App sources
        init.go       Interceptor registration
        controllers/  App controllers go here
        views/        Templates directory

    messages/         Message files

    public/           Public static assets
        css/          CSS files
        js/           Javascript files
        images/       Image files

    tests/            Test suites


## Help

* The [Getting Started with Revel](http://revel.github.io/tutorial/gettingstarted.html).
* The [Revel guides](http://revel.github.io/manual/index.html).
* The [Revel sample apps](http://revel.github.io/examples/index.html).
* The [API documentation](https://godoc.org/github.com/revel/revel).

