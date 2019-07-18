$(document).ready(function(){ 

  function Websocket(callback_func) {      
    const socket = new WebSocket('ws://' + window.location.host + '/ws/join?username=' + $('#username').val());
    socket.onmessage = function (event) {
      console.log("event.data: ", event.data)
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

    function initFriroom() {
      var frirooms = document.getElementsByClassName("friroom");
      var temp;
      for (var i = 0; i < frirooms.length; ++i) {
       temp = frirooms[i].getAttribute('id');
       temp = temp.slice(0, -5);
       FriroomListen(temp, socket);
      }
    }
  
    function initFribox() {
      var friboxes = document.getElementsByClassName("fribox");
      var temp;
      for (var i = 0; i < friboxes.length; ++i) {
        temp = friboxes[i].getAttribute('id');
        temp = temp.slice(0, -4); // getFriname
        FriboxListen(temp, socket);
      }
    }

    function initFrinfo() {
      var frinfos = document.getElementsByClassName("frinfo");
      var temp;
      for (var i = 0; i < frinfos.length; ++i) {
        temp = frinfos[i].getAttribute('id');
        temp = temp.slice(0, -5);
        FrinfoListen(temp, socket);
      }
    }
    
    initFrinfo();
    initFriroom();
    initFribox();
    callback_func(socket);
    
  }//end Websocket()
 
  /*
   * t is data type
   * 0: invitation
   * 1: reply
   * 2: friendinfo
   * 3: msg
   * 4: del(delete)
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
      fif: {
        username: "",
        nickname: "",
        email: "",
        motto: "",
        icon: ""
      },
      msg: {
        sender: "",
        receiver: "",
        content: "",
        sendtime: ""
      },
      del: {
        sender: "",
        exfri: "",
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
    case 3:
      ws.msg = data;
      break;
    case 4:
      ws.del = data;
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
    case 3:
      rec_msg(data.msg, socket);
      break;
    case 4:
      rec_del(data.del, socket);
      break;
    }
  }

  function rec_del(data, socket) {
    console.log("rec.del: ", data);
    document.location.href = "/app";
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
    //在对方同意我的邀请的情况下，如果对方曾向我发起邀请，那么就删除该邀请
    if ($("#" + data.me + "-invite").length != 0) {
      $("#" + data.me + "-invite").remove();
    }
  }

  function rec_fif(data, socket) {
    if (!isFriExist(data.username)) {
      $("#frimain").append(newFriBox(data));
      FriboxListen(data.username, socket);
      //收到好友信息后，应该建立相应的friroom和frinfo
      //建立完后添加相应的事件响应（callback）
      $("#app-room").append(newFriroom(data));
      FriroomListen(data.username, socket);

      $("#app-frinfo").append(newFrinfo(data));
      FrinfoListen(data.username, socket);
    }
  }

  function rec_msg(data, socket){
    console.log("rec msg data: ", data);
    var usr = $("#username").val();
    var friname = "";
    if(data.sender == usr) {
      //发送者为本人，放在相应的friroom中
      friname = data.receiver;
      $('#' + friname + "-room-msglist").append(newMsgBox(data.content, "right"));
    } else {
      //发送者为朋友，放到相应的friroom中
      friname = data.sender;
      $('#' + friname + "-room-msglist").append(newMsgBox(data.content, "left"));
    }

    var ta = document.getElementById(friname + "-room-msglist");
    if (ta.scrollHeight != null) {
      ta.scrollTop = ta.scrollHeight;
    }
    $("#" + friname + "-box-icon>img:first").attr("class", "fribox-img-rec fribox-img img-circle img-responsive");
  }

  /**********************************
  * invitation part
  ***********************************/
  function recEnvelope(sendername, socket) {
    $("#" + sendername + "-invite-content").hide();

    $("#" + sendername + "-invite-brand").click(function(){ 
      $("#" + sendername + "-invite-content").slideToggle();
    });


    var obj1 = {
      sdrname: sendername,
      att: "agree",
      socket: socket,
      removeEnvelope: function() {
        var data = InitData({
          me: $("#username").val(),
          obj: this.sdrname,
          attitude: this.att
        }, 1);

        this.socket.send(data);

        var idname = "#" + this.sdrname + "-invite";
        $(idname).remove();  
      }
    }
    
    var obj2 = {
      sdrname: sendername,
      att: "refuse",
      socket: socket,
      removeEnvelope: function() {
        var data = InitData({
          me: $("#username").val(),
          obj: this.sdrname,
          attitude: this.att
        }, 1);

        this.socket.send(data);

        var idname = "#" + this.sdrname + "-invite";
        $(idname).remove();  
      }
    }
    /*
    function removeEnvelope(sdrname, att) {
      var data = InitData({
        me: $("#username").val(),
        obj: sdrname,
        attitude: att
      }, 1);

      socket.send(data);

      var idname = "#" + sdrname + "-invite";
      $(idname).remove();
    }
    */
    


    $("#" + sendername + "-agree").click(jQuery.proxy(obj1, "removeEnvelope"));
    
    $("#" + sendername + "-refuse").click(jQuery.proxy(obj2, "removeEnvelope")); 
  }


  function recEnvelopeCtlr(socket) {
    var recArr = $('#rec-envelope').children();
    var sdrname;
    for (var i = 0; i <recArr.length; ++i) {
      sdrname = recArr.eq(i).attr("id").slice(0, -7);
      recEnvelope(sdrname, socket);
    }
  }
 
  /*******************************
  * fribox
  *******************************/
  function clickFribox(friname, socket) {
    $("#app").attr("class", "apphide");
    $("#app-room").attr("class", "appshow");
    $("#app-frinfo").attr("class", "apphide");

    var frirooms = document.getElementsByClassName("friroom");
    var temp1;
    var temp2;
    for (var i = 0; i < frirooms.length; ++i) {
      temp1 = frirooms[i].getAttribute('id');
      temp1 = temp1.slice(0, -5);
      if (friname == temp1) {
        frirooms[i].setAttribute('class', 'app-active friroom');
      } else {
        frirooms[i].setAttribute('class', 'app-unactive friroom');
      }
    }

    $('#' + friname + "-room-btn").click(function(){
      var data = InitData({
        sender: $("#username").val(),
        receiver: friname,
        content: $("#" + friname + "-room-input").val(),
        sendtime: new Date().Format("yyyy-MM-dd hh:mm:ss") 
      }, 3);
      
      socket.send(data);
      $("#" + friname + "-room-input").val("");
    });

    $("#" + friname + "-box-icon>img:first").attr("class", "fribox-img img-circle img-responsive");
  }

  function FriboxListen(name, socket) {
    $('#' + name + "-box").click(function(){
      clickFribox(name, socket);
      var ta = document.getElementById(name + "-room-msglist");
      ta.scrollTop = ta.scrollHeight;
      $("#" + name + "-box-icon>img:first").attr("class", "fribox-img img-circle img-responsive");
    });
  }

  

  /**********************************************
  * friroom
  **********************************************/
  function clickRoomBack(name, socket) {
    $('#' + name + "-room").attr('class', 'app-unactive friroom');
    $('#app-room').attr('class', 'apphide');
    $('#app').attr('class', 'appshow');
  }

  function clickRoomIcon(name, socket) {
    $('#' + name + "-room").attr('class', 'app-unactive friroom');
    $('#app-room').attr('class', 'apphide');
    $('#app-frinfo').attr('class', 'appshow');
    $('#' + name + '-info').attr('class', 'app-active frinfo');
  }

  function FriroomListen(name, socket) {
    $("#" + name + "-room-back").click(function(){
      clickRoomBack(name, socket);
      $("#" + name + "-box-icon>img:first").attr("class", "fribox-img img-circle img-responsive");
    });

    $("#" + name + "-room-icon").click(function(){
      clickRoomIcon(name, socket);
      $("#" + name + "-box-icon>img:first").attr("class", "fribox-img img-circle img-responsive");
    });
  }

  
  
  /*********************************************************
  * frinfo
  *********************************************************/
  
  function clickInfoBack(name, socket) {
    $('#' + name + '-info').attr('class', 'app-unactive frinfo');
    $('#app-frinfo').attr('class', 'apphide');
    $('#app-room').attr('class', 'appshow');
    $('#' + name + '-room').attr('class', 'app-active friroom');

  }

  function FrinfoListen(name, socket) {
    $('#' + name + '-info-back').click(function(){
      clickInfoBack(name, socket);
      var ta = document.getElementById(name + "-room-msglist");
      ta.scrollTop = ta.scrollHeight;
    });

    $('#' + name + "-info-delete-btn").click(function(){
      var data = InitData({
        sender: $("#username").val(),
        exfri: name
      }, 4);

      socket.send(data);
    });
  }

  


  /*********************************
  * main
  *********************************/
  Websocket(recEnvelopeCtlr);


  /************************************
  * helper
  ************************************/
  function isSenderExist(sdr) {
    return $("#" + sdr).length != 0;
  }

  function isFriExist(fri) {
    return $("#" + fri + "-box").length != 0;
  }
  

  function newFriBox(data) {
    var fribox = document.createElement('div')
    fribox.setAttribute('id', data.username + "-box");
    fribox.setAttribute('class', 'fribox col-xs-4 col-sm-3 col-md-2')

    var content = document.createElement('div');
    content.setAttribute('id', data.username + '-box-content');
    content.setAttribute('class', 'fribox-content container-fluid')

    var icon = document.createElement('div');
    icon.setAttribute('id', data.username +'-box-icon');
    icon.setAttribute('class', 'row fribox-icon');

    var img = document.createElement('img');
    img.setAttribute('src', data.icon);
    img.setAttribute('class', "fribox-img img-circle img-responsive");
    img.setAttribute('alt', data.username + '-icon');

    img.onload = function() {
      img.height = img.width;
    }

    icon.append(img);

    var usr = document.createElement('div');
    usr.setAttribute('id', data.username + "-box-username");
    usr.setAttribute('class', "row fribox-username");

    var p1 = document.createElement('p');
    p1.setAttribute('class', 'text-muted text-center');
    p1.innerHTML = data.username;

    usr.append(p1);
    
    var nik = document.createElement('div');
    nik.setAttribute('id', data.username + "-box-nickname");
    nik.setAttribute('class', "row fribox-nickname");

    var p2 = document.createElement('p');
    p2.setAttribute('class', 'text-primary text-center');
    p2.innerHTML = data.nickname;

    nik.append(p2);

    content.append(icon, usr, nik);
    fribox.append(content);
    return fribox;
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


  function newMsgBox(ct, dir) {
    var row = document.createElement("div");
    row.setAttribute("class", "row");

    var outer = document.createElement("div");
    var temp = "";
    if (dir == "right") {
      outer.setAttribute("class", "col-xs-offset-4 col-xs-8");
    } else {
      outer.setAttribute("class", "col-xs-8");
    }
    var mid = document.createElement("p");
    mid.setAttribute("class", "text-" + dir + " my-msg");

    var sp = document.createElement("span");
    sp.innerHTML = ct;

    mid.append(sp);
    outer.append(mid);
    row.append(outer);

    return row;
  }




  function cre_b(elm, attr_class, attr_id, innerHTML) {
    var e = document.createElement(elm);
    if (attr_class != null) {
      e.setAttribute("class", attr_class);
    }
    if (attr_id != null) {
      e.setAttribute("id", attr_id);
    }
    if (innerHTML != null) {
      e.innerHTML = innerHTML;
    }
    return e;
  }

  function cre_img(attr_class, attr_id, attr_src, attr_alt) {
    var e = document.createElement("img");
    if (attr_class != null) {
      e.setAttribute("class", attr_class);
    }

    if (attr_id != null) {
      e.setAttribute("id", attr_id);
    }

    if(attr_src != null) {
      e.setAttribute("src", attr_src);
    }

    if (attr_alt != null) {
      e.setAttribute("alt", attr_alt);
    }
    return e;
  }

  function cre_input(attr_class, attr_id, attr_type, attr_ph, attr_val) {
    var e = document.createElement("input");
    if (attr_class != null) {
      e.setAttribute("class", attr_class);
    }
    if (attr_id != null) {
      e.setAttribute("id", attr_id);
    }
    if (attr_type != null) {
      e.setAttribute("type", attr_type);
    }
    if (attr_ph != null) {
      e.setAttribute("placeholder", attr_ph);
    }
    if (attr_val != null) {
      e.setAttribute("value", attr_val);
    }
    return e;
  }

  function newFriroom(data) {
    var b1 = cre_b("div", "app-unactive friroom", data.username + "-room");
    var b11 = cre_b("div", "container-fluid friwrapper");
    
    var b111 = cre_b("div", "row frimid");
    
    var b1111 = cre_b("div", "frinner col-sm-offset-2 col-sm-8 col-md-offset-3 col-md-6 col-lg-offset-4 col-lg-4");
    
    var b11111 = cre_b("div", "row friroom-brand", data.username + "-room-brand");
    
    var b111111 = cre_b("div", "friroom-icon col-xs-2", data.username + "-room-icon");
    var b1111111 = cre_img("img-circle img-responsive", null, data.icon, data.username + "-room-icon");
    b111111.append(b1111111);
    var b111112 = cre_b("div", "friroom-nick col-xs-8", data.username + "-room-nickname");
    var b1111121 = cre_b("p", "text-muted", null, data.nickname);
    b111112.append(b1111121);

    var b111113 = cre_b("div", "friroom-back col-xs-2", data.username + "-room-back");
    var b1111131 = cre_img("img-circle img-responsive", null, "/static/img/back.png", "back-icon");
    b111113.append(b1111131);
    b11111.append(b111111, b111112, b111113);

    var b11112 = cre_b("div", "row friroom-msglist", data.username + "-room-msglist");
    b1111.append(b11111, b11112);
    b111.append(b1111);

    var b112 = cre_b("div", "row friroom-edit col-sm-offset-2 col-sm-8 col-md-offset-3 col-md-6 col-lg-offset-4 col-lg-4", data.username + "-room-edit");
    var b1121 = cre_b("div", "row");
    var b11211 = cre_input("friroom-input col-xs-9", data.username + "-room-input", "text", "input something....");
    b11211.setAttribute("autocomplete", "off");
    var b11212 = cre_b("button", "friroom-btn col-xs-3 btn btn-default", data.username + "-room-btn", "Send");
    b1121.append(b11211, b11212);
    b112.append(b1121);

    b11.append(b111, b112);
    b1.append(b11);

    return b1;
  }

  function newFrinfo(data) {
    var b1 = cre_b("div", "app-unactive frinfo", data.username + "-info");
    var b11 = cre_b("div", "container-fluid");
    var b111 = cre_b("div", "row");
    var b1111 = cre_b("div", "col-sm-offset-2 col-sm-8 col-md-offset-3 col-md-6 col-lg-offset-4 col-lg-4");
    
    var b11111 = cre_b("div", "row frinfo-brand", data.username + "-info-brand");
    var b111111 = cre_b("div", "col-xs-2", data.username + "-info-icon");
    var b1111111 = cre_img("frinfo-img img-circle img-responsive", null, data.icon, data.username + "-info-icon");
    b111111.append(b1111111);
    var b111112 = cre_b("div", "col-xs-offset-8 col-xs-2", data.username + "-info-back");
    var b1111121 = cre_img("frinfo-img img-circle img-responsive", null, "/static/img/back.png", "back-icon");
    b111112.append(b1111121);
    b11111.append(b111111, b111112);

    var b11112 = cre_b("div", "row frinfo-content text-center", data.username + "-info-content");
    
    var b111121 = cre_b("div", "row", data.username + "-info-username");
    var b1111211 = cre_b("label", "col-sm-2 control-label", null, "Username");
    var b1111212 = cre_b("p", "col-sm-10", data.username + "-info-username-ct", data.username);
    b111121.append(b1111211);
    b111121.append(b1111212);

    var b111122 = cre_b("div", "row", data.username + "-info-nickname");
    var b1111221 = cre_b("label", "col-sm-2 control-label", null, "Nickname");
    var b1111222 = cre_b("p", "col-sm-10", data.nickname + "-info-nickname-ct", data.nickname);
    b111122.append(b1111221);
    b111122.append(b1111222);

    var b111123 = cre_b("div", "row", data.username + "-info-email");
    var b1111231 = cre_b("label", "col-sm-2 control-label", null, "Email");
    var b1111232 = cre_b("p", "col-sm-10", data.username + "-info-email-ct", data.email);
    b111123.append(b1111231);
    b111123.append(b1111232);

    var b111124 = cre_b("div", "row", data.username + "-info-motto");
    var b1111241 = cre_b("label", "col-sm-2 control-label", null, "Motto");
    var b1111242 = cre_b("p", "col-sm-10", data.username + "-info-motto-ct", data.motto);
    b111124.append(b1111241);
    b111124.append(b1111242);

    var b111125 = cre_b("div", "row", data.username + "-info-delete");
    var b1111251 = cre_b("button", "col-sm-offset-2 col-sm-8 btn btn-danger", data.username + "-info-delete-btn", "Delete");
    b111125.append(b1111251);
    b11112.append(b111121, b111122, b111123, b111124, b111125);
    
    b1111.append(b11111, b11112);
    b111.append(b1111);
    b11.append(b111);
    b1.append(b11);

    return b1;
  }


  // 对Date的扩展，将 Date 转化为指定格式的String
  // 月(M)、日(d)、小时(h)、分(m)、秒(s)、季度(q) 可以用 1-2 个占位符，
  // 年(y)可以用 1-4 个占位符，毫秒(S)只能用 1 个占位符(是 1-3 位的数字)
  // 例子：
  // (new Date()).Format("yyyy-MM-dd hh:mm:ss.S") ==> 2006-07-02 08:09:04.423
  // (new Date()).Format("yyyy-M-d h:m:s.S")      ==> 2006-7-2 8:9:4.18
  Date.prototype.Format = function (fmt) {
      var o = {
          "M+": this.getMonth() + 1, //月份
          "d+": this.getDate(), //日
          "h+": this.getHours(), //小时
          "m+": this.getMinutes(), //分
          "s+": this.getSeconds(), //秒
          "q+": Math.floor((this.getMonth() + 3) / 3), //季度
          "S": this.getMilliseconds() //毫秒
      };
      if (/(y+)/.test(fmt)) fmt = fmt.replace(RegExp.$1, (this.getFullYear() + "").substr(4 - RegExp.$1.length));
      for (var k in o)
      if (new RegExp("(" + k + ")").test(fmt)) fmt = fmt.replace(RegExp.$1, (RegExp.$1.length == 1) ? (o[k]) : (("00" + o[k]).substr(("" + o[k]).length)));
      return fmt;
  }


});//end jquery



