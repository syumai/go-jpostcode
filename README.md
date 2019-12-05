# go-jpostcode

* go-jpostcode is a Go package to find Japanese address data from Japanese postal code.
  - This package was created to provide data from https://github.com/kufu/jpostcode-data

## Usage

```go
// Find an address
address, err := jpostcode.Find("0010928")

// Search addresses (Some addresses have same postal code)
addresses, err := jpostcode.Search("1138654")

// Print address as a JSON
addressJSON, err := address.ToJSON()
if err != nil { // error handling }
fmt.Println(addressJSON)
```

### Example

* [HTTP server example](https://github.com/syumai/go-jpostcode/blob/master/example/server/main.go)

```console
$ go run example/server/main.go
$ curl localhost:8080/0010928
```

## Install a CLI tool to get address from postcode

* A CLI tool is given as [jpost](https://github.com/syumai/go-jpostcode/blob/master/cmd/jpost).

### Installation

```
go get -u github.com/syumai/go-jpostcode/cmd/jpost
```

### Usage of jpost

* To get address, **just give postal code** as argument.

```
# Get address from postal code: 0010928.
$ jpost 0010928
{"postcode":"0010928","prefecture":"北海道",...
```

## License

* MIT

## Author

* [syumai](https://github.com/syumai)

## Original data

* [jpostcode-data](https://github.com/kufu/jpostcode-data)
