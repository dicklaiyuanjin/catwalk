/*
 * edit btn
 */
function edit_su_func(xhr){
  console.log(xhr);
  var obj = JSON.parse(xhr.responseText);
  switch (obj.state) {
    case 0:
      var hint = document.getElementById("hint");
      hint.innerHTML = 'edit fail, please ask for administrator...';
      break;
    case 1:
      break;
  }
}

function editbtn_click_handler() {
  var editbtn = document.getElementById("editbtn");

  var nick = document.getElementById("nickname");
  var email = document.getElementById("email");
  var motto = document.getElementById("motto");

  var readonly2 = email.getAttribute('readonly');
  var readonly3 = motto.getAttribute('readonly');
  var readonly4 = nick.getAttribute('readonly');


  if (readonly2 != null || readonly3 != null || readonly4 != null) {
    nick.removeAttribute('readonly');
    email.removeAttribute('readonly');
    motto.removeAttribute('readonly');
  
    editbtn.innerHTML = 'Submit';
    editbtn.setAttribute('class', 'widthbtn btn btn-default');
  
  } else {
    nick.setAttribute('readonly', 'readonly');
    email.setAttribute('readonly', 'readonly');
    motto.setAttribute('readonly', 'readonly');
    
    editbtn.innerHTML = 'Edit';
    editbtn.setAttribute('class', 'widthbtn btn btn-info');
    
    var val_usr = document.getElementById("username").value;
    var val_nick = document.getElementById("nickname").value;
    var val_email = document.getElementById("email").value;
    var val_motto = document.getElementById("motto").value;
    var jsondata = {
      username: val_usr,
      nickname: val_nick,
      email: val_email,
      motto: val_motto,
      icon: ""
    };

    ajax_request('/app/edit', 'POST', JSON.stringify(jsondata), edit_su_func);
  
  }

}

var editbtn = document.getElementById("editbtn");
editbtn.addEventListener('click', editbtn_click_handler , false);


/*
 * signout btn
 */
function signout_su_func(xhr) {
  var obj = JSON.parse(xhr.responseText);
  switch (obj.state) {
    case 0:
      var hint = document.getElementById("hint");
      hint.innerHTML = 'sign out fail, please ask for administrator...';
      break;
    case 1:
      document.location.href = "/";
      break;
  }
}

function signoutbtn_click_handler() {
  ajax_request('/app/signout','GET', null, signout_su_func);
}

var signoutbtn = document.getElementById("signoutbtn");
signoutbtn.addEventListener('click', signoutbtn_click_handler, false);




/*
 * upload btn
 */
function upload_su_func(xhr) {
  console.log(xhr);
  var obj = JSON.parse(xhr.responseText);
  switch (obj.state) {
    case 0:
      var hint = document.getElementById("hint");
      hint.innerHTML = 'upload fail, please ask for administrator...';
      break;
    case 1:
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
    cvs.width = sp.width;
    cvs.height = sp.width;
    ctx.drawImage(sp, 0, 0, cvs.width, cvs.height);
    var base64 = cvs.toDataURL("image/jpeg", 0.5);
    var jsondata = JSON.stringify({
      "username": username,
      "nickname": "",
      "email": "",
      "motto": "",
      "icon": base64
    });
    console.log(jsondata);

    ajax_request('/app/upload', 'POST', jsondata, upload_su_func);
  }
}

var imgInput = document.getElementById("icon-user");
imgInput.addEventListener('change', upload_handler);
