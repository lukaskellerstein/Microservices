var http = require('http');

var server = http.createServer(function (request, response) {

    if (request.url == '/') {

        response.writeHead(200, { "Content-Type": "text/html" });
        response.end("<strong>Hello</strong> world");

    } else if (request.url == '/end') {

        response.writeHead(200, { "Content-Type": "text/html" });
        response.end("<strong>Goodbye</strong> world");

    } else if (request.url == '/other') {

        if (request.method == 'GET') {

            response.writeHead(200, { "Content-Type": "text/html" });
            response.end("<strong>GET</strong> Other world");

        } else if (request.method == 'POST') {

            response.writeHead(200, { "Content-Type": "text/html" });
            response.end("<strong>POST</strong> Other world");

        }
    }

});

server.listen(10000, function () {
    console.log("server started on port 10000")
})