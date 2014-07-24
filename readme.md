# phant go client for data.sparkfun.com

quick data poster for [http://data.sparkfun.com](http://data.sparkfun.com)

## example

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

# Docs

[https://godoc.org/github.com/derekpitt/phant](https://godoc.org/github.com/derekpitt/phant)

# TODO:

1. read X-Rate-Limit- headers
2. <del>CreateStream method</del>
3. <del>Clear method</del>
4. <del>Delete method</del>
5. <del>Getting Data</del>
6. <del>Stats</del>
