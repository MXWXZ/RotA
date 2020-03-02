import React, { Component } from 'react'
import { Tag, Table } from 'antd'
import { Link, Redirect } from 'react-router-dom'
import Server from '../../server'
import { GetID } from '../../storage'

const columns = [
    {
        title: '房间号',
        dataIndex: 'ID',
        align: 'center',
        width: 150,
        sorter: (a, b) => a.ID - b.ID,
    },
    {
        title: '房间名',
        dataIndex: 'Name',
        align: 'center',
    },
    {
        title: '模式',
        dataIndex: 'Type',
        width: 100,
        align: 'center',
        sorter: (a, b) => a.Type - b.Type,
        render: type => {
            let color;
            let text;
            switch (type) {
                case 1:
                    color = 'red';
                    text = '1v1';
                    break;
                default:
                    color = 'grey';
                    text = 'Unknown';
                    break;
            }
            return (
                <Tag color={color}>
                    {text}
                </Tag>
            );
        },
    },
    {
        title: '人数',
        key: 'Player',
        width: 100,
        align: 'center',
        sorter: (a, b) => (a.Capacity - a.Size) - (b.Capacity - b.Size),
        render: (_, record) => (
            <span>{record.Size} / {record.Capacity}</span>
        )
    },
    {
        title: '状态',
        dataIndex: 'Status',
        width: 100,
        align: 'center',
        sorter: (a, b) => a.Status - b.Status,
        render: (status, record) => {
            if (status)
                return <font color='grey'>游戏中</font>
            else if (record.Size === record.Capacity)
                return <font color='grey'>已满</font>
            else
                return <font color='green'>可加入</font>
        }
    },
    {
        title: '操作',
        key: 'action',
        width: 100,
        align: 'center',
        render: (_, record) => {
            if (record.Status || record.Size === record.Capacity)
                return <Link to={'/room/' + record.ID.toString()} disabled>加入</Link>
            else
                return <Link to={'/room/' + record.ID.toString()}>加入</Link>
        },
    },
];

class RoomView extends Component {
    state = {
        data: [],
        created: 0,
    }

    componentWillUnmount() {
        Server.DeleteHandler('NewRoomRsp');
        Server.DeleteHandler('DeleteRoomRsp');
        Server.DeleteHandler('RoomInfoRsp');
        Server.DeleteHandler('GetRoomsRsp');
    }

    componentDidMount() {
        Server.AddHandler('NewRoomRsp', (d) => {
            if (d.Msg.Master === GetID())
                this.setState((prev) => { return { data: prev.data.concat(d.Msg), created: d.Msg.ID } })
            else
                this.setState((prev) => { return { data: prev.data.concat(d.Msg) } })
            return true;
        });
        Server.AddHandler('DeleteRoomRsp', (d) => {
            let newdata = [...this.state.data];
            for (let i = 0, len = newdata.length; i < len; ++i) {
                if (newdata[i].ID === d.Msg.ID) {
                    newdata.splice(i, 1);
                    break;
                }
            }
            this.setState({ data: newdata });
            return true;
        });
        Server.AddHandler('RoomInfoRsp', (d) => {
            let newdata = [...this.state.data];
            for (let i = 0, len = newdata.length; i < len; ++i)
                if (newdata[i].ID === d.Msg.ID) {
                    newdata[i] = d.Msg
                    break;
                }

            this.setState({ data: newdata });
            return true;
        });
        Server.AddHandler('GetRoomsRsp', (d) => {
            this.setState({ data: d.Msg.RoomInfo })
        });
        Server.Send('GetRooms', {});
    }

    render() {
        if (this.state.created)
            return <Redirect to={'/room/' + this.state.created} />
        return (
            <Table rowKey={record => record.ID} bordered={true} columns={columns} dataSource={this.state.data} />
        )
    }
}

export default RoomView;