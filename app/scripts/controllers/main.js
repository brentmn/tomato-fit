'use strict';
(function () {

    var pomodoro = {
        pomodoro: 1500000,
        short: 300000,
        long: 600000
    };

    var timer, authorized;

    angular.module('tomtatoFitApp')
        .controller('MainCtrl', function ($window, $scope, $timeout, $cookies, $filter, api, pomodoro) {
            if ($cookies.jwt) {
                $window.sessionStorage.token = $cookies.jwt;
                authorized = true;
            }

            $scope.authorized = authorized;
            $scope.pomodoro = pomodoro;

            $scope.filterAlarms = function(a){
                return a.enabled && !a.deleted;
            };

        });
}());
