module.exports = function() {

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

    JSONWebSocket.prototype.isReady = function() {
        return this.connection.readyState == 1;
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
}
