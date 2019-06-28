function validate_userinfo(usr, pwd, cit) {
  var hint = document.getElementById("hint");
  hint.innerHTML = "";
  var err = document.getElementById("errmsg");
  err.innerHTML = "";
  var errmsg;
  errmsg = "";
  
  if (usr.length <= 0) {
    errmsg = errmsg.concat("username is required...<br/>");
  }

  if (usr.length > 20) {
    errmsg = errmsg.concat("username is too long(maxlength is 20)...<br/>");
  }

  if (pwd.length < 8) {
    errmsg = errmsg.concat("password is too short(minlength is 8)...</br>");
  }

  if (pwd.length > 20) {
    errmsg = errmsg.concat("password is too long(maxlength is 20)...<br/>");
  }

  if (cit.length <= 0) {
     errmsg = errmsg.concat("captcha is required...<br/>");
  }

  console.log("errmsg: ", errmsg);
  if (errmsg.length === 0) {
    return true;
  } else {
    err.innerHTML = errmsg;
    return false;
  }
}

/*
 * string action: signin or signup 
 */
function ajaxSign(action) {
  var usr = document.getElementById("username").value;
  var pwd = document.getElementById("password").value;
  var cit = document.getElementById("captchainput").value;
  if (validate_userinfo(usr, pwd, cit) == true)  {
    var userjson = JSON.stringify({
      "username": usr,
      "password": pwd,
      "captchainput": cit
    });
  
    var url = document.location.protocol + "//" + document.location.host + "/auth/" + action;
    var xhr = new XMLHttpRequest();
    xhr.open('POST', url, true);
    xhr.setRequestHeader("Content-Type", "application/json;charset=utf-8");
    xhr.send(userjson);
  
    xhr.onreadystatechange = function() {
      var hint = document.getElementById('hint');
      if(xhr.readyState === 4) {
        if(xhr.status === 200 || xhr.status === 302) { 
          var obj = JSON.parse(xhr.responseText);
          switch (obj.state) {
            case 0:
              ajaxCaptcha();
              hint.innerHTML = 'fail, please ensure your input...';
              break;
            case 1:
              hint.innerHTML = 'Success, going to app...';
              document.location.href = document.location.protocol + "//" + document.location.host + "/app";
              break;
          } 
          
          
        } else {
          hint.innerHTML = 'Fail, Server error';
        }
  
      } else {
        hint.innerHTML = 'Loading...'
      }
    };
  }
}

/*
 * action：signin or signup
 */
function signListener(action) {
  var form = document.getElementById("signform");
  form.addEventListener("submit", function(event){
    event.preventDefault();
    ajaxSign(action);
  });
}
