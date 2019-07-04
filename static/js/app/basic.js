/*
 * ajax request
 */
function ajax_request(url, method, jsondata, su_func) {
  var xhr = new XMLHttpRequest();
  xhr.onreadystatechange = function(){
    // 通信成功时，状态值为4
    if (xhr.readyState === 4){
      if (xhr.status === 200){
        su_func(xhr);
      } else {
        console.error(xhr.statusText);
      }
    }
  };
  
  xhr.onerror = function (e) {
    console.error(xhr.statusText);
  };
  
  xhr.open(method, url, true);
  xhr.setRequestHeader("Content-Type", "application/json;charset=utf-8");
  xhr.send(jsondata);
}



