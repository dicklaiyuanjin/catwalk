<div class="container-fluid">
  <div class="row">
    <div id="frimain" class="col-sm-offset-2 col-sm-8">
    
      {{range .Fm}}
      <div id="{{.Friusername}}-box" class="fribox col-xs-4 col-sm-3 col-md-2">
        <div id="{{.Friusername}}-box-content" class="fribox-content container-fluid">
          <div id="{{.Friusername}}-box-icon" class="row fribox-icon">
            <img src="{{.Friicon}}" alt="{{.Friusername}}-icon" class="fribox-img-rec fribox-img img-circle img-responsive">
          </div>
  
          <div id="{{.Friusername}}-box-username" class="row fribox-username">
            <p class="text-muted text-center">{{.Friusername}}</p>
          </div>
  
          <div id="{{.Friusername}}-box-nickname" class="row fribox-nickname">
            <p class="text-primary text-center">{{.Frinickname}}</p>
          </div>
        
        </div>
      </div>
      {{end}}
     


    </div>
  </div>
</div>
