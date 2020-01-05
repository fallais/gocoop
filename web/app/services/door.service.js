// DoorService
app.service('DoorService', ['$http', '$q', 'BASE_URL', function($http, $q, BASE_URL){
  // Status
  function status(id){
    var deferred = $q.defer();

    $http({
      method: 'GET',
      url: BASE_URL + '/api/v1/door'
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
       url: BASE_URL + '/api/v1/door/use?action=open&force=false'
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
      url: BASE_URL + "/api/v1/door/use?action=close&force=false"
    }).success(function (data) {
      deferred.resolve(data);
    }).error(function (msg) {
      deferred.reject(msg);
     });

    return deferred.promise;
  }

  // Open (forced)
  function openForced(){
    var deferred = $q.defer();

     $http({
       method: 'GET',
       url: BASE_URL + '/api/v1/door/use?action=open&force=true'
     }).success(function (data) {
       deferred.resolve(data);
     }).error(function (msg) {
       deferred.reject(msg);
     });

     return deferred.promise;
  }

  // Close (forced)
  function closeForced(){
    var deferred = $q.defer();

    $http({
      method: 'GET',
      url: BASE_URL + "/api/v1/door/use?action=close&force=true"
    }).success(function (data) {
      deferred.resolve(data);
    }).error(function (msg) {
      deferred.reject(msg);
     });

    return deferred.promise;
  }

  // Stop (forced)
  function stop(){
    var deferred = $q.defer();

    $http({
      method: 'GET',
      url: BASE_URL + "/api/v1/door/use?action=stop"
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
    Close: close,
    OpenForced: openForced,
    CloseForced: closeForced,
    Stop: stop
  };
}]);