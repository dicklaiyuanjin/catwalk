<div class="container-fluid">
      <div class="row">
        <div class="col-md-4 col-md-offset-4 col-xs-10 col-xs-offset-1 col-sm-8 col-sm-offset-2 col-lg-4 col-lg-offset-4">
          <form name="signform" id="signform" class="form-horizontal">
            <div  class="form-group">
              <label for="username" class="col-sm-2 control-label">Username</label>
              <div class="col-sm-10">
                <input type="text" class="form-control" id="username" name="username" placeholder="Username...">
              </div>
            </div>
            <div class="form-group">
              <label for="password" class="col-sm-2 control-label">Password</label>
              <div class="col-sm-10">
                <input type="password" class="form-control" id="password" name="password" placeholder="Password...">
              </div>
            </div>
<div class="form-group">
      <label for="captchainput" class="col-sm-2 control-label">Captcha</label>
      <div class="col-sm-5">
        <input type="text" class="form-control" id="captchainput" placeholder="Captcha...">
      </div>
     <div class="col-sm-5">
        <img id="captcha" src="{{.captcha}}" alt="Captcha" />
      </div>
    </div>
    <div class="form-group text-right">
      <div class="col-sm-offset-2 col-sm-10">
        <button type="submit" id="{{.btnsignid}}" class="btn {{.btnclass}}">{{.btnvalue}}</button>
        </div>
            </div>
            <div class="form-group">
              <div class="col-sm-offset-2 col-sm-10">
                <em id="hint"></em>
                <em id="errmsg"></em>
              </div>
            </div>
          </form>
        </div>
      </div>
    </div>
<script src="/static/js/captcha.js"></script>
<script src="/static/js/sign.js"></script>
<script src="/static/js/{{.jsfile}}"></script>
