import React, { Component } from 'react';
import { Button } from 'antd';
import { message, Row, Col, Form, Icon, Input, Typography } from 'antd';
import Server from './server';

const { Title } = Typography;

class Login extends Component {
    state = {
        loading: false,
    }

    handleLogin = (e) => {
        e.preventDefault();
        this.setState({ loading: true });
        this.props.form.validateFields((err, values) => {
            if (!err) {
                Server.Connect().then(() => {
                    Server.AddHandler("Login", (data) => {
                        if (data.Msg.Status === 1) {
                            message.error("Invalid username or password");
                            this.setState({ loading: false });
                        } else {
                            sessionStorage.setItem('token', data.Msg.Token);
                            message.success("Login success!");
                            setTimeout(() => { window.location.reload(); }, 1000);
                        }
                    })
                    Server.Send("Login", values);
                })
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
                Server.Connect().then(() => {
                    Server.AddHandler("Signup", (data) => {
                        if (data.Msg.Status === 1) {
                            message.error("Username is already existed");
                            this.setState({ loading: false });
                        } else {
                            message.success("Signup success!");
                            setTimeout(() => { window.location.reload(); }, 1000);
                        }
                    })
                    Server.Send("Signup", values);
                })
            } else {
                this.setState({ loading: false });
            }
        });
    }

    render() {
        const { getFieldDecorator } = this.props.form;
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
                                    getFieldDecorator('UserName', { rules: [{ required: true, message: 'Please input your username!' }] })(
                                        <Input prefix={<Icon type='user' />} placeholder='Username' />)
                                }
                            </Form.Item>
                            <Form.Item>
                                {
                                    getFieldDecorator('UserPass', { rules: [{ required: true, message: 'Please input your Password!' }] })(
                                        <Input prefix={<Icon type='lock' />} type='password' placeholder='Password' />)
                                }
                            </Form.Item>
                            <Form.Item>
                                <Col span={6} offset={4}>
                                    <Button type='primary' htmlType='submit' loading={this.state.loading} block>Sign in</Button>
                                </Col>
                                <Col span={6} offset={2}>
                                    <Button loading={this.state.loading} onClick={this.handleSignup} block>Sign up</Button>
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