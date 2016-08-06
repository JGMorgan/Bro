//Packages

var express = require('express');
var brApp = express();
var encApp = express();
var brPort = 8080;
var encPort = 4020;

//TODO: Routes:
//      -brApp:
//          -Login
//          -Registration
//          -GetFriendsList
//          -Any bro after first
//      -encApp
//          -firstBro
//              -Keys included

brApp.get('/', function(req, res) {
    res.send('Bro, welcome'); // Placeholder, not needed for a server
}); 



brApp.listen(brPort);
console.log('Brooooooo x' + brPort);
encApp.listen(encPort);
console.log('secret message on ' + encPort);
