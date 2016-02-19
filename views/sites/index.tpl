<div class="form-group has-success has-feedback">
    <label class="control-label sr-only" for="inputGroupSuccess4">Input group with success</label>
    <div class="input-group">
        <span class="input-group-addon">搜索</span>
        <input type="text" placeholder="根据域名,IP,端口,描述进行搜索" class="form-control" id="searchBox"
               aria-describedby="inputGroupSuccess4Status">
    </div>
</div>


<a class="btn btn-primary btn-success" href="/sites/new">新建</a>

<table class="table table-hover">
    <thead>
    <th class="text-center" width="15%">*</th>
    <th class="text-center" width="10%">子域名</th>
    <th class="text-center" width="25%">链接</th>
    <th width="20%">目标地址</th>
    <th>描述</th>
    </thead>
    <tbody>
    {{ range .Sites }}
    <tr>
        <td class="text-center">
            <form method="post" action="/sites/{{ .Domain }}/delete">
                <a class="btn btn-primary btn-sm btn-warning" href="/sites/{{ .Domain}}/">编辑</a>
                <input type="submit" class="btn btn-primary btn-sm btn-danger" value="删除" />
            </form>
        </td>
        <td class="text-center">{{ .Domain }} </td>
        <td><a href="http://{{ .Domain }}.{{ $.SiteDomain }}/{{ .DefaultPath }}">http://{{ .Domain }}.{{ $.SiteDomain }}/{{ .DefaultPath }}</a></td>
        <td>http://{{ .Host }}:{{ .Port }}/{{ .DefaultPath }}</td>
        <td>
            <div>{{ .Description }}</div>
        </td>
    </tr>
    {{ end }}
    </tbody>
</table>