var app = angular.module("single", []);

app.controller('SingleCtr', function ($scope, $http) {
    $scope.ssss = "!!!!";

    $scope.initData = function () {
        $http.get("/article", {})
            .success(function (resp) {
                if (resp.code == 200) {
                    $scope.article = resp.data;
                } else {
                    alert(resp.msg);
                }
            });
    };
    $scope.initData();

    $scope.formatTime = function (str) {
        return formatDate(str)
    };
});

function formatDate(str, format) {
    if ("" == str) return "-";

    format = format || 'Y-m-d H:i';

    if ('0001-01-01T00:00:00Z' == str || str == undefined) {
        return '';
    }
    if ('0001-01-01 00:00:00 +0000' == str || str == undefined) {
        return '';
    }
    if (!str.match(/T/)) {
        str = str.replace(/\-/g, '/')
    }
    var dt = new Date(str),
        mon = dt.getMonth() + 1,
        d = dt.getDate(),
        h = dt.getHours(),
        m = dt.getMinutes(),
        s = dt.getSeconds();
    if (mon < 10) mon = '0' + mon;
    if (d < 10) d = '0' + d;
    if (h < 10) h = '0' + h;
    if (m < 10) m = '0' + m;
    if (s < 10) s = '0' + s;

    format = format.replace(/Y|y/, dt.getFullYear());
    format = format.replace(/M|m/, mon);
    format = format.replace(/d/, d);
    format = format.replace(/H|h/, h);
    format = format.replace(/i/, m);
    format = format.replace(/s/, s);

    return format;
}