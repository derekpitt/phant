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
2. Create method
3. Clear method
4. Delete method
5. Getting Data
