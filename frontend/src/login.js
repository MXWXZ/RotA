import React, { Component } from 'react';
import { Button } from 'antd';
import { message, Row, Col, Form, Icon, Input, Typography } from 'antd';
import { Redirect } from 'react-router-dom';
import Server from './server';


const { Title } = Typography;

class Login extends Component {
    state = {
        loading: false,
        login: false,
    }

    handleLogin = (e) => {
        e.preventDefault();
        this.setState({ loading: true });
        this.props.form.validateFields((err, values) => {
            if (!err) {
                Server.AddHandler("LoginRsp", (data) => {
                    if (data.Msg.Status === 1) {
                        message.error("用户名或密码错误");
                        this.setState({ loading: false });
                    } else if (data.Msg.Status === 2) {
                        message.error("服务器错误");
                        this.setState({ loading: false });
                    } else {
                        sessionStorage.setItem('id', data.Msg.ID);
                        sessionStorage.setItem('token', data.Msg.Token);
                        message.success("登录成功", 1.5).then(() => this.setState({ login: true }));
                    }
                });
                Server.Send("Login", values);
            } else {
                this.setState({ loading: false });
            }
        });
    }

    handleSignup = (e) => {
        e.preventDefault();
        this.setState({ loading: true });
        this.props.form.validateFields((err, values) => {
            if (!err) {
                Server.AddHandler("SignupRsp", (data) => {
                    if (data.Msg.Status === 1) {
                        message.error("用户名已存在");
                        this.setState({ loading: false });
                    } else if (data.Msg.Status === 2) {
                        message.error("服务器错误");
                        this.setState({ loading: false });
                    } else {
                        message.success("注册成功");
                        setTimeout(() => { window.location.reload(); }, 1000);
                    }
                });
                Server.Send("Signup", values);
            } else {
                this.setState({ loading: false });
            }
        });
    }

    componentWillUnmount() {
        Server.DeleteHandler("LoginRsp");
        Server.DeleteHandler("SignupRsp");
    }

    render() {
        const { getFieldDecorator } = this.props.form;
        const isLogged = sessionStorage.getItem("id") != null &&
            sessionStorage.getItem("token") != null;
        if (this.state.login || isLogged)
            return <Redirect to='/room' />
        else
            return (
                <div>
                    <Row>
                        <Col span={8} offset={8}>
                            <Title level={2} style={{ paddingTop: '20px' }}>RotA - Reunion of the Ancients</Title>
                        </Col>
                    </Row>
                    <Row>
                        <Col span={7} offset={8}>
                            <Form onSubmit={this.handleLogin}>
                                <Form.Item>
                                    {
                                        getFieldDecorator('UserName', { rules: [{ required: true, message: '请输入用户名！' }] })(
                                            <Input prefix={<Icon type='user' />} placeholder='用户名' />)
                                    }
                                </Form.Item>
                                <Form.Item>
                                    {
                                        getFieldDecorator('UserPass', { rules: [{ required: true, message: '请输入密码！' }] })(
                                            <Input prefix={<Icon type='lock' />} type='password' placeholder='密码' />)
                                    }
                                </Form.Item>
                                <Form.Item>
                                    <Col span={6} offset={4}>
                                        <Button type='primary' htmlType='submit' loading={this.state.loading} block>登录</Button>
                                    </Col>
                                    <Col span={6} offset={2}>
                                        <Button loading={this.state.loading} onClick={this.handleSignup} block>注册</Button>
                                    </Col>
                                </Form.Item>
                            </Form>
                        </Col>
                    </Row>
                </div>
            );
    }
}

const LoginForm = Form.create({ name: 'login' })(Login);

export default LoginForm;