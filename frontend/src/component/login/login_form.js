import React, { Component } from 'react'
import { LockOutlined, UserOutlined } from '@ant-design/icons'
import { message, Row, Col, Input, Form, Button } from 'antd'
import Server from '../../server'
import { SetID, SetStorage } from '../../storage'

// onLogin: Login success callback
// onSignup: Signup success callback
// span: Wrap span
// offset: Wrap offset
class LoginForm extends Component {
    state = {
        loading: false,
    }

    handleLogin = (v) => {
        this.setState({ loading: true });
        Server.AddHandler('LoginRsp', (data) => {
            if (data.Msg.Status === 1) {
                message.error('用户名或密码错误');
                this.setState({ loading: false });
            } else if (data.Msg.Status === 2) {
                message.error('服务器错误');
                this.setState({ loading: false });
            } else {
                SetID(data.Msg.ID);
                SetStorage('Token', data.Msg.Token);
                message.success('登录成功', 1).then(() => this.props.onLogin());
            }
        });
        Server.Send('Login', v);
    }

    handleSignup = (v) => {
        this.setState({ loading: true });
        Server.AddHandler('SignupRsp', (data) => {
            if (data.Msg.Status === 1) {
                message.error('用户名已存在');
                this.setState({ loading: false });
            } else if (data.Msg.Status === 2) {
                message.error('服务器错误');
                this.setState({ loading: false });
            } else {
                message.success('注册成功', 1).then(() => this.props.onSignup());
            }
        });
        Server.Send('Signup', v);
    }

    componentWillUnmount() {
        Server.DeleteHandler('LoginRsp');
        Server.DeleteHandler('SignupRsp');
    }

    render() {
        return (
            <Form wrapperCol={{ span: this.props.span, offset: this.props.offset }} onFinish={this.handleLogin}>
                <Form.Item name='UserName' rules={[
                    {
                        required: true,
                        message: '请输入用户名！',
                    },
                ]}>
                    <Input prefix={<UserOutlined />} placeholder='用户名' />
                </Form.Item>
                <Form.Item name='UserPass' rules={[
                    {
                        required: true,
                        message: '请输入密码！',
                    },
                ]}>
                    <Input.Password prefix={<LockOutlined />} placeholder='密码' />
                </Form.Item>
                <Form.Item>
                    <Row>
                        <Col span={6} offset={4}>
                            <Button type='primary' htmlType='submit' loading={this.state.loading} block>登录</Button>
                        </Col>
                        <Col span={6} offset={2}>
                            <Button loading={this.state.loading} onClick={this.handleSignup} block>注册</Button>
                        </Col>
                    </Row>
                </Form.Item>
            </Form>
        )
    }
}

export default LoginForm;