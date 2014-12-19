module.exports = function RoverJSONWebSocketService(jsonWebSocketFactory) {
    return jsonWebSocketFactory.create('rover','ws://' + location.host + '/api/ws');
}
