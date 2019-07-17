{{range .Fm}}
<div id="{{.Friusername}}-room" class="app-active friroom">
  <div class="container-fluid friwrapper">
    <div class="row frimid">
      <div class="frinner col-sm-offset-2 col-sm-8 col-md-offset-3 col-md-6 col-lg-offset-4 col-lg-4">
        
        <div id="{{.Friusername}}-room-brand" class="row friroom-brand">
          <div id="{{.Friusername}}-room-icon" class="friroom-icon col-xs-2">
            <img src="{{.Friicon}}" class="img-circle img-responsive" alt="{{.Friusername}}-room-icon"> 
          </div>

          <div id="{{.Friusername}}-room-nickname" class="friroom-nick col-xs-8">
            <p class="text-muted">{{.Frinickname}}</p>
          </div>

          <div id="{{.Friusername}}-room-back" class="friroom-back col-xs-2">
            <img src="/static/img/back.png" class="img-circle img-responsive" alt="back-icon">
          </div>
        </div>

        <div id="{{.Friusername}}-room-msglist" class="row friroom-msglist">
          {{$Fri := .Friusername}}
          {{range .Msg}}
            {{if ne $Fri .Sender}}
              <div class="row">
                <div class="col-xs-offset-4 col-xs-8">
                  <p class="text-right my-msg">
                    <span>{{.Content}}</span>
                  </p>
                </div>
              </div>
            {{else}} 
              <div class="row">
                <div class="col-xs-8">
                  <p class="text-left fri-msg">
                    <span>{{.Content}}</span>
                  </p>
                </div>
              </div>
            {{end}}{{/*end if*/}}

          {{end}}{{/*end range .Msg*/}}
        </div> 
      </div>
    </div>


    <div id="{{.Friusername}}-room-edit" class="row friroom-edit col-sm-offset-2 col-sm-8 col-md-offset-3 col-md-6 col-lg-offset-4 col-lg-4">
      <div class="row">
        <input type="text" id="{{.Friusername}}-room-input" class="friroom-input col-xs-9" placeholder="input something....">
        <button id="{{.Friusername}}-room-btn" class="friroom-btn col-xs-3 btn btn-default">Send</button>
      </div>
    </div>

  </div>
</div>
{{end}}{{/*end range .Fm*/}}


