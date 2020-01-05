// DoorGetCtrl
app.controller("DoorGetCtrl", ['$scope', '$state', 'door', function($scope, $state, door) {
  // Get the door
  $scope.door = door;

  // Open the door
  $scope.openDoor = function(){
    $scope.doorOpeningLoading = true;
    DoorService.Open()
    .then(function (data, status, headers, config) {
      $scope.doorOpeningLoading = false;
      console.log("Door opened")
      $state.reload();
      //Flash.create('success', "Door has been opened !", 5000, null, true);
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
      console.log("Door closed")
      $state.reload();
      //Flash.create('success', "Door has been opened !", 5000, null, true);
    })
    .catch(function (error) {
      $scope.doorClosingLoading = false;
      console.log(error);
      $scope.error = error;
    });
  }

  // Open the door (forced)
  $scope.openDoorForced = function(){
    $scope.doorOpeningLoading = true;
    DoorService.OpenForced()
    .then(function (data, status, headers, config) {
      $scope.doorOpeningLoading = false;
      console.log("Door opened by force")
      $state.reload();
      //Flash.create('success', "Door has been opened !", 5000, null, true);
    })
    .catch(function (error) {
      $scope.doorOpeningLoading = false;
      console.log(error);
      $scope.error = error;
    });
  }

  // Close the door (forced)
  $scope.closeDoorForced = function(){
    $scope.doorClosingLoading = true;
    DoorService.CloseForced()
    .then(function (data, status, headers, config) {
      $scope.doorClosingLoading = false;
      console.log("Door closed by force")
      $state.reload();
      //Flash.create('success', "Door has been opened !", 5000, null, true);
    })
    .catch(function (error) {
      $scope.doorClosingLoading = false;
      console.log(error);
      $scope.error = error;
    });
  }

  // Stop the door (forced)
  $scope.stopDoor = function(){
    DoorService.Stop()
    .then(function (data, status, headers, config) {
      $scope.doorClosingLoading = false;
      console.log("Door stopped by force")
      $state.reload();
      //Flash.create('success', "Door has been opened !", 5000, null, true);
    })
    .catch(function (error) {
      $scope.doorClosingLoading = false;
      console.log(error);
      $scope.error = error;
    });
  }
}]);