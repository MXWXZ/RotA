import React from 'react';
import ReactDOM from 'react-dom';
import LoginForm from './login';
import { BrowserRouter as Router, Route } from 'react-router-dom';
import 'antd/dist/antd.css';

const RotaRouter = (
    <Router>
        <Route exact path='/' component={LoginForm} />
    </Router>
);

ReactDOM.render(RotaRouter, document.getElementById('root'));