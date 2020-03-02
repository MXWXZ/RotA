import ReconnectingWebSocket from 'reconnecting-websocket'
import { message } from 'antd'
import { GetID, SetStorage, GetStorage, Clear } from './storage'

var ws = new ReconnectingWebSocket(process.env.REACT_APP_BACKEND);
var evt = {};
var checked = false;

class Server {
    static CheckToken(cb) {
        if (checked === false) {
            Server.AddHandler("CheckTokenRsp", (data) => {
                if (data.Msg.Code === 1) {
                    message.error("凭据已过期，请重新登录");
                    Clear();
                    setTimeout(() => {
                        window.location.href = "/"
                    }, 1000);
                } else {
                    SetStorage('Name', data.Msg.Name);
                    SetStorage('Room', data.Msg.Room);
                    SetStorage('Status', data.Msg.Status);
                    checked = true;
                    cb();
                }
            });
            Server.Send("CheckToken", {
                ID: GetID(),
                Token: GetStorage('Token'),
            });
        } else {
            cb();
        }
    }

    static AddHandler(n, h) {
        evt[n] = h;
    }

    static DeleteHandler(n) {
        delete evt[n];
    }

    static handler(data) {
        var reader = new FileReader();
        reader.onload = (e) => {
            var d = JSON.parse(e.target.result);
            if (d.Response.Type in evt)
                if (!evt[d.Response.Type](d.Response))
                    delete evt[d.Response.Type];
        }
        reader.readAsText(data.data);
    }

    static Send(n, d) {
        var msg = {};
        msg[n] = d;

        ws.send(JSON.stringify(msg));
    }
}

ws.addEventListener('message', Server.handler);
Server.AddHandler("NeedTokenRsp", () => {
    checked = false;
    Server.CheckToken(() => {
        message.success("重新连接成功");
    });
    return true;
});
Server.AddHandler("Error", (d) => {
    message.error(d.Msg.Info);
    return true;
});

export default Server;