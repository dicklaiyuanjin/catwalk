$(document).ready(function(){
  function ivttWebsocket() {
    const socket = new WebSocket('ws://' + window.location.host + '/ws/ivtt/join?username=' + $('#username').val());
    
    socket.onmessage = function (event) {
      var data = JSON.parse(event.data);
      console.log(event);
      console.log(data);
    }

  }
 


  /**********************************
  * invitation part
  ***********************************/
  function recEvelope(sendername) {
    $("#" + sendername + "-invite-content").hide();

    $("#" + sendername + "-invite-brand").click(function(){ 
      $("#" + sendername + "-invite-content").slideToggle();
    });

    function removeEnvelope(sdrname, action) {
      var data = JSON.stringify({
        sender: sdrname,
        receiver: $('#username').val(),
        msg: ""
      });

      if (action == "agree") {
        $.post("/app/invitation/agree", data, function(data, status){
          console.log("status: ", status);
          console.log("data: ", data);
        });
      } else if (action == "refuse") {
        $.post("/app/invitation/refuse", data, function(data, status){
          console.log("status: ", status);
          console.log("data: ", data);
        });
      }

      var idname = "#" + sdrname + "-envelope";
      $(idname).remove();
    }
     
    $("#" + sendername + "-agree").click(function(){
      removeEnvelope(sendername, "agree"); 
    });
    
    $("#" + sendername + "-refuse").click(function(){ 
      removeEnvelope(sendername, "refuse");
    }); 
  }

  function recEvelopeCtlr() {
    var recArr = $('#rec-envelope').children();
    var sdrname;
    for (var i = 0; i <recArr.length; ++i) {
      sdrname = recArr.eq(i).attr("id").slice(0, -7);
      recEvelope(sdrname);
    }
  }
  

  function senderEvelopeCtlr() {
    $("#invite-send").click(function(){
      var data = JSON.stringify({
        sender: $('#nickname').val(),
        receiver: $('#receiver-name').val(),
        msg: $('#invite-message').val()
      });
      $('#receiver-name').val("");
      $('#invite-message').val("");

      $.post("/app/invitation/send", data, function(data, status){
        console.log("status: ", status);
        console.log("data: ", data);
      });

    });
  }


  recEvelopeCtlr();
  senderEvelopeCtlr();
  ivttWebsocket();



  
});

