import React, { Component } from 'react'
import ChatList from './chat_list'
import Server from '../../server';


class ChatBox extends Component {
    list = React.createRef();

    state = {
        data: [],
    }

    componentDidMount() {
        Server.AddHandler('ChatRsp', (d) => {
            this.setState((prev) => { return { data: prev.data.concat(<div><font color='green'>{d.Msg.Name}: </font>{d.Msg.Msg}</div>) } });
            this.list.current.scrollTop = this.list.current.scrollHeight;
            return true;
        });
    }

    componentWillUnmount() {
        Server.DeleteHandler('ChatRsp');
    }

    render() {
        return (
            <div ref={this.list} style={{ height: '100%', overflowY: 'scroll', border: '1px solid #d9d9d9', borderRadius: '2px' }} >
                <ChatList data={this.state.data} />
            </div>
        )
    }
}

export default ChatBox;