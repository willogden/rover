module.exports = function HomeController($scope,roverJSONWebSocketService) {

    var roverJSONWebSocket = roverJSONWebSocketService;

    roverJSONWebSocket.addEventListener("onopen",function() {
        //console.log("Open");
        //roverJSONWebSocket.send({type: "location",data: {lon: 1.23, lat: -1.05}}); // Send the message to the server
        //roverJSONWebSocket.send({type: "motorspeed",data: {motor: 0, speed: 100}}); // Send the message to the server
    });

    roverJSONWebSocket.addEventListener("onerror",function(error) {
        console.log("Error",error);
    });

    roverJSONWebSocket.addEventListener("onmessage",function(data) {
        console.log("Message", data);
    });

    $scope.motorspeed0 = 0;
    $scope.motorspeed1 = 0;

    $scope.$watch('motorspeed0', function(newValue, oldValue) {
        if(roverJSONWebSocket.isReady()) {
            roverJSONWebSocket.send({type: "motorspeed",data: {motor: 0, speed: Number(newValue)}})
        }
    });

    $scope.$watch('motorspeed1', function(newValue, oldValue) {
        if(roverJSONWebSocket.isReady()) {
            roverJSONWebSocket.send({type: "motorspeed",data: {motor: 1, speed: Number(newValue)}})
        }
    });

};
