# go-url-fuzz
Fuzz URLs in Go! Just a fun tool I am developing as a way learn Go and to fuzz URLs in Go. Find those hidden directories on websites that developers forgot to remove! 

Note: Inspired by [wfuzz](https://github.com/xmendez/wfuzz), I wanted to develop a similar fuzzer in Go, while learning Go!

## Features
  - Fuzz GET requests for an URL using a wordlist
  - Colored response by response codes to make it easier to find meaningful responses!
  - Concurred requests made using fasthttp, resulting in quick fuzzing!

## Install
Installing is really easy. First make sure your go environment is set up, including `$GOPATH` and `$GOPATH/bin`. Then, git clone this repo. 
- `cd go-url-fuzz`
- `go install`
- `go-url-fuzz --help`

## Usage
### Flags:
  - `-X="GET"` string (default "GET") - Request type
  - `-url="https://example.com/FUZZ"` string *Required* - The URL to fuzz, note that the word `FUZZ` will be replaced in the URL by words from th wordlist for each request
  - `-wordlist=/PATH/TO/WORDLIST/FILE` string *Required* - The wordlist to use
  - `-c=100` int (default 100) - Number of concurrent requests to make, defaults to 100. Must be a number between 1 and 200.

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
  
