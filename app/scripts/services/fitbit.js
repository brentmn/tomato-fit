(function ($ng) {
    $ng.module('app.services')
        .factory('fitbit', ['$http', function ($http) {

            function getDevices() {
                return $http.get('/device');
            }

            function getAlarms(deviceId) {
                return $http.get('/alarm/' + deviceId);
            }

            function setAlarm(alarm) {
                return $http.post('/alarm', alarm);
            }

            function cancelAlarm(deviceId, alarmId) {
                return $http.delete('/alarm/' + deviceId + '/' + alarmId);
            }

            return {
                getDevices: getDevices,
                getAlarms: getAlarms,
                setAlarm: setAlarm,
                cancelAlarm: cancelAlarm
            };
        }]);
}(angular));