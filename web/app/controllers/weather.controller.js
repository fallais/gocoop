// WeatherGetCtrl
app.controller("WeatherGetCtrl", ['$scope', 'weather', 'sunrise', 'sunset', function($scope, weather, sunrise, sunset) {
  // Get the weather
  $scope.weather = weather;
  $scope.sunrise = sunrise;
  $scope.sunset = sunset;
}]);