'use strict';

var HomeCtrl = require("./home-controller");
var ControlStickDirective = require("../components/control-stick/control-stick-directive");

module.exports = angular.module('app.home', [
    //'home/home.tpl.html',
    //'home/home-subnav/home-subnav.tpl.html',
    //'components/control-stick/control-stick.tpl.html',
    'ui.router',
    'ngTouch'
])
.config(function config($stateProvider) {

    $stateProvider.state('home', {
        url: '/home',
        views: {
            "main": {
                controller: 'HomeCtrl',
                templateUrl: 'home/home.tpl.html'
            },
        },
        data: {
            pageTitle: 'Home'
        }
    });

})
.controller('HomeCtrl', HomeCtrl)
.directive('controlStick', ControlStickDirective);
