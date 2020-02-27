import React, { Component } from 'react';
import { Tag, Table, Col, Row, Button, Modal } from 'antd';
import { Form, Icon, Input, Select } from 'antd';
import { Link } from 'react-router-dom';
import Server from './server';

const { Option } = Select;

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
        render: status => {
            if (status)
                return <font color="grey">游戏中</font>
            else
                return <font color="green">可加入</font>
        }
    },
    {
        title: '操作',
        key: 'action',
        width: 100,
        align: 'center',
        render: (_, record) => {
            if (record.Status)
                return <Link to={'/room/' + record.ID.toString()} disabled>加入</Link>
            else
                return <Link to={'/room/' + record.ID.toString()}>加入</Link>
        },
    },
];

class RoomForm extends Component {
    state = {
        data: [],
        visible: false,
        loading: false,
    }

    showModal = (e) => {
        this.setState({ visible: true });
    }

    handleCancel = () => {
        this.setState({ visible: false });
    }

    handleOk = (e) => {
        this.setState({ loading: true });
        this.props.form.validateFields((err, values) => {
            if (!err) {
                values.Type = parseInt(values.Type);
                Server.Send("NewRoom", values);
                this.setState({
                    visible: false,
                    loading: false,
                });
            } else {
                this.setState({ loading: false });
            }
        })
    }

    componentWillUnmount() {
        Server.DeleteHandler("NewRoomRsp");
        Server.DeleteHandler("DeleteRoomRsp");
        Server.DeleteHandler("RoomInfoRsp");
        Server.DeleteHandler("GetRoomsRsp");
    }

    componentDidMount() {
        Server.AddHandler("NewRoomRsp", (d) => {
            this.setState({ data: this.state.data.concat(d.Msg) })
            return true;
        });
        Server.AddHandler("DeleteRoomRsp", (d) => {
            let newdata = this.state.data;
            for (let i = 0, len = newdata.length; i < len; ++i) {
                if (newdata[i].ID === d.Msg.ID) {
                    newdata.splice(i, 1);
                    break;
                }
            }
            this.setState({ data: newdata });
            return true;
        });
        Server.AddHandler("RoomInfoRsp", (d) => {
            let newdata = this.state.data;
            for (let i = 0, len = newdata.length; i < len; ++i) {
                if (newdata[i].ID === d.Msg.ID) {
                    newdata[i] = d.Msg
                    break;
                }
            }
            this.setState({ data: newdata });
            return true;
        });
        Server.AddHandler("GetRoomsRsp", (d) => {
            this.setState({ data: d.Msg.RoomInfo })
        });
        Server.Send("GetRooms", {});
    }

    render() {
        const { getFieldDecorator } = this.props.form;
        return (
            <Row>
                <Col span={16} offset={4} style={{ marginTop: '10px', marginBottom: '10px' }}>
                    <Button type='primary' style={{ float: 'right' }} onClick={this.showModal}>新建房间</Button>
                    <Modal visible={this.state.visible} title='新建房间' onOk={this.handleOk} onCancel={this.handleCancel}
                        width={400} confirmLoading={this.state.loading}>
                        <Form labelCol={{ span: 6 }} wrapperCol={{ span: 18 }} labelAlign='left'>
                            <Form.Item label="房间名">
                                {
                                    getFieldDecorator('Name', { rules: [{ required: true, message: '请输入房间名！' }] })(
                                        <Input prefix={<Icon type="home" />} placeholder='房间名' />)
                                }
                            </Form.Item>
                            <Form.Item label="模式">
                                {
                                    getFieldDecorator('Type', { initialValue: "1", rules: [{ required: true, message: '请选择模式！' }] })(
                                        <Select>
                                            <Option value="1">1v1</Option>
                                        </Select>)
                                }
                            </Form.Item>
                        </Form>
                    </Modal>
                </Col>
                <Col span={16} offset={4}>
                    <Table rowKey={record => record.ID} bordered={true} columns={columns} dataSource={this.state.data} />
                </Col>
            </Row >
        );
    }
}

const Room = Form.create({ name: 'room' })(RoomForm);

export default Room;