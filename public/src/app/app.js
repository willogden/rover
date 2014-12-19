'use strict';

require('angular');
require('angular-touch');
require('angular-ui-router');
require('famous-angular');
require('../../build/templates');

module.exports = angular.module('app', [
    'ui.router',
    'famous.angular',
    // Load the partials
    'templates',

    // Add modules/sections as dependencies
    require('./home/home').name,
    require('./about/about').name
])
.config(function myAppConfig($stateProvider, $urlRouterProvider) {
    $urlRouterProvider.otherwise('/home');
})
.controller('AppCtrl',require('./app-controller'))
.factory('jsonWebSocketFactory', require('./components/jsonwebsocket-factory'))
.service('roverJSONWebSocketService', require('./components/rover-jsonwebsocket-service'))
.run(['$rootScope', '$state', '$stateParams','jsonWebSocketFactory', function ($rootScope,   $state,   $stateParams, jsonWebSocketFactory) {
    $rootScope.$state = $state;
    $rootScope.$stateParams = $stateParams;
}]);
