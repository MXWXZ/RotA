import React, { Component } from 'react'
import { Button } from 'antd'
import Server from '../../server'
import RoomForm from './room_form'

class RoomCreate extends Component {
    formRef = React.createRef();

    state = {
        visible: false,
        loading: false,
    }

    handleSubmit = (v, a) => {
        this.setState({ loading: true });
        v.Type = parseInt(v.Type);
        Server.Send('NewRoom', v);
        this.setState({
            visible: false,
            loading: false,
        });
    }

    render() {
        return (
            <div>
                <Button type='primary' style={{ float: 'right' }} onClick={() => this.setState({ visible: true })}>新建房间</Button>
                <RoomForm visible={this.state.visible} title='新建房间' onCancel={() => this.setState({ visible: false })}
                    onFinish={this.handleSubmit} loading={this.state.loading} />
            </div>
        )
    }
}

export default RoomCreate;