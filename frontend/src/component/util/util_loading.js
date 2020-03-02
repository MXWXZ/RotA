import React from 'react'
import { Result, Spin, Button } from 'antd'

function UtilLoading() {
    return (
        <Result icon={<Spin />} title="正在加载"
            extra={<Button onClick={() => window.location.href = '/'}>刷新</Button>}
        />
    )
}

export default UtilLoading;