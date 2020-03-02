import React, { Component } from 'react'
import { Button } from 'antd'
import Server from '../../server'

// type: 1 for ready 2 for master
// disabled: true for disabled
class InroomReady extends Component {
    state = {
        ready: false,
    }

    handleReady = () => {
        if (this.props.type === 1) {
            Server.Send('ReadyRoom', { Ready: this.state.ready === false ? 1 : 0 });
        } else {

        }
        this.setState((prev) => { return { ready: !prev.ready } });
    }

    render() {
        let text = '准备';
        if (this.props.type === 1 && this.state.ready === true)
            text = '取消准备';
        if (this.props.type === 2)
            text = '开始游戏';
        return (
            <Button type='primary' onClick={this.handleReady} block disabled={this.props.disabled}>{text}</Button>
        )
    }
}

export default InroomReady;