    </div><!-- end div#app -->
    <div id="app-room" class="apphide">
      {{template "tpl/app/room.tpl" .}}
    </div>
    <div id="app-frinfo" class="apphide">
      {{template "tpl/app/frinfo.tpl" .}}
    </div>

    <script src="/static/js/app/basic.js"></script>
    <script src="/static/js/app/setting.js"></script>
    <!-- jQuery (Bootstrap 的所有 JavaScript 插件都依赖 jQuery，所以必须放在前边) -->
    <script src="https://cdn.jsdelivr.net/npm/jquery@1.12.4/dist/jquery.min.js"></script>
    <!-- 加载 Bootstrap 的所有 JavaScript 插件。你也可以根据需要只加载单个插件。 -->
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@3.3.7/dist/js/bootstrap.min.js"></script>

    <script src="/static/js/app/app.js"></script>
  </body>
</html>

