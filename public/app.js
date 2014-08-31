(function(){
    "use strict";

    var roverApp = angular.module('roverApp', [])

    roverApp.factory('jsonWebSocket', function(){

        var _connections = {};

        var JSONWebSocket = function(name,URI) {

            this.registeredEventListeners = {
                "onopen":[],
                "onerror":[],
                "onmessage":[]
            };

            this.name = name;
            this.connection = new WebSocket(URI);

            // When the connection is open, send some data to the server
            this.connection.onopen = function (event) {
                for(var c in this.registeredEventListeners.onopen) {
                    this.registeredEventListeners.onopen[c].call(this);
                }
            }.bind(this);

            // Log errors
            this.connection.onerror = function (error) {
                for(var c in this.registeredEventListeners.onerror) {
                    this.registeredEventListeners.onerror[c].call(this,error);
                }
            }.bind(this);

            // Log messages from the server
            this.connection.onmessage = function (e) {
                for(var c in this.registeredEventListeners.onmessage) {
                    this.registeredEventListeners.onmessage[c].call(this,JSON.parse(e.data));
                }

            }.bind(this);
        }

        JSONWebSocket.prototype.send = function(pojo) {
            this.connection.send(JSON.stringify(pojo));
        }

        JSONWebSocket.prototype.addEventListener = function(event,callback) {
            this.registeredEventListeners[event].push(callback);
        }

        return {
            create: function(name,URI) {

                if(!(name in _connections)) {

                    _connections[name] = new JSONWebSocket(name,URI);

                }

                return _connections[name];
            },

            get: function(name) {
                return _connections[name];
            }
        }

    });

    roverApp.controller('MainCtrl', function ($scope,jsonWebSocket) {

        var roverJSONWebSocket = jsonWebSocket.get("rover");

        roverJSONWebSocket.addEventListener("onopen",function() {
            console.log("Open");
            roverJSONWebSocket.send({type: "location",data: {lon: 1.23, lat: -1.05}}); // Send the message to the server
        });

        roverJSONWebSocket.addEventListener("onerror",function(error) {
            console.log("Error",error);
        });

        roverJSONWebSocket.addEventListener("onmessage",function(data) {
            console.log("Message", data);
        });

    });

    roverApp.run(function(jsonWebSocket) {
        var connection = jsonWebSocket.create('rover','ws://' + location.host + '/api/ws');
    })

})();
