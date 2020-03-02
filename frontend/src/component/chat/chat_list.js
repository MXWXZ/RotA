import React from 'react'
import { List } from 'antd'

// data: list data
function ChatList(props) {
    return (
        <List size='small' dataSource={props.data}
            renderItem={item => <List.Item style={{ paddingLeft: '10px', paddingRight: '10px', paddingTop: 0, paddingBottom: 0 }}>{item}</List.Item>}
            split={false} />
    )
}

export default ChatList;