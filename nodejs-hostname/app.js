var http = require('http');
var os = require('os');

http.createServer(function (req, res) {
    res.writeHead(200, {'Content-Type': 'text/html'});
    res.end(`My Hostname: ${os.hostname()}\n`);
}).listen(8080);

console.log('Listening on :8080');
