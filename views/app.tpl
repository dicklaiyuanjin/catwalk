{{template "tpl/app/head.tpl" .}}
  <div class="tab-content">
    <div role="tabpanel" class="tab-pane active" id="chatroom">{{template "tpl/app/chatroom.tpl" .}}</div>
    <div role="tabpanel" class="tab-pane" id="invitation">{{template "tpl/app/invitation.tpl" .}}</div>
    <div role="tabpanel" class="tab-pane" id="setting">{{template "tpl/app/setting.tpl" .}}</div>
  </div>
{{template "tpl/app/tail.tpl" .}}
