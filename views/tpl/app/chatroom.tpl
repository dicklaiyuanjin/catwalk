<div class="container-fluid">
  <div class="row">
    <div id="frimain" class="col-sm-offset-2 col-sm-8">
    
      {{range .Fiflist}}
      <div id="{{.Username}}-box" class="fribox col-xs-4 col-sm-3 col-md-2">
        <div id="{{.Username}}-box-content" class="fribox-content container-fluid">
          <div id="{{.Username}}-box-icon" class="row fribox-icon">
            <img src="{{.Icon}}" alt="{{.Username}}-icon" class="fribox-img img-circle img-responsive">
          </div>
  
          <div id="{{.Username}}-box-username" class="row fribox-username">
            <p class="text-muted text-center">{{.Username}}</p>
          </div>
  
          <div id="{{.Username}}-box-nickname" class="row fribox-nickname">
            <p class="text-primary text-center">{{.Nickname}}</p>
          </div>
        
        </div>
      </div>
      {{end}}
     


    </div>
  </div>
</div>
