import React from 'react'
import { Col, Row } from 'antd'
import RoomView from '../component/room/room_view'
import RoomCreate from '../component/room/room_create'

function Room() {
    return (
        <Row>
            <Col span={16} offset={4} style={{ marginTop: '10px', marginBottom: '10px' }}>
                <RoomCreate />
            </Col>
            <Col span={16} offset={4}>
                <RoomView />
            </Col>
        </Row >
    )
}

export default Room;