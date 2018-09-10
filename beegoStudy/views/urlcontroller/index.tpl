<h5>OKOKOKOK!!!!!</h5>
<p>Controller中传过来的username为：{{.username}}</p>
<p>{{urlfor "UrlController.List"}}</p> <!--urlfor不能大写，只能小写；多个参数用空格隔开-->
<p>{{urlfor "UrlController.Get" ":last" "hai" ":first" "ou"}}</p> 
<!--有bug：这里并不能转换为：/person/hai/ou ；而是：/url/get?:first=ou&:last=hai 。 后面又可以了，随便修改过router.go就可以了，貌似编译并不是很靠谱-->
<p>{{urlfor "UrlController.GetUrl"}}</p> <!--/url/geturl--> 
<p>{{urlfor "UrlController.Myext"}}</p>
<p>{{.Ext}}</p> <!--大小写是有区别的，需要和controller中传的参数名一致-->
<p>{{.XXX}}</p> <!-- 在模板中如果引用了不存在的变量并不会报错 -->
<p>{{.PersonUrl}}</p>