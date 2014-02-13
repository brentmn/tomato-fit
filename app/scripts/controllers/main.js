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

            console.log('pom', pomodoro);
            $scope.authorized = authorized;
            $scope.pomodoro = pomodoro;

            return;

            $scope.alarms = [];
            $scope.authorized = authorized;
            $scope.pomodoro = pomodoro;
            $scope.pause = clear;
            $scope.delete = deleteAlarm;
            $scope.valid = function (a) {
                return !a.deleted && a.enabled;
            };
            $scope.start = function (kind) {
                clear();
                var length = pomodoro[kind];
                if (!length) {
                    return;
                }

                $scope.pomodoro.counter = length;

                var d = new Date();

                d.setSeconds(d.getSeconds() + (length / 1000));
                var formatted = $filter('date')(d, 'hh:mmZ');

                var alarm = {
                    deviceId: $scope.selected,
                    id: null,
                    time: formatted,
                    enabled: true,
                    recurring: false,
                    weekDays: [],
                    label: kind
                    //snoozeLength int
                    //snoozeCount  int
                    //vibe         string
                };

                fitbit.setAlarm(alarm).success(function (res, code) {
                    if (code === 200) {
                        $scope.alarms.push(res.trackerAlarm);
                        doPomodoro();
                    }
                });

            };

            function doPomodoro() {
                var t = $scope.pomodoro.counter - 1000;
                $scope.pomodoro.counter = t;
                if (t <= 0)return;
                timer = $timeout(doPomodoro, 1000);
            }

            if (!authorized) {
                api.authorize().success(function (res, status) {
                    if (status === 200) {
                        $scope.authUrl = res.Url;
                    }
                });
            } else {
                fitbit.getDevices().success(function (res, status) {
                    $scope.selected = res[0].id;
                    $scope.devices = res;
                })
                    .then(getAlarms);
            }

            function getAlarms() {
                fitbit.getAlarms($scope.selected).success(function (res, code) {
                    if (code === 200) {
                        $scope.alarms = res.trackerAlarms;

                        var last = res.trackerAlarms[res.trackerAlarms.length - 1];

                        if (last.enabled && !last.deleted) {
                            var d = new Date();
                            var d2 = new Date();
                            d2.setDate(last.time);
                            console.log(d, d2);
                            //$scope.pomodoro.counter = (d.getSeconds - d2.getSeconds());
                        }
                    }
                });
            }
            function clear(){
                $timeout.cancel(timer);
            }
            function deleteAlarm(alarm) {

                fitbit.cancelAlarm($scope.selected, alarm.alarmId)
                    .then(getAlarms);
            }

        });
}());
