// LightService
app.service('LightService', ['$http', '$q', 'BASE_URL', function($http, $q, BASE_URL){
  // Status
  function status(id){
    var deferred = $q.defer();

    $http({
      method: 'GET',
      url: BASE_URL + '/api/v1/light'
    }).success(function (data) {
      deferred.resolve(data);
    }).error(function (msg) {
      deferred.reject(msg);
    });

    return deferred.promise;
  }

  // Open
  function open(){
    var deferred = $q.defer();

     $http({
       method: 'GET',
       url: BASE_URL + '/api/v1/light?action=open'
     }).success(function (data) {
       deferred.resolve(data);
     }).error(function (msg) {
       deferred.reject(msg);
     });

     return deferred.promise;
  }

  // Close
  function close(){
    var deferred = $q.defer();

    $http({
      method: 'GET',
      url: BASE_URL + "/api/v1/light?action=close"
    }).success(function (data) {
      deferred.resolve(data);
    }).error(function (msg) {
      deferred.reject(msg);
     });

    return deferred.promise;
  }

  // Publish functions
  return {
    Status: status,
    Open: open,
    Close: close
  };
}]);