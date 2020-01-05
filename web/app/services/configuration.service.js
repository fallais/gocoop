// ConfigurationService
app.service('ConfigurationService', ['$http', '$q', 'BASE_URL', function($http, $q, BASE_URL){
  // Get
  function get(id){
    var deferred = $q.defer();

    $http({
      method: 'GET',
      url: BASE_URL + '/api/v1/configuration'
    }).success(function (data) {
      deferred.resolve(data);
    }).error(function (msg) {
      deferred.reject(msg);
    });

    return deferred.promise;
  }

  // Publish functions
  return {
    Get: get
  };
}]);