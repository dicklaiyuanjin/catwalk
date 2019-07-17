{{range .Fm}}
<div id="{{.Friusername}}-info" class="app-unactive frinfo">
  <div class="container-fluid">
    <div class="row">
      <div class="col-sm-offset-2 col-sm-8 col-md-offset-3 col-md-6 col-lg-offset-4 col-lg-4">
        
        <div id="{{.Friusername}}-info-brand" class="row frinfo-brand">
          <div id="{{.Friusername}}-info-icon" class="col-xs-2">
            <img src="{{.Friicon}}" class="frinfo-img img-circle img-responsive" alt="{{.Friusername}}-info-icon">
          </div>
          <div id="{{.Friusername}}-info-back" class="col-xs-offset-8 col-xs-2">
            <img src="/static/img/back.png" class="frinfo-img img-circle img-responsive" alt="back-icon">
          </div>
        </div>

        <div id="{{.Friusername}}-info-content" class="row frinfo-content text-center">
          <div id="{{.Friusername}}-info-username" class="row">
            <label class="col-sm-2 control-label">Username</label>
            <p id="{{.Friusername}}-info-username-ct" class="col-sm-10">{{.Friusername}}</p>
          </div>

          <div id="{{.Friusername}}-info-nickname" class="row">
            <label class="col-sm-2 control-label">Nickname</label>
            <p id="{{.Friusername}}-info-nickname-ct" class="col-sm-10">{{.Frinickname}}</p>
          </div>

          <div id="{{.Friusername}}-info-email" class="row">
            <label class="col-sm-2 control-label">Email</label>
            <p id="{{.Friusername}}-info-email-ct" class="col-sm-10">{{.Friemail}}</p>
          </div>

          <div id="{{.Friusername}}-info-motto" class="row">
            <label class="col-sm-2 control-label">Motto</label>
            <p id="{{.Friusername}}-info-motto-ct" class="col-sm-10">{{.Frimotto}}</p>
          </div>

          <div id="{{.Friusername}}-info-delete" class="row">
            <button id="{{.Friusername}}-info-delete-btn" class="col-sm-offset-2 col-sm-8 btn btn-danger">Delete</button>
          </div>
        </div>
      
      </div>
    </div>
  </div>
</div>
{{end}}
