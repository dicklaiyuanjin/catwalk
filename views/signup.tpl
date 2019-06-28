{{template "tpl/head.tpl"}}    
  {{template "tpl/signformhead.tpl"}}
    <div class="form-group">
      <label for="captchainput" class="col-sm-2 control-label">Captcha</label>
      <div class="col-sm-5">
        <input type="text" class="form-control" id="captchainput" placeholder="Captcha...">
      </div>
     <div class="col-sm-5">
        <img id="captcha" src="{{ .captcha }}" alt="Captcha" />
      </div>
    </div>
    <div class="form-group text-right">
      <div class="col-sm-offset-2 col-sm-10">
        <button type="submit" id="signupbtn" class="btn btn-success">Sign up</button>
  {{template "tpl/signformtail.tpl"}}
  <script src="/static/js/signup.js"></script>
{{template "tpl/tail.tpl"}}


