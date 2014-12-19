var AppCtrl = require('./app-controller');

describe('AppCtrl', function() {

    describe('isCurrentUrl', function() {
        var $scope,$location,ctrl;

        beforeEach(inject(function($controller, $rootScope) {
            $scope = $rootScope.$new();
            ctrl = new AppCtrl($scope);
        }));


        it('should pass a dummy test', function() {
            expect($scope.navVisible).toBeFalsy();
        });
    });
});
