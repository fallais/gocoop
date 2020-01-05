app.config(['$stateProvider', '$locationProvider', '$urlRouterProvider', '$httpProvider', function($stateProvider, $locationProvider, $urlRouterProvider, $httpProvider) {

  $locationProvider.html5Mode(false);
  $locationProvider.hashPrefix('!');

  $stateProvider.
    state('dashboard', {
      url: '/dashboard',
      templateUrl: '/app/views/dashboard.html',
      controller: 'DashboardCtrl',
      resolve : {
        door: function(DoorService) {
          return DoorService.Status();
        },
      }
    }).
    state('door', {
      url: '/door',
      templateUrl: '/app/views/door.html',
      controller: 'DoorGetCtrl',
      resolve: {
        door: function(DoorService) {
          return DoorService.Status();
        }
      }
    }).
    state('configuration', {
      url: '/configuration',
      templateUrl: '/app/views/configuration.html',
      controller: 'ConfigurationGetCtrl',
      resolve: {
        configuration: function(ConfigurationService) {
          return ConfigurationService.Get();
        }
      }
    });

    $urlRouterProvider.otherwise('/dashboard');
}]);