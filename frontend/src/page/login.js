import React from 'react'
import { Row, Col, Typography } from 'antd'
import { Redirect, useHistory } from 'react-router-dom'
import LoginForm from '../component/login/login_form'
import { GetID, GetStorage } from '../storage'

const { Title } = Typography;

function Login() {
    const history = useHistory();
    const isLogged = GetID() != null &&
        GetStorage('Token') != null;
    if (isLogged)
        return <Redirect to='/room' />
    else
        return (
            <div>
                <Row style={{ marginBottom: '10px' }}>
                    <Col span={8} offset={8}>
                        <Title level={2} style={{ paddingTop: '20px' }}>RotA - Reunion of the Ancients</Title>
                    </Col>
                </Row>
                <Row>
                    <Col span={7} offset={8}>
                        <LoginForm onLogin={() => history.push("/room")} onSignup={() => window.location.reload()}
                            span={20} offset={2} />
                    </Col>
                </Row>
            </div>
        )
}

export default Login;