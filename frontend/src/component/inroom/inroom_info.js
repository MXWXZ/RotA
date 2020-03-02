import React from 'react'
import { Tag, PageHeader, Descriptions } from 'antd'
import { useHistory } from 'react-router-dom'
import Server from '../../server'

// rid: room id
// name: room name
// type: room type
// status: room status
// capacity: room capacity
// size: room size
function InroomInfo(props) {
    const history = useHistory();
    const handleBack = () => {
        Server.Send('ExitRoom', {});
        history.push("/room");
    }

    let mod, status;
    switch (props.type) {
        case 1:
            mod = <Tag color='red'>1v1</Tag>
            break;
        default:
            mod = <Tag color='grey'>Unknown</Tag>
            break;
    }

    if (props.status === 1)
        status = <font color='grey'>游戏中</font>
    else if (props.capacity === props.size)
        status = <font color='grey'>已满</font>
    else
        status = <font color='green'>可加入</font>

    return (
        <PageHeader title={'房间号：' + props.rid} subTitle={props.name} onBack={handleBack}
            style={{ paddingLeft: 0 }}>
            <Descriptions size='small' column={3}>
                <Descriptions.Item key='mod' label='模式'>{mod}</Descriptions.Item>
                <Descriptions.Item key='size' label='人数'>{props.size} / {props.capacity}</Descriptions.Item>
                <Descriptions.Item key='status' label='状态'>{status}</Descriptions.Item>
            </Descriptions>
        </PageHeader>
    )
}

export default InroomInfo;