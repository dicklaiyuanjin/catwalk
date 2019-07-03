<div class="container-fluid">
  <div class="row">
    <div id="setting-content" class="col-xs-offset-1 col-xs-10 col-sm-offset-3 col-sm-6">
      <form class="form-horizontal" id="setting-form">
        <div class="form-group">
          <div class="col-xs-offset-4 col-xs-4 col-sm-offset-4 col-sm-4">
            <img id="icon-img" src="{{.icon}}" class="img-responsive" alt="user icon" >
          </div>
        </div>

        <div class="form-group">
          <div class="col-xs-offset-3 col-xs-6 col-sm-offset-4 col-sm-4">
            <label id="uploadbtn" class="btn btn-info widthbtn" for="icon-user">upload</label>
          </div>
          <input type="file" id="icon-user" accept="image/png, image/jpeg" style="position:absolute;clip:rect(0 0 0 0);">
        </div>

        <div class="form-group">
          <label for="username" class="col-sm-2 control-label">username</label>
          <div class="col-sm-10">
            <input type="text" class="form-control" id="username" readonly="readonly" value="{{.userinfo.Username}}">
          </div>
        </div>
        
        <div class="form-group">
          <label for="nickname" class="col-sm-2 control-label">nickname</label>
          <div class="col-sm-10">
            <input type="text" class="form-control" id="nickname" readonly="readonly" value="{{.userinfo.Nickname}}">
          </div>
        </div>
        
        <div class="form-group">
          <label for="email" class="col-sm-2 control-label">e-mail</label>
          <div class="col-sm-10">
            <input type="text" class="form-control" id="email" readonly="readonly" value="{{.userinfo.Email}}">
          </div>
        </div>
        
        <div class="form-group">
          <label for="motto" class="col-sm-2 control-label">motto</label>
          <div class="col-sm-10">
            <input type="text" class="form-control" id="motto" readonly="readonly" value="{{.userinfo.Motto}}">
          </div>
        </div>

        <div class="form-group">
          <div class="col-sm-12">
            <button id="editbtn" type="button" class="widthbtn btn btn-info">Edit</button>
          </div>
        </div>
        
        <div class="form-group">
          <div class="col-sm-12">
            <button id="signoutbtn" type="button" class="widthbtn btn btn-info">Sign out</button>
          </div>
        </div>
        <em id="hint"></em> 
      </form> 
    </div>
  </div>
</div>
