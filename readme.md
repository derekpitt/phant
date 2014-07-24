# phant go client for data.sparkfun.com

quick data poster for [http://data.sparkfun.com](http://data.sparkfun.com)

## post data example

    package main

    import (
      "github.com/derekpitt/phant"
    )

    func main() {
      c := phant.Create("<publickey>", "<privatekey>")

      err := c.Post(map[string]string{"derek": "1", "test": "2"})

      if err != nil {
        // dang..
      }
    }

## get data example

    package main

    import (
      "fmt"
      "github.com/derekpitt/phant"
    )

    // define what fields are in the stream here
    type testType struct {
      Derek     string
      Test      string
      Timestamp string
    }

    func main() {
      // declare your data holder to be an array!
      var d []testType

      // don't forget to use your address operator here! (&)
      phant.AllData("RMMW22NWWzh6oj1NyADE", &d)

      fmt.Printf("%#v", d)
    }

# Docs

[https://godoc.org/github.com/derekpitt/phant](https://godoc.org/github.com/derekpitt/phant)

# TODO:

1. read X-Rate-Limit- headers
2. <del>CreateStream method</del>
3. <del>Clear method</del>
4. <del>Delete method</del>
5. <del>Getting Data</del>
6. <del>Stats</del>
