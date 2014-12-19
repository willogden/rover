module.exports = function($scope) {

    $scope.$on('$stateChangeSuccess', function(event, toState, toParams, fromState, fromParams) {
        if (angular.isDefined(toState.data.pageTitle)) {
            $scope.pageTitle = toState.data.pageTitle + ' | app';
        }
    });

    $scope.navVisible = false;

    $scope.toggleNav = function() {

        $scope.navVisible = !$scope.navVisible;

    };

};
