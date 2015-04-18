(function() {
    var postsControllers = angular.module("postsControllers", []);

    postsControllers.controller("PostsListCtrl", function($scope, $http) {
        $scope.posts = [];
        $http.get("/api/posts").success(function(data) {
            $scope.posts = data;
        });
        $scope.deletePost = function(postId) {
            $http.delete("/api/posts/" + postId).success(function() {
                $scope.posts.forEach(function(element, index) {
                    if (element.id == postId) {
                        $scope.posts.splice(index, 1);
                    }
                });
            });
        };
    });

    postsControllers.controller("PostDetailsCtrl", function($scope, $http, $routeParams) {
        $http.get("/api/posts/" + $routeParams.postId).success(function(data) {
            $scope.post = data;
        });
    });

    postsControllers.controller("EditPostCtrl", function($scope, $http, $routeParams, $location) {
        $scope.submitIsDisabled = true;

        if ($routeParams.postId == 0) {
            $scope.submitIsDisabled = false;
        } else {
            $http.get("/api/posts/" + $routeParams.postId).success(function(data) {
                $scope.post = data;
                $scope.submitIsDisabled = false;
            });
        }

        $scope.submitForm = function() {
            var successCallback = function(data) {
                $location.path("/posts");
            };
            if ($routeParams.postId == 0) {
                $http.post("/api/posts", $scope.post).success(successCallback);
            } else {
                $http.patch("/api/posts/" + $routeParams.postId, $scope.post).success(successCallback);
            }
        };
        $scope.cancelForm = function() {
            $location.path("/posts");
        }
    });
})();
