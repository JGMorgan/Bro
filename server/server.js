//Routes
var express = require('express'),
  brApp = express(),
  encApp = express(),
  fs = require('fs'),
  config = require('./config.js');

//WebSocket cut and paste code
var brWebSocket = require('websocket').server;
var http = require('http');

var server = http.createServer(function(req, res) {
   console.log((new Date()) + ' Received request for ' + req.url);
   res.writeHead(404);
   res.end();
});
server.listen(config.brPort, function() {
    console.log((new Date()) + ' - Brooooo x' + config.brPort);
});


wsServer = new brWebSocket({
    httpServer: server,
    autoAcceptConnections: false
});

function originIsAllowed(origin) {
    return true;
}

wsServer.on('request', function(request) {
    if(!originIsAllowed(request.origin)) {
        request.reject();
        console.log((new Date()) + ' Connection from origin ' + request.origin + ' rejected.');
        return;
    }

    var connection = request.accept('echo-protocol', request.origin);
    console.log((new Date()) + ' Connection accepted.');
    connection.on('message', function(message) {
        if(message.type === 'utf8') {
            console.log('Received Message: ' + message.utf8Data);
            connection.sendUTF(message.utf8Data);
        }
        else  { console.log('Bad Data Input') }
    });
    connection.on('close', function(reasonCode, description) {
        console.log((new Date()) + ' Peer ' + connection.remoteAddress + ' disconnected.' );
    });
});



//Database
var mongoose = require('mongoose');
//mongoose.connect('mongodb://localhost/');
//var db = mongoose.connection;


//TODO: Routes:
//      -brApp:
//          -Any bro after first
//      -encApp
//          -firstBro
//              -Keys included
//          -GetFriendsList

/*
 *
 * brApp Section
 *
*/

//Placeholder for testing - might be route for websocket
brApp.get('/', function(req, res) {
    res.send('Bro, welcome'); // This is not needed, used for testing
});

// Websocket, not complete


/*
 *
 * encApp Section
 *
*/

//Placeholder for testing - might be route for websocket
encApp.get('/', function(req, res) {
    res.send('shhh'); // This is not needed, used for testing
});

// TODO Empty Function
encApp.post('/registration', function(res, req, next) {

});

// TODO Empty Function
encApp.post('/login', function() {

});

// TODO Empty Function
encApp.post('/addFriend', function() {

});

// TODO Empty Function
encApp.post('/removeFriend', function() {

});

// TODO Empty Function
encApp.get('/getFriendsList', function() {

});


/*
 *
 * Server Setup and Starting
 *
*/

//brApp.listen(brPort);
//console.log('Brooooooo x' + brPort);
encApp.listen(config.encPort);
console.log('secret message on ' + config.encPort);
