var obJoke = (function() {
    function JokeWindow() {
        // 变量定义
        var self = this;
        // 绑定登入信息
        $("#butLogin").bind("click", function() { self.doLogin(); });
        // 评审不通过
        $("#butPass").bind("click", function() { self.doCheck(1); });
        // 评审通过
        $("#butOk").bind("click", function() { self.doCheck(2); });
        // 评审为精避
        $("#butGood").bind("click", function() { self.doCheck(3); });

        this.arrButLabel = [];
        for (var i = 1; i <= 6; i++) {
            var but = $("#jokeType0" + i);
            but.bind("click", function() {
                self.doSelectLabel(this);
            });
            this.arrButLabel.push(but);
        }
    }

    var _proto = JokeWindow.prototype;

    _proto.init = function() {
        // 判断当前是否有在登入状态
        this.getRndJoke();
    }

    _proto.getLabelColor = function(id) {
        switch(id) {
            case "1": return "label label-primary";
            case "2": return "label label-success";
            case "3": return "label label-info";
            case "4": return "label label-warning";
            case "5": return "label label-danger";
            case "6": return "label label-primary";
            default:
                return "label label-primary";
        }
    }

    _proto.showLogin = function() {
        $("#myModal").modal({
            backdrop: "static",
            keyboard: false
        });
    }

    _proto.doSelectLabel = function(nodeLabel) {
        var ob = $("#" + nodeLabel.id);
        if (ob.attr("class") == "label label-default") {
            var className = this.getLabelColor(ob.attr("value"));
            ob.attr("class", className);
            // 设定标签值
        } else {
            ob.attr("class", "label label-default");
        }
    }

    _proto.getLabelVal = function() {
        var arr = [];
        for (var i = 0; i < this.arrButLabel.length; i++) {
            var obBut = this.arrButLabel[i];
            if (obBut.attr("class") != "label label-default") {
                arr.push(obBut.attr("value"));
            }
        }
        return arr.join(",");
    }

    _proto.doCheck = function(state) {
        var cmd = "joke check " + $("#jokeID").val() + " " + state + " " + this.getLabelVal();
        // this.send(cmd);
        this.showMsg("这只是个测试!");
    }

    _proto.showJokeInfo = function(dbInfo) {
        $("#jokeID").val(dbInfo.id);
        $("#jokeContent").html(dbInfo.content);
        $("#jokeVote").text(dbInfo.vote);
        $("#jokeComment").text(dbInfo.comment);
        // 清空标签
        this.clearLabel();
    }

    _proto.clearLabel = function() {
        for (var i = 0; i < this.arrButLabel.length; i++) {
            var but = this.arrButLabel[i];
            but.attr("class", "label label-default");
        }
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

    _proto.showMsg = function(strMsg) {
        // $("#msgBox").attr("class", "alert alert-success");
        // $("#msgBoxInfo").html(strMsg);
        $("msgBox").show();
    }

    _proto.resultCmd = function(result) {
        console.log(result);
        switch(result.Cmd) {

            // 没有登入显示登入窗口
            case "LoginOut":
            this.showLogin();
            break;

            // 登入消息返回时执行
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

            // 获得笑话对象执行
            case "JOKE_RND":
            var info = result.Info;
            if (info.result != true)
                return;
            this.showJokeInfo(info.info);
            break;

            // 过审成功
            case "JOKE_CHECK":
            var info = result.Info;
            if (info.result == true) {
                this.showMsg("操作成功！");
                this.getRndJoke();
            } else {
                this.showMsg(info.info);
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
