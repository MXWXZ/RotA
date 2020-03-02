import React, { Component } from 'react'
import { Row, Col, message } from 'antd'
import InroomTeam from '../component/inroom/inroom_team'
import InroomReady from '../component/inroom/inroom_ready'
import InroomInfo from '../component/inroom/inroom_info'
import UtilLoading from '../component/util/util_loading'
import Server from '../server'
import { Redirect } from 'react-router-dom'
import { GetID } from '../storage'
import ChatBox from '../component/chat/chat_box'
import ChatInput from '../component/chat/chat_input'

class Inroom extends Component {
    state = {
        back: false,
        loaded: false,
        disabled: true,
        data: {},
        cteam: 0,
        team: [[], [], [], [], [], [], [], [], [], [], [], [], [], [], [], []],
    }

    update = (d) => {
        this.setState({ data: d });
        let newteam = [[], [], [], [], [], [], [], [], [], [], [], [], [], [], [], []];
        for (let i = 0, len = this.state.data.Members.length; i < len; ++i) {
            if (this.state.data.Members[i].ID === GetID())
                this.setState({ cteam: this.state.data.Members[i].Team });
            newteam[this.state.data.Members[i].Team].push([this.state.data.Members[i].Name,
            this.state.data.Members[i].ID === this.state.data.Master,
            this.state.data.Members[i].Ready]);
        }
        let block = true;
        if (this.state.data.Master === GetID()) {
            if (this.state.data.Capacity === this.state.data.Size) {
                let cnt = 0;
                newteam.forEach((t) => { t.forEach((p) => cnt += p[2]) });
                if (cnt === this.state.data.Capacity - 1)
                    block = false;
            }
        } else {
            block = false;
        }
        this.setState({ team: newteam, loaded: true, disabled: block });
    }

    componentDidMount() {
        Server.AddHandler('JoinRoomRsp', (d) => {
            if (d.Msg.Code === 1) {
                message.error('房间已满');
            } else if (d.Msg.Code === 2) {
                message.error('加入失败，你已经在一个房间中了');
                window.location.href = '/';
            } else if (d.Msg.Code === 3) {
                message.error('房间不存在');
            } else {
                this.update(d.Msg.Info);
            }
            if (d.Msg.Code !== 0)
                this.setState({ loaded: true, back: true });
        });
        Server.AddHandler('RoomInfoRsp', (d) => {
            this.update(d.Msg);
            return true;
        });
        Server.Send('JoinRoom', {
            ID: parseInt(this.props.match.params.rid)
        })
    }

    componentWillUnmount() {
        Server.DeleteHandler('RoomInfoRsp');
        Server.DeleteHandler('JoinRoomRsp');
    }

    render() {
        if (!this.state.loaded)
            return <UtilLoading />

        if (this.state.back)
            return <Redirect to='/room' />

        let room;
        switch (this.state.data.Type) {
            case 1:
                room = (
                    <Col span={12}>
                        <InroomTeam name='近卫' data={this.state.team[1]} capacity={1} team={1} changable={this.state.cteam !== 1} />
                        <InroomTeam name='天灾' data={this.state.team[2]} capacity={1} team={2} changable={this.state.cteam !== 2} />
                    </Col>
                )
                break;
            default:
                break;
        }

        return (
            <Row>
                <Col span={16} offset={4}>
                    <Row>
                        <InroomInfo rid={this.state.data.ID} name={this.state.data.Name} type={this.state.data.Type}
                            status={this.state.data.Status} size={this.state.data.Size} capacity={this.state.data.Capacity} />
                    </Row>
                    <Row>
                        {room}
                        <Col span={10} offset={1}>
                            <InroomReady type={this.state.data.Master === GetID() ? 2 : 1} disabled={this.state.disabled} />
                            <div style={{ marginTop: '15px', marginBottom: '15px', height: 300 }} >
                                <ChatBox />
                            </div>
                            <ChatInput />
                        </Col>
                    </Row>

                </Col>
            </Row>
        );
    }
}

export default Inroom;