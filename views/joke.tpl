<!DOCTYPE html>
<html>
<head>
	<meta charset="utf-8">
    <title>笑话过审系统</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
	<meta http-equiv="Pragma" content="no-cache">
	<meta http-equiv="Cache-control" content="no-cache">
	<meta http-equiv="Cache" content="no-cache">
    <!-- Bootstrap -->
    <link rel="stylesheet" href="/static/css/bootstrap.min.css">  
	
<style>
 body {
	/* padding-top: 50px; */
	/* background: #f5f5f5; */
    /* padding-left: 50px; */
}
</style>
<!--[if lt IE 9]>
	<script src="https://apps.bdimg.com/libs/html5shiv/3.7/html5shiv.min.js"></script>
<![endif]-->
</head>
<body>


<div class="navbar navbar-fixed-top navbar-inverse" role="navigation">
	<div class="container">
		<div class="navbar-header">
			<button type="button" class="navbar-toggle" data-toggle="collapse" data-target=".navbar-collapse">
				<span class="sr-only">Toggle navigation</span>
				<span class="icon-bar"></span>
				<span class="icon-bar"></span>
				<span class="icon-bar"></span>
			</button>
			<a class="navbar-brand" href="#">笑话过审系统<small style="font-size:8px;color:#838B8B"> Ver0.01 by vava6.com</small></a>
		</div>
		<div class="collapse navbar-collapse">
			<ul class="nav navbar-nav">
				<li class="active"><a href="#"><i class="glyphicon glyphicon-home"></i> 首页</a></li>
				<li><a href="#shop"><i class="glyphicon glyphicon-flag"></i> 标签</a></li>
				<li><a href="#support"><i class="glyphicon glyphicon-log-out"></i> 退出</a></li>
			</ul>
		</div>
	</div>
</div>
<!-- display:none;-->
<div class="container" style="position:fixed;width:100%;">
	<div class="alert alert-success" id="msgBox">
		<a href="#" class="close" data-dismiss="alert">
			&times;
		</a>
		<strong>提示！</strong><span id="msgBoxInfo">您的网络连接有问题。</span>
	</div>
</div>

<div class="container" style="margin-top:60px;margin-bottom:10px;min-height:100%;overflow:scroll;overflow-x:visible;overflow-y:visible">
	<p style="margin-top:15px;font-size:17px" id="jokeContent"></p>
	<hr />
</div>

<div class="container">
	<input type="hidden" id="jokeID" value="" />
 	<h3><i class="glyphicon glyphicon-edit"></i> 请评审</h3>
  	<div class="btn-group btn-group-justified">
  	  	<div class="btn-group"><button type="button" class="btn btn-primary" id="butPass"><i class="glyphicon glyphicon-remove"></i> 不过审</button></div>
    	<div class="btn-group"><button type="button" class="btn btn-primary" id="butOk"><i class="glyphicon glyphicon-ok"></i> 过审</button></div>
    	<div class="btn-group"><button type="button" class="btn btn-primary" id="butGood"><i class="glyphicon glyphicon-thumbs-up"></i> 精典</button></div>
  	</div>
	<p class="text-info" style="margin-top:5px;">&nbsp;好评&nbsp;<span style="color:#3c763d;" id="jokeVote">0</span>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;评论&nbsp;<span style="color:#3c763d" id="jokeComment">0</span></p>
	<!-- <p class="text-success">本行内容带有一个 success class</p> -->
	<!-- 337ab7
	<blockquote>
		这是一个带有源标题的引用。 primary
		<small>Someone famous in <cite title="Source Title">Source Title</cite></small>
	</blockquote>
	-->
		
	<div style="margin-top:10px;margin-right:10px;margin-bottom:100px;margin-left:10px;line-height:2">
		<span class="label label-default" value="2" id="jokeType02">朋友闺蜜</span>
		<span class="label label-default" value="3" id="jokeType03">夫妻情侣</span>
		<span class="label label-default" value="4" id="jokeType04">生活家庭</span>
		<span class="label label-default" value="5" id="jokeType05">职场工作</span>
		<span class="label label-default" value="6" id="jokeType06">校园学习</span>
		<span class="label label-default" value="1" id="jokeType01">其它类型</span>
	</div>
</div>

<!-- 模态框（Modal）aria-hidden="true" -->
<div class="modal fade" id="myModal" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true" style="margin-top: 50px;">
	<div class="modal-dialog">
		<div class="modal-content">
			<div class="modal-header">
				<!--
				<button type="button" class="close" data-dismiss="modal" aria-hidden="true">
					&times;
				</button>
				-->
				<h4 class="modal-title" id="myModalLabel">
					登入
				</h4>
			</div>
			<div class="modal-body">
				<div class="form-group">
					<div class="col-sm-12">
						<input type="password" class="form-control" id="loginpass" placeholder="请输入登入码" />
					</div>
				</div>
			</div>
			<div class="modal-footer">
				<button type="button" class="btn btn-primary" id="butLogin">
					登入系统
				</button>
			</div>
		</div><!-- /.modal-content -->
	</div><!-- /.modal -->
</div>

<!-- jQuery (Bootstrap 插件需要引入) -->
<script src="/static/js/jquery.min.js"></script>
<!-- 包含了所有编译插件 -->
<script src="/static/js/bootstrap.min.js"></script>

<script src="/static/js/webjokes.js"></script>

</body>
</html>
