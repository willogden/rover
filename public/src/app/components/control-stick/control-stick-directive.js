module.exports = function() {
    return {
        templateUrl: 'components/control-stick/control-stick.tpl.html',

        scope: {
            controlStickValue: '=ngModel'
        },

        link: function(scope, element, attrs) {
            scope.controlStickValue = 0;
            element.addClass("components__control-stick");
        }
    };
}
