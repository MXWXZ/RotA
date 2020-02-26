import React from 'react';
import ReactDOM from 'react-dom';
import LoginForm from './login';
import Room from './room';
import { BrowserRouter as Router, Route, Redirect } from 'react-router-dom';
import Server from './server';
import { Spin, ConfigProvider } from 'antd';
import 'antd/dist/antd.css';
import zhCN from 'antd/es/locale/zh_CN';

class AuthRoute extends React.Component {
    state = {
        loading: true,
    }

    checked = () => {
        if (this.state.loading === true)
            this.setState({ loading: false });
    }

    render() {
        const { component: Component, ...rest } = this.props;
        const isLogged = sessionStorage.getItem("id") != null &&
            sessionStorage.getItem("token") != null;

        if (!isLogged) {
            return <Redirect to='/' />
        } else {
            Server.CheckToken(this.checked);
            if (this.state.loading === true)
                return <Spin />
            else
                return <Component {...rest} />;
        }
    }
}

const RotaRouter = (
    <ConfigProvider locale={zhCN}>
        <Router>
            <Route exact path='/' component={LoginForm} />
            <Route exact path='/room' render={(props) => <AuthRoute {...props} component={Room} />} />
        </Router>
    </ConfigProvider>
);

ReactDOM.render(RotaRouter, document.getElementById('root'));