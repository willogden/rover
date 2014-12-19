(function(module) {
try {
  module = angular.module('templates');
} catch (e) {
  module = angular.module('templates', []);
}
module.run(['$templateCache', function($templateCache) {
  $templateCache.put('about/about.tpl.html',
    '<ul><li ng-repeat="choice in dropdownDemoItems"><a>{{choice}}</a></li></ul>');
}]);
})();

(function(module) {
try {
  module = angular.module('templates');
} catch (e) {
  module = angular.module('templates', []);
}
module.run(['$templateCache', function($templateCache) {
  $templateCache.put('home/home.tpl.html',
    '<div class="home" ng-controller="HomeCtrl"><control-stick ng-model="motorspeed0"></control-stick><control-stick ng-model="motorspeed1"></control-stick><fa-app style="height: 200px"><fa-surface fa-background-color="\'red\'">Hello world</fa-surface></fa-app></div>');
}]);
})();

(function(module) {
try {
  module = angular.module('templates');
} catch (e) {
  module = angular.module('templates', []);
}
module.run(['$templateCache', function($templateCache) {
  $templateCache.put('components/control-stick/control-stick.tpl.html',
    '<input ng-model="controlStickValue" type="range" min="0" max="100" value="0">');
}]);
})();
