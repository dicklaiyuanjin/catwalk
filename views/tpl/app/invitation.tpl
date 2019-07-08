<div class="container-fluid">
  <div class="row">
    <div class="col-xs-offset-1 col-xs-10 col-sm-offset-3 col-sm-6">
      <ul class="list-group">
        <li class="list-group-item">
          <span class="glyphicon glyphicon-pencil" aria-hidden="true"></span> Invitation Envelope
        </li>
        <li class="list-group-item">  
          <div class="form-horizontal" id="invite-envelope">
            <div class="form-group">
              <label for="receiver-name" class="col-sm-2 control-label">Name</label>
              <div class="col-sm-10">
                <input type="text" class="form-control" id="receiver-name" placeholder="Receiver's Name">
              </div>
            </div>
            <div class="form-group">
              <label for="invite-message" class="col-sm-2 control-label">Message</label>
              <div class="col-sm-10">
                <input type="text" class="form-control" id="invite-message" placeholder="invitation message">
              </div>
            </div>
            <div class="form-group">
              <div class="col-sm-offset-2 col-sm-10">
                <button type="submit" class="widthbtn btn btn-info" id="invite-send">Send</button>
              </div>
            </div>
          </div>
        </li>

      </ul>
    </div>  
  </div>

  <div class="row">
    <div class="col-xs-offset-1 col-xs-10 col-sm-offset-3 col-sm-6">
      <ul id="rec-envelope" class="list-group">        
        {{range .Invitations}}
        <div id="{{.Sender}}-invite">
          <li id="{{.Sender}}-invite-brand" class="list-group-item">
            <span class="glyphicon glyphicon-envelope" aria-hidden="true"></span> <span id="{{.Sender}}">{{.Sender}}</span> <span class="myarrow glyphicon glyphicon-arrow-down" aria-hidden="true"></span>
          </li>
          
          <li class="list-group-item" id="{{.Sender}}-invite-content">
            <div class="form-horizontal" id="{{.Sender}}-inviteform"> 
              <div class="form-group">
                <label class="col-sm-2 control-label">Message</label>
                <div class="col-sm-10 padtop7px">
                  <p type="text" id="{{.Sender}}-invitemsg">{{.Msg}}</p>
                </div>
              </div>
              <div class="form-group">
                <div class="col-sm-offset-2 col-sm-10">
                  <button type="submit" class="widthbtn btn btn-primary" id="{{.Sender}}-agree">Agree</button>
                </div>
              </div>
              <div class="form-group">
                <div class="col-sm-offset-2 col-sm-10">
                  <button type="submit" class="widthbtn btn btn-success" id="{{.Sender}}-refuse">Refuse</button>
                </div>
              </div>
            </div>
          </li>
        </div>
        {{end}}
      </ul>
    </div>
  </div>
</div>
