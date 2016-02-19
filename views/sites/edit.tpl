<div class="bs-example" data-example-id="textarea-form-control">
    <form method="post" role="form" data-toggle="validator">
        <div class="input-group">
            <div class="input-group-addon">http://</div>
            <div class="form-group">
                <input {{ .Disabled }} id="domain" type="text" data-placement="bottom" title="必须有'字母','数字','-','_'构成" data-toggle="tooltip"
                       class="form-control" pattern="[a-zA-Z0-9-_]+" placeholder="site name" value="{{ .Site.Domain }}"
                       name="domain" required>
            </div>
            <div class="input-group-addon">.{{ $.SiteDomain }}</div>

            <div class="input-group-addon">➩➩➩</div>
            <div class="input-group-addon">http://</div>
            <div class="form-group">
                <input class="form-control" data-placement="bottom" title="必须填入合法IP的地址" data-toggle="tooltip"
                       placeholder="target host" value="{{ .Site.Host }}" name="host" type="text" required
                       pattern="\d+(\.\d+){3}">
            </div>

            <div class="input-group-addon">:</div>
            <div class="form-group">
                <input type="number" data-placement="bottom" data-toggle="tooltip" title="必须填入合法的端口"
                       class="form-control" placeholder="target port" value="{{ .Site.Port }}" name="port" required
                       max="65535" min="1">
            </div>
        </div>

        <div class="form-group">
            <label for="rootPath">rootPath</label>
            <div id="rootPath" class="form-group">
                <input data-placement="bottom" name="defaultPath" data-toggle="tooltip" title="首页路径"
                          class="form-control" placeholder="首页" value="{{ .Site.DefaultPath }}" />
            </div>
        </div>

        <div class="form-group">
            <label for="exampleInputPassword1">描述</label>
            <div id="exampleInputPassword1" class="form-group">
                <textarea required data-placement="bottom" name="description" data-toggle="tooltip" title="对此服务的描述"
                          class="form-control" rows="3" placeholder="用途/背景/负责人等...">{{ .Site.Description }}</textarea>
            </div>
        </div>
        <button type="submit" class="btn btn-default">Submit</button>
    </form>
</div><!-- /.bs-example -->