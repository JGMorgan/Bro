angular.module('app.controllers', [])

.controller('homeCtrl', ['$scope', '$stateParams','$state', // The following is the constructor function for this page's controller. See https://docs.angularjs.org/guide/controller
// You can include any angular dependencies as parameters for this function
// TIP: Access Route Parameters for your page via $stateParams.parameterName
function ($scope, $stateParams) {


}])

.controller('loginCtrl', ['$scope', '$stateParams','$state','$http', // The following is the constructor function for this page's controller. See https://docs.angularjs.org/guide/controller
// You can include any angular dependencies as parameters for this function
// TIP: Access Route Parameters for your page via $stateParams.parameterName
function ($scope, $stateParams, $state,$http) {
      $scope.login=function(){
           $http.post('',{
                username:this.formdata.username,
                password:this.formdata.password
           })
           console.log(this.formdata.username);
           console.log(this.formdata.password);
           $state.go('menu.home');
     }

}])

.controller('broWithCtrl', ['$scope', '$stateParams','$state', // The following is the constructor function for this page's controller. See https://docs.angularjs.org/guide/controller
// You can include any angular dependencies as parameters for this function
// TIP: Access Route Parameters for your page via $stateParams.parameterName
function ($scope, $stateParams) {
     var intervalId;
     var websocket = new WebSocket('ws://localhost:8420/chat');
     var brotext;
     $scope.brotextcalc=function($scope, $stateParams, $state){
          brotext = document.getElementById('broText');
          broText.value="Br";
          intervalId=setInterval(appendText, 100);
     }

     $scope.stopText=function(){

          clearInterval(intervalId);
          var chatbox = document.getElementById('chatbox');
          angular.element(document.querySelector('#chatbox')).html(broText.value);
          websocket.send(broText);
     }
     function appendText(){
          brotext = document.getElementById('broText');

          broText.value+="o";
     }
     websocket.onmessage = function (evt)
     {
          var received_msg = evt.data;
          alert("Message is received...");
     };

}])

.controller('signupCtrl', ['$scope', '$stateParams','$state', // The following is the constructor function for this page's controller. See https://docs.angularjs.org/guide/controller
// You can include any angular dependencies as parameters for this function
// TIP: Access Route Parameters for your page via $stateParams.parameterName
function ($scope, $stateParams) {


}])
