# taxesGo

Get the basic sales tax

### Installation

taxesGo requires latest [Golang](https://golang.org/doc/install) to run.

### Run the app
Go to project folder and type:

```sh
$ cd taxesGo
$ go run main.go
```
Create the headers:
```sh
$ quantity, price, product
```
* This can be in any order
    * product, price, quantity
    * Price, Product, Quantity
* Switching from one order to another requires the app to be restarted

Next create the rows. Note that this needs to follow the order of the header. Press enter after each row.
```sh
$ 1, 3, discount iphone
$ 3, 20.65, expensive coffee
```

Lastly, press enter twice to see the results:
```sh
$ 1, 3.30, discount iphone
$ 3, 68.15, expensive coffee
$
$ Sales Taxes: 6.5
$ Total: 71.45
```

Our final entry should look like this:
```sh
$ quantity, price, product
$ 1, 3, discount iphone
$ 3, 20.65, expensive coffee
$
$ 1, 3.30, discount iphone
$ 3, 68.15, expensive coffee
$
$ Sales Taxes: 6.5
$ Total: 71.45
```

You can add more items at the end like this:
```sh
$ quantity, price, product
$ 1, 3, discount iphone
$ 3, 20.65, expensive coffee
$
$ 1, 3.30, discount iphone
$ 3, 68.15, expensive coffee
$
$ Sales Taxes: 6.5
$ Total: 71.45
$
$ 3, 10000, limited edition shoes
$ 1, 5000, air guitar
$
$ 3, 33000.00, limited edition shoes
$ 1, 550.00, air guitar
$
$ Sales Taxes: 3050
$ Total: 33550
```
### Tests
Run the tests:
```sh
$ go test -cover ./... -v
```
### Documentation
Run the docs:
```sh
$ godoc -http=":6060"
```
Then visit: [http://localhost:6060/pkg/github.com/clybs/](http://localhost:6060/pkg/github.com/clybs/)
