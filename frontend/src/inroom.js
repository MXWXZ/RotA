import React, { Component } from 'react';
import { Card, PageHeader, Row, Col, List, Descriptions, Tag, Typography, Input, Icon, Button } from 'antd';
import { Redirect } from 'react-router-dom';
import Server from './server';

const { StarTwoTone } = Icon;
const { Search } = Input;

class Inroom extends Component {
    state = {
        back: false,
        data: { "ID": 1, "Name": "3", "Type": 1, "Size": 1, "Capacity": 2, "Master": 1, "Status": 0, "Members": [{ "ID": 1, "Name": "Admin", "Team": 1 }] },
        team: [[], [], [], [], [], [], [], [], [], [], [], [], [], [], [], []],
    }

    componentDidMount() {
        // Server.AddHandler("RoomInfoRsp", (d) => {
        //     this.setState({ data: d.Msg })
        //     return true;
        // });
        let newteam = [[], [], [], [], [], [], [], [], [], [], [], [], [], [], [], []];
        for (let i = 0, len = this.state.data.Members.length; i < len; ++i)
            newteam[this.state.data.Members[i].Team].push(this.state.data.Members[i].Name);
        this.setState({ team: newteam });
    }

    componentWillUnmount() {
        Server.DeleteHandler("RoomInfoRsp");
    }

    render() {
        let mod, status, room;
        switch (this.state.data.Type) {
            case 1:
                mod = <Tag color="red">1v1</Tag>
                room = (
                    <Col span={12}>
                        <List header={<div>近卫</div>}
                            dataSource={this.state.team[1]}
                            renderItem={item => (
                                <List.Item>
                                    <StarTwoTone twoToneColor="#eb2f96" /> {item}
                                </List.Item>
                            )}
                        />
                        <List header={<div>天灾</div>}
                            dataSource={this.state.team[2]}
                            renderItem={item => (
                                <List.Item>
                                    <Typography.Text mark>[ITEM]</Typography.Text> {item}
                                </List.Item>
                            )}
                        />
                    </Col>
                )
                break;
            default:
                mod = <Tag color="grey">Unknown</Tag>
                break;
        }

        if (this.state.data.Status === 1)
            status = <font color="grey">游戏中</font>
        else if (this.state.data.Capacity === this.state.data.Size)
            status = <font color="grey">已满</font>
        else
            status = <font color="green">可加入</font>

        return (
            <Row>
                <Col span={16} offset={4}>
                    <Row>
                        <PageHeader title={'房间号：' + this.state.data.ID} subTitle={this.state.data.Name} onBack={() => this.setState({ back: true })}
                            style={{ paddingLeft: 0 }}>
                            <Descriptions size="small" column={3}>
                                <Descriptions.Item label="模式">{mod}</Descriptions.Item>
                                <Descriptions.Item label="人数">{this.state.data.Size} / {this.state.data.Capacity}</Descriptions.Item>
                                <Descriptions.Item label="状态">{status}</Descriptions.Item>
                            </Descriptions>
                        </PageHeader>
                    </Row>
                    <Row>
                        {room}
                        <Col span={10} offset={1}>
                            <Button key="1" type="primary" style={{ marginBottom: "15px" }} block>准备</Button>
                            <Card style={{ marginBottom: "15px", height: 300, overflowY: "scroll" }}>
                                <p>A: Card content</p>
                                <p>A: Card content</p><p>A: Card content</p><p>A: Card content</p><p>A: Card content</p><p>A: Card content</p><p>A: Card content</p>
                                <p>A: Card content</p><p>B: Card content</p>
                            </Card>
                            <Search placeholder="input search text" onSearch={value => console.log(value)} enterButton={<Icon type="enter" />} />
                        </Col>
                    </Row>

                </Col>
            </Row>
        );
    }
}

export default Inroom;