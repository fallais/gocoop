'use strict';

var app = angular.module('gocoop', ['ui.router', 'ngFlash']);

//------------------------------------------------------------------------------
// Constant
//------------------------------------------------------------------------------

app.constant('C_VERSION', '0.1.0');
app.constant('BASE_URL', 'http://localhost:8000');

//------------------------------------------------------------------------------
// Filters
//------------------------------------------------------------------------------

// weatherIcon
app.filter('weatherIcon', [ function(){
  return function(val) {
    switch(val[0].main){
      case "Clear":
        return "wi-day-sunny";
        break;
      case "Rain":
      case "Drizzle":
        return "wi-day-rain";
        break;
      case "Thunderstorm":
        return "wi-day-thunderstorm";
        break;
      case "Snow":
        return "wi-day-snow";
        break;
      case "Clouds":
        return "wi-day-cloudy";
        break;
      case "Atmosphere":
        return "wi-day-haze";
        break;
      case "Extreme":
        return "";
        break;
      default:
        return "";
        break;
    }
  };
}]);

// fancyDate with Moment
app.filter('fancyDate', [ function(){
  return function(dateval) {
    return moment(dateval).format("DD MMM YYYY @ HH:mm")
  }
}]);

// fancyDateSimple with Moment
app.filter('fancyDateSimple', [ function(){
  return function(dateval) {
    return moment(dateval).format("HH:mm")
  }
}]);

//------------------------------------------------------------------------------
// Run
//------------------------------------------------------------------------------

app.run(['$rootScope', '$state', function($rootScope, $state) {
  $rootScope.loading = false;

  $rootScope.$on('$stateChangeStart', function() {
    $rootScope.loading = true;
    $rootScope.routeError = false;
  });

  $rootScope.$on('$stateChangeSuccess', function() {
    $rootScope.loading = false;
    $rootScope.routeError = false;
  });
  
  $rootScope.$on('$stateChangeError', function (event, current, previous, rejection) {
    $rootScope.loading = false;
    $rootScope.routeError = true;
    event.preventDefault();
  });
}]);
