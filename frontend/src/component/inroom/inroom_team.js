import React from 'react'
import { Button, List } from 'antd'
import { StarTwoTone, CheckCircleTwoTone } from '@ant-design/icons'
import Server from '../../server'

function expandData(d, c) {
    while (d.length < c)
        d.push([]);
    return d;
}

function renderItem(item) {
    let inside;
    if (item.length === 0)
        inside = <Button type="link" block disabled>待加入</Button>
    else if (item[1] === true)
        inside = <div><StarTwoTone twoToneColor='#eb2f96' />{item[0]}</div>
    else
        inside = item[0]

    return (
        <List.Item extra={item[2] === 1 ? <CheckCircleTwoTone twoToneColor="#52c41a" /> : ''}>
            {inside}
        </List.Item>
    )
}

function changeTeam(id) {
    Server.Send('ChangeTeam', {
        Team: id,
    })
}

// name: team name
// data: data source
// capacity: team capacity
// team: teamid
// changable: true to show change
function InroomTeam(props) {
    return (
        <List header={<div>{props.name} {props.changable ?
            <Button style={{ float: "right" }} size="small" onClick={() => { changeTeam(props.team) }}>加入阵营</Button> : ''}</div>}
            itemLayout="vertical" dataSource={expandData(props.data, props.capacity)} renderItem={item => renderItem(item)}
        />
    )
}

export default InroomTeam;