// LightGetCtrl
app.controller("LightGetCtrl", ['$scope', 'light', function($scope, light) {
  // Get the light
  $scope.light = light;
}]);