import ReconnectingWebSocket from 'reconnecting-websocket';

var ws = null
var evt = {}

class Server {
    static Connect() {
        return new Promise((resolve, reject) => {
            if (!ws) {
                ws = new ReconnectingWebSocket(process.env.REACT_APP_BACKEND);
                evt = {}
                ws.addEventListener('message', this.handler);
                ws.addEventListener('open', () => {
                    resolve();
                });
            } else {
                resolve();
            }
        })
    }

    static AddHandler(n, h) {
        evt[n] = h;
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

export default Server;