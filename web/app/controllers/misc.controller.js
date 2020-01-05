// DashboardCtrl
app.controller("DashboardCtrl", ['$scope', 'door', 'DoorService', function($scope, door, DoorService) {
  // Get the door
  $scope.door = door;


  // Open the door
  $scope.openDoor = function(){
    $scope.doorOpeningLoading = true;
    DoorService.Open()
    .then(function (data, status, headers, config) {
      $scope.doorOpeningLoading = false;
      console.log("Door has been opened !")
    })
    .catch(function (error) {
      $scope.doorOpeningLoading = false;
      console.log(error);
      $scope.error = error;
    });
  }

  // Close the door
  $scope.closeDoor = function(){
    $scope.doorClosingLoading = true;
    DoorService.Close()
    .then(function (data, status, headers, config) {
      $scope.doorClosingLoading = false;
      console.log("Door has been closed !")
    })
    .catch(function (error) {
      $scope.doorClosingLoading = false;
      console.log(error);
      $scope.error = error;
    });
  }

}]);

// NavCtrl
app.controller("NavCtrl", ['$scope', '$state', function($scope, $state) {}]);

// CamerasCtrl
app.controller("CamerasCtrl", ['$scope', '$state', function($scope, $state) {}]);

// ConfigurationCtrl
app.controller("ConfigurationGetCtrl", ['$scope', '$state', 'configuration', function($scope, $state, configuration) {
  $scope.configuration = configuration;
}]);