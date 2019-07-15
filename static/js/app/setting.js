/*
 * edit btn
 */
var rawData =  {
  username: document.getElementById("username").value,
  nickname: document.getElementById("nickname").value,
  email: document.getElementById("email").value,
  motto: document.getElementById("motto").value,
  icon: ""
};

function pushRawData (data) {
  var nickname = document.getElementById("nickname");
  var email = document.getElementById("email");
  var motto = document.getElementById("motto");

  nickname.value = data.nickname;
  email.value = data.email;
  motto.value = data.motto;
}

function isReadonly(arr) {
  flag = true;
  for (var i in arr) {
    if(arr[i].getAttribute('readonly') == null) {
      flag = false;
    }
  }
  return flag;
}

function removeAttrReadonly(arr) {
  for(var i in arr) {
    arr[i].removeAttribute('readonly');
  }
}

function setAttrReadonly(arr) {
  for(var i in arr) {
    arr[i].setAttribute('readonly', 'readonly');
  }
}

function getJsondata() {
  return {
    username: document.getElementById("username").value,
    nickname: document.getElementById("nickname").value,
    email: document.getElementById("email").value,
    motto: document.getElementById("motto").value,
    icon: ""
  };
}

function setEditbtn(btn, html, val){
  btn.innerHTML = html;
  btn.setAttribute('class', val);
}

function sameAsBefore(data) {
  if(data.nickname == rawData.nickname && 
     data.email == rawData.email && 
     data.motto == rawData.motto) {
    return true;
  }
  return false;
}

function editbtn_click_handler() {
  var editbtn = document.getElementById("editbtn");
  
  var arr = [];
  arr[0] = document.getElementById("nickname");
  arr[1] = document.getElementById("email");
  arr[2] = document.getElementById("motto");

  if (isReadonly(arr) == true) {
    removeAttrReadonly(arr);
    setEditbtn(editbtn, 'Submit', 'widthbtn btn btn-default');
  } else {
    setAttrReadonly(arr);
    setEditbtn(editbtn, 'Edit', 'widthbtn btn btn-info'); 
    var jsondata = getJsondata();
    if(sameAsBefore(jsondata) == false) {
      ajax_request('/app/setting/edit', 'POST', JSON.stringify(jsondata), function(xhr){
        var obj = JSON.parse(xhr.responseText);
        var hint = document.getElementById("hint");
        var errmsg = ""; 

        if (obj.existnick == 1 && rawData.nickname != jsondata.nickname) {
          errmsg += 'nickname is exist, please change it.';
        } else {
          rawData.nickname = document.getElementById("nickname").value;
        }

        if (obj.existemail == 1 && rawData.email != jsondata.email) {
          errmsg += 'email is exist, please change it.';
        } else {
          rawData.email = document.getElementById("email").value;
        }

        rawData.motto = document.getElementById("motto").value;
        pushRawData(rawData);
        hint.innerHTML = errmsg;
      });
    }
  }
}

var editbtn = document.getElementById("editbtn");
editbtn.addEventListener('click', editbtn_click_handler , false);


/*
 * signout btn
 */
function signout_su_func(xhr) {
  document.location.href = "/"; 
}

function signoutbtn_click_handler() {
  ajax_request('/app/setting/signout','GET', null, signout_su_func);
}

var signoutbtn = document.getElementById("signoutbtn");
signoutbtn.addEventListener('click', signoutbtn_click_handler, false);




/*
 * upload btn
 */
function upload_su_func(xhr) {
  var obj = JSON.parse(xhr.responseText);
  var hint = document.getElementById("hint");
  switch (obj.state) {
    case 0:
      hint.innerHTML = 'upload fail, please ask for administrator...';
      break;
    case 1:
      hint.innerHTML = 'upload success';
      break;
  }
}

function upload_handler() {
  var username = document.getElementById("username").value;
  var icon = document.getElementById("icon-user");
  var sp = document.getElementById("icon-img");
  var f = window.URL.createObjectURL(icon.files[0]);
  sp.src = f;

  sp.onload = function() {
    var cvs = document.createElement("canvas");
    var ctx = cvs.getContext('2d');
    cvs.width = 125;
    cvs.height = 125;
    ctx.drawImage(sp, 0, 0, cvs.width, cvs.height);
    var base64 = cvs.toDataURL("image/jpeg", 0.5);
    var jsondata = JSON.stringify({
      "username": username,
      "nickname": "",
      "email": "",
      "motto": "",
      "icon": base64
    });

    ajax_request('/app/setting/upload', 'POST', jsondata, upload_su_func);
  }
}

var imgInput = document.getElementById("icon-user");
imgInput.addEventListener('change', upload_handler);
