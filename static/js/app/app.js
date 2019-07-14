$(document).ready(function(){
  function ivttWebsocket() {
    const socket = new WebSocket('ws://' + window.location.host + '/ws/join?username=' + $('#username').val());
    
    socket.onmessage = function (event) {
      console.log("event.data: ", event.data);
      if (event.data != "") {
        var data = JSON.parse(event.data);
        RecData(data, socket);
      }
    }
    
    
    $("#invite-send").click(function(){
      var data = InitData({
        sender: $('#username').val(),
        receiver: $('#receiver-name').val(),
        msg: $('#invite-message').val()
      }, 0);

      $('#receiver-name').val("");
      $('#invite-message').val(""); 
      socket.send(data);
    });

       
  }
 
  /*
   * t is data type
   * 0: invitation
   * 1: reply
   * 2: friendinfo
   */
  function InitData(data, t) {
    var ws = {
      code: t,
      ivtt: {
        sender: "",
        receiver: "",
        msg: ""
      },
      rpl: {
        me: "",
        obj: "",
        attitude: ""
      },
      fif : {
        username: "",
        nickname: "",
        email: "",
        motto: "",
        icon: ""
      }
   };

    switch (t) {
    case 0:
      ws.ivtt = data;
      break;
    case 1:
      ws.rpl = data;
      break;
    case 2:
      ws.fif = data;
      break;
    }
    return JSON.stringify(ws);
  }

  function RecData(data, socket) {
    switch (data.code) {
    case 0:
      rec_ivtt(data.ivtt, socket);
      break;
    case 1:
      rec_rpl(data.rpl, socket);
      break;
    case 2:
      rec_fif(data.fif, socket);
      break;
    }
  }

  function rec_ivtt(data, socket) {
    if (data.sender != $('#username').val()) {
      if (!isSenderExist(data.sender)) {
        $("#rec-envelope").append(newEnvelope(data));
        recEnvelope(data.sender, socket);
      }
    }
  }

  function rec_rpl(data, socket) {
    console.log("ivtt.rpl: ", data);
  }

  function rec_fif(data, socket) {
    console.log("fif: ", data);
  }

  /**********************************
  * invitation part
  ***********************************/
  function recEnvelope(sendername, socket) {
    $("#" + sendername + "-invite-content").hide();

    $("#" + sendername + "-invite-brand").click(function(){ 
      $("#" + sendername + "-invite-content").slideToggle();
    });

    function removeEnvelope(sdrname, att) {
      var data = InitData({
        me: $("#username").val(),
        obj: sdrname,
        Attitude: att
      }, 1);

      socket.send(data);
      
      console.log("removeEnvelope: ", data);
      var idname = "#" + sdrname + "-invite";
      console.log("remove idname: ", idname);
      $(idname).remove();
    }
     
    $("#" + sendername + "-agree").click(function(){
      removeEnvelope(sendername, "agree"); 
    });
    
    $("#" + sendername + "-refuse").click(function(){ 
      removeEnvelope(sendername, "refuse");
    }); 
  }


  function recEnvelopeCtlr() {
    var recArr = $('#rec-envelope').children();
    var sdrname;
    for (var i = 0; i <recArr.length; ++i) {
      sdrname = recArr.eq(i).attr("id").slice(0, -7);
      recEnvelope(sdrname);
    }
  }
  

  


  recEnvelopeCtlr(); 
  ivttWebsocket();



  function isSenderExist(sdr) {
    return $("#" + sdr).length != 0;
  }
  

  function newEnvelope(data) {
    var invite = document.createElement('div');
    invite.setAttribute('id', data.sender + "-invite");

    var inviteBrand = document.createElement('li');
    inviteBrand.setAttribute('id', data.sender + "-invite-brand");
    inviteBrand.setAttribute('class', "list-group-item");

    var enelopeIcon = document.createElement('span');
    enelopeIcon.setAttribute('class', "glyphicon glyphicon-envelope");
    enelopeIcon.setAttribute('aria-hidden', "true");

    var brandname = document.createElement('span');
    brandname.setAttribute('id', data.sender);
    brandname.innerHTML = " " + data.sender;

    var arrowIcon = document.createElement('span');
    arrowIcon.setAttribute('class', "myarrow glyphicon glyphicon-arrow-down");
    arrowIcon.setAttribute('aria-hidden', "true");

    inviteBrand.append(enelopeIcon, brandname, arrowIcon);


    var inviteContent = document.createElement('li');
    inviteContent.setAttribute('id', data.sender + "-invite-content");
    inviteContent.setAttribute('class', "list-group-item");

    var inviteForm = document.createElement('div');
    inviteForm.setAttribute("id", data.sender + "-inviteform");
    inviteForm.setAttribute("class", "form-horizontal");

    var formgroup1 = document.createElement('div');
    formgroup1.setAttribute("class", "form-group");
    
    var label = document.createElement('label');
    label.setAttribute("class", "col-sm-2 control-label");
    label.innerHTML = "Message";

    var labelcontent = document.createElement('div');
    labelcontent.setAttribute("class", "col-sm-10 padtop7px");
    
    var labeltext = document.createElement('p');
    labeltext.setAttribute("type", "text");
    labeltext.setAttribute("id", data.sender + "invitemsg");
    labeltext.innerHTML = data.msg;

    labelcontent.append(labeltext);
    formgroup1.append(label, labelcontent);

    var formgroup2 = document.createElement('div');
    formgroup2.setAttribute("class", "form-group");

    var btnouter1 = document.createElement('div');
    btnouter1.setAttribute("class", "col-sm-offset-2 col-sm-10");

    var btn1 = document.createElement('button');
    btn1.setAttribute("id", data.sender + "-agree");
    btn1.setAttribute("class", "widthbtn btn btn-primary");
    btn1.setAttribute("type", "submit");
    btn1.innerHTML = "Agree";

    btnouter1.append(btn1);
    formgroup2.append(btnouter1);
    
    var formgroup3 = document.createElement('div');
    formgroup3.setAttribute("class", "form-group");
    
    var btnouter2 = document.createElement('div');
    btnouter2.setAttribute("class", "col-sm-offset-2 col-sm-10");

    var btn2 = document.createElement('button');
    btn2.setAttribute("id", data.sender + "-refuse");
    btn2.setAttribute("class", "widthbtn btn btn-success");
    btn2.setAttribute("type", "submit");
    btn2.innerHTML = "Refuse";

    btnouter2.append(btn2);
    formgroup3.append(btnouter2);
    inviteForm.append(formgroup1, formgroup2, formgroup3);

    inviteContent.append(inviteForm);
    
    invite.append(inviteBrand, inviteContent);
    return invite;
  }

});



