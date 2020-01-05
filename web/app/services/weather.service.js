// WeatherService
app.service('WeatherService', ['$http', '$q', 'BASE_URL', function($http, $q, BASE_URL){
  // Get
  function getToday(){
    var deferred = $q.defer();

    $http({
      method: 'GET',
      url: BASE_URL + '/api/v1/weather/today'
    }).success(function (data) {
      deferred.resolve(data);
    }).error(function (msg) {
      deferred.reject(msg);
    });

    return deferred.promise;
  }

  // Get
  function getSunrise(){
    var deferred = $q.defer();

    $http({
      method: 'GET',
      url: BASE_URL + '/api/v1/weather/sunrise'
    }).success(function (data) {
      deferred.resolve(data);
    }).error(function (msg) {
      deferred.reject(msg);
    });

    return deferred.promise;
  }

  // Get
  function getSunset(){
    var deferred = $q.defer();

    $http({
      method: 'GET',
      url: BASE_URL + '/api/v1/weather/sunset'
    }).success(function (data) {
      deferred.resolve(data);
    }).error(function (msg) {
      deferred.reject(msg);
    });

    return deferred.promise;
  }

  // Publish functions
  return {
    GetToday: getToday,
    GetSunrise: getSunrise,
    GetSunset: getSunset
  };
}]);