function ajaxCaptcha() {
  var xhr = new XMLHttpRequest();
  
  xhr.onreadystatechange = function() {
    if(xhr.readyState === 4) {
      if(xhr.status === 200) {
        var url = JSON.parse(xhr.responseText);
        setCaptcha(url);
      } else {
        console.error("err: ", xhr.statusText);
      }
    }
  };

  xhr.onerror = function (e) {
    console.error(xhr.statusText);
  };
  
  var url = document.location.protocol + "//" + document.location.host + "/auth/captcha";
  xhr.open('POST', url, true);
  xhr.send();
}


function setCaptcha(url) {
  var img = document.getElementById("captcha");
  img.setAttribute('src', url);
}

var img = document.getElementById("captcha");
img.addEventListener('click', ajaxCaptcha, false);
