// call the packages we need, Load all of our middleware
var express       = require('express');        // call express
var app           = express();                 // using express
var bodyParser    = require('body-parser');
var http          = require('http')
var fs            = require('fs');
var Fabric_Client = require('fabric-client');
var path          = require('path');
var util          = require('util');
var os            = require('os');
// configure app to use bodyParser(),get the data from a POST
// app.use(express.static(__dirname + '/client'));
app.use(bodyParser.urlencoded({ extended: true }));
app.use(bodyParser.json());
// instantiate the app
var app = express();
// runs routes.js file and passes it app
require('./routes.js')(app);
// set up a static file server to the "Index" directory
app.use(express.static(path.join(__dirname, './Index')));
//http.createServer(function(request,response){}).listen(8000);
// Save our port
var port = process.env.PORT || 8000;
// Start the server and listen on port
app.listen(port,function(){
  console.log("Live on port: " + port);
});
