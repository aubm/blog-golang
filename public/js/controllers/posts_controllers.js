(function() {
    var postsControllers = angular.module("postsControllers", []);

    postsControllers.controller("PostsListCtrl", function($scope) {
        $scope.pageHeader = "Hello world !";
    });

    postsControllers.controller("PostDetailsCtrl", function($scope) {
    });

    postsControllers.controller("EditPostCtrl", function($scope) {
    });
})()
