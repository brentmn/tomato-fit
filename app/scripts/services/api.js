(function ($ng) {
    $ng.module('app.services')
        .factory('api', ['$http', '$log', function ($http, $log) {

            function authorize(){
                $log.log('api authorizing');
                return $http.post('/authorize');
            }

            return {
                authorize: authorize
            };
        }]);
}(angular));