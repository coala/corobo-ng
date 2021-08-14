import React from 'react';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import { Provider } from 'react-redux';

import NavBar from './components/Navbar/NavBar';
import Login from './pages/AuthPages/Login';
import AuthComplete from './pages/AuthPages/AuthComplete';
import HomePage from './pages/HomePage/HomePage';

import store from './redux/store';

import './assets/stylesheets/base.scss';
import './assets/stylesheets/app.scss';

const App = () => {
  return (
    <Provider store={store}>
      <Router>
        <div className="app">
          <NavBar />
        </div>
        <div className="app__body">
          <Switch>
            <Route exact path="/" component={HomePage} />
            <Route exact path="/login" component={Login} />
            <Route
              exact
              path="/login/github/complete"
              render={() => <AuthComplete provider="github" />}
            />
            <Route
              exact
              path="/login/gitlab/complete"
              render={() => <AuthComplete provider="gitlab" />}
            />
          </Switch>
        </div>
      </Router>
    </Provider>
  );
};

export default App;
