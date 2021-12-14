# HttpRequestTester
A simple CLI tool for testing if a web address gives the expected response.

## How to use
To use this tool, download the exe and run it on a CLI. The following flags can be set:
- `httptype` _(REQUIRED)_ - The http method to be used. Current method types that are supported include: GET & POST
- `url` _(REQUIRED)_ - The URL to make the request to. eg- `"http://google.com"` NOTE: `http://` or `https://` must be included in the URL.
- `number` _(OPTIONAL)_ - The number of times to repeat the request. By default, there will only be 1 request made.
- `postbody` _(OPTIONAL)_ -  The data to be posted, if a POST request is made. By default the data that is sent is `{data: 'hello world!'}`- which uses JSON format
- `delay` _(OPTIONAL)_ - The delay between each request (if multiple requests are sent.) Represented in milliseconds. For example, if a delay of 5000 is given, there will be a 5 second delay between each HTTP request.
- `urls` _(OPTIONAL)_ - A path to a file to be used, if there are multiple HTTP requests to be made. URLs in this file, **must** be separated by a comma (eg - _,_). Appropriate file types - `.txt`, `.csv`
