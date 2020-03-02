import React, { useState, useEffect } from 'react'
import { Modal, Form, Input, Select } from 'antd'

const { Option } = Select;

// visible: false to hide
// loading: false to enable click
// title: form title
// onCancel: cancel callback
// onFinish: finish callback
function RoomForm(props) {
    const [form] = Form.useForm();
    const [enreset, setReset] = useState(false);

    useEffect(() => {
        if (props.loading)
            setReset(true);
        if (!props.loading && enreset) {
            form.resetFields();
            setReset(false);
        }
    }, [props.loading, enreset, form]);

    return (
        <Modal visible={props.visible} title={props.title} onCancel={props.onCancel} onOk={() => { form.submit(); }}
            width={400} confirmLoading={props.loading} >
            <Form initialValues={{ Type: '1' }} onFinish={props.onFinish} form={form} labelCol={{ span: 6 }} wrapperCol={{ span: 18 }} labelAlign='left'>
                <Form.Item label='房间名' name='Name' rules={[
                    {
                        required: true,
                        message: '请输入房间名！',
                    },
                ]}>
                    <Input placeholder='房间名' />
                </Form.Item>
                <Form.Item label='模式' name='Type' rules={[
                    {
                        required: true,
                        message: '请选择模式！',
                    }
                ]}>
                    <Select>
                        <Option value='1'>1v1</Option>
                    </Select>
                </Form.Item>
            </Form>
        </Modal>
    )
}

export default RoomForm