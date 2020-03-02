import React from 'react'
import { Input } from 'antd'
import Server from '../../server'
import { EnterOutlined } from '@ant-design/icons'

const { Search } = Input;

function ChatInput() {
    const r = React.createRef();
    return (
        <Search ref={r} onSearch={v => {
            if (v !== '') {
                r.current.input.setValue('');
                Server.Send('Chat', { Channel: 1, Msg: v, });
            }
        }} enterButton={<EnterOutlined />} />
    )
}

export default ChatInput;