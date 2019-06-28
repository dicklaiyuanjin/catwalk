{{template "tpl/head.tpl"}}
  {{template "tpl/signformhead.tpl"}}    
    <div class="form-group">
      <label for="captchainput" class="col-sm-2 control-label">Captcha</label>
      <div class="col-sm-5">
        <input type="text" class="form-control" id="captchainput" name="captcha" placeholder="Captcha...">
      </div>
      <div class="col-sm-5">
        <img id="captcha" src="{{.captcha}}" alt="Captcha" />
      </div>
    </div>
    <div class="form-group text-right">
       <div class="col-sm-offset-2 col-sm-10">
         <button type="submit" id="signinbtn" class="btn btn-primary">Sign in</button>
  {{template "tpl/signformtail.tpl"}}
  <script src="/static/js/signin.js"></script>
{{template "tpl/tail.tpl"}}
