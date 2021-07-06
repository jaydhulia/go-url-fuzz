# go-url-fuzz
Fuzz URLs in Go! Just a fun tool I am developing as a way learn Go and to fuzz URLs in Go. Find those hidden directories on websites that developers forgot to remove! 

Given an URL containing the word `FUZZ` and a wordlist, it fuzzes the URL, replacing `FUZZ` in the URL with words from the wordlist.

Note: Inspired by [wfuzz](https://github.com/xmendez/wfuzz), I wanted to develop a similar fuzzer in Go, while learning Go!

## Features
  - Fuzz GET requests for an URL using a wordlist
  - Colored response by response codes to make it easier to find meaningful responses!
  - Concurred requests made using fasthttp, resulting in quick fuzzing!
  - Only show requests that return certain response codes! (`-s / --show flag`)

## Install
Installing is really easy. First make sure your go environment is set up, including `$GOPATH` and `$GOPATH/bin`. Then, git clone this repo. 
- `cd go-url-fuzz`
- `go install`
- `go-url-fuzz --help`

## Usage
```
go-url-fuzz [flags]
```

### Options

```
  -c, --concurrency int   path to a file containing wordlist to fuzz through (default 100)
  -h, --help              help for go-url-fuzz
  -s, --show ints         example: show 200,204,300
  -X, --type string       request type (currently only GET supported) (default "GET")
  -u, --url string        URL to fuzz, for example: "https://example.com/FUZZ" 
  -w, --wordlist string   path to a file containing wordlist to fuzz through
```

### Sample usage:
`go-url-fuzz -X=GET -url="https://www.google.com/?q=FUZZ" -wordlist=/path/to/wordlist.txt -c 150`


## Coming soon:
  - POST requests with data
  - More configuration options
  - Ability to have clearer output (only show if requests are a certain code)
  - Customizable headers, cookies
  
## Example Outputs
<table>
<tr>
  <th> <img src="https://github.com/jaydhulia/go-url-fuzz/blob/master/images/200.png" width="100%" height="400"> </th>
  <th> <img src="https://github.com/jaydhulia/go-url-fuzz/blob/master/images/404.png" width="100%" height="400"> </th>
</tr>
</table>
  
