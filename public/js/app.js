(function () {
    var blogApp = angular.module("blogApp", [
        "ngRoute",
        "postsControllers"
    ]);

    blogApp.config(['$routeProvider', 
        function($routeProvider) {
            $routeProvider.
                when('/posts', {
                    templateUrl: 'views/posts.html',
                    controller: 'PostsListCtrl'
                }).
                when('/posts/:postId', {
                    templateUrl: 'views/post_details.html',
                    controller: 'PostDetailsCtrl'
                }).
                when('/posts/edit/:postId', {
                    templateUrl: 'views/edit_post.html',
                    controller: 'EditPostCtrl'
                }).
                otherwise({
                    redirectTo: "/posts"
                });
        }]);
})()
