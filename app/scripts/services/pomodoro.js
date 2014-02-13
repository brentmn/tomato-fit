(function ($ng) {
    'use strict';
    var svc = function ($timeout, $log, fitbit) {
        var pub = {
            devices: [],
            device: undefined,
            alarms: [],
            alarm: undefined,
            timers: {
                pomodoro: 1500000,
                short: 300000,
                long: 600000
            },
            counter: 0,
            countdown: "",
            timer: undefined
        };


        fitbit.getDevices().success(function (data, status) {
            if (status !== 200) {
                $log.log('get devices error', status, data);
                return;
            }
            pub.devices = pub.devices.concat(data);
            if (!pub.devices || !pub.devices.length) return;
            pub.device = pub.devices[0];

        }).then(function () {
                fitbit.getAlarms(pub.device.id).success(function (data, status) {
                    if (status !== 200) {
                        $log.log('get device alarms error', status, data);
                        return;
                    }
                    if (data && data.trackerAlarms && data.trackerAlarms.length) {
                        pub.alarms = pub.alarms.concat(data.trackerAlarms);
                        pub.alarm = pub.alarms[0];
                        startAlarm(pub.alarm)
                    }
                });
            });

        function startAlarm(trackerAlarm) {
            var d = moment(trackerAlarm.time, 'HH:mmZ').toDate();
            var d2 = new Date();

            pub.counter = (d.getTime() - d2.getTime());

            start();
        }

        function start(kind) {
            clear();
            if (kind) {
                pub.counter = pub.timers[kind];
            }
            var d = new Date();
            d.setSeconds(d.getSeconds() + (pub.counter / 1000));
            var alarm = {
                deviceId: pub.device.id,
                id: null,
                time: moment(d).format('HH:mmZ'),
                enabled: true,
                recurring: false,
                weekDays: [],
                label: kind
            };
            //if generate by the user
            if (kind) {
                fitbit.setAlarm(alarm).success(function (data, status) {
                    if (status !== 200) {
                        $log.log('set alarm error', alarm, status, data);
                        return;
                    }
                    pub.alarm = data.trackerAlarm;
                    pub.alarms.push(data.trackerAlarm);
                });
            }

            pub.alarm = alarm;
            doCountdown();
        }

        function doCountdown() {
            pub.timer = $timeout(function () {
                pub.counter = (pub.counter - 1000);
                pub.countdown = moment(pub.alarm.time, 'HH:mm:ssZ').fromNow(true);
                doCountdown();
            }, 1000);
        }

        function remove(alarm) {
            fitbit.cancelAlarm(pub.device.id, alarm.alarmId).success(function (data, status) {
                if (status !== 200) {
                    $log.log('remove alarm error', alarm, status, data);
                }
            });
        }

        function clear() {
            $timeout.cancel(pub.timer);
        }

        pub.start = start;
        pub.remove = remove;
        return pub;
    };

    $ng.module('app.services').factory('pomodoro', svc);


}(angular));