var obJoke = (function() {
    function JokeWindow() {
        // 变量定义
        var self = this;
        // 绑定登入信息
        $("#butLogin").bind("click", function() { self.doLogin(); });
    }

    var _proto = JokeWindow.prototype;

    _proto.init = function() {
        // 判断当前是否有在登入状态
        this.send("checklogin");
    }

    _proto.showLogin = function() {
        $("#myModal").modal({
            backdrop: "static",
            keyboard: false
        });
    }

    _proto.doLogin = function() {
		var loginPass = $("#loginpass").val()
		this.send("login " + loginPass)
    }
	
	_proto.getRndJoke = function() {
		this.send("joke rnd")
	}

    _proto.send = function(cmd) {
        var self = this;
        $.ajax({
            url: "/cmds?cmd=" + cmd,
            success: function(result) {
                self.resultCmd(result);
            }
        })
    }

    _proto.resultCmd = function(result) {
        console.log(result, this);
        switch(result.Cmd) {
            case "LoginOut":
            this.showLogin();
            break;
			
			case "LOGIN":
			var info = result.Info;
			if (info.result != true) {
				console.log(info.info);
			} else {
				// 隐藏登入框
				$("#myModal").modal("hide");
				// 获取一个笑话对象
				this.getRndJoke();
			}
			break;
        }
    }

    return new JokeWindow();
}());

(function() {
    // obJoke.showLogin();
    obJoke.init();
})()
