import React from 'react'
import Login from './page/login'
import Room from './page/room'
import Inroom from './page/inroom'
import { BrowserRouter as Router, Route, Redirect } from 'react-router-dom'
import Server from './server'
import { ConfigProvider } from 'antd'
import zhCN from 'antd/es/locale/zh_CN'
import UtilLoading from './component/util/util_loading'
import { GetID, GetStorage } from './storage'

class AuthRoute extends React.Component {
    state = {
        loading: true,
        islogin: GetID() != null && GetStorage('Token') != null,
    }

    checked = () => {
        if (this.state.loading === true)
            this.setState({ loading: false });
    }

    componentDidMount() {
        const tmp = GetID() != null && GetStorage('Token') != null;
        this.setState({ islogin: tmp });
        if (tmp && this.state.loading)
            Server.CheckToken(this.checked);
    }

    render() {
        const { component: Component, ...rest } = this.props;

        if (!this.state.islogin) {
            return <Redirect to='/' />
        } else {
            if (this.state.loading === true)
                return <UtilLoading />
            else
                return <Component {...rest} />;
        }
    }
}

const RotaRouter = (
    <ConfigProvider locale={zhCN}>
        <Router>
            <Route exact path='/' component={Login} />
            <Route exact path='/room' render={(props) => <AuthRoute {...props} component={Room} />} />
            <Route exact path='/room/:rid' render={(props) => <AuthRoute {...props} component={Inroom} />} />
        </Router>
    </ConfigProvider>
);

export default RotaRouter;