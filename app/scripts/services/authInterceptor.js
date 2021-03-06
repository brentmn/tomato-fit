(function ($ng) {
    $ng.module('app.services')
        .factory('authInterceptor', ['$window', '$q', function ($window, $q) {
            return {
                request: function (config) {
                    config.headers = config.headers || {};
                    if ($window.sessionStorage.token) {
                        config.headers.Authorization = 'Bearer ' + $window.sessionStorage.token;
                    }
                    return config;
                },
                response: function (response) {
                    if (response.status === 401) {
                        // handle the case where the user is not authenticated
                    }
                    return response || $q.when(response);
                }
            };
        }])
        .config(function($httpProvider){
            $httpProvider.interceptors.push('authInterceptor');
        });
}(angular));