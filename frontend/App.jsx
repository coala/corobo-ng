import React, { useEffect } from 'react';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import { useDispatch } from 'react-redux';

import NavBar from './components/Navbar/NavBar';
import Login from './pages/AuthPages/Login';
import Logout from './pages/AuthPages/Logout';
import AuthComplete from './pages/AuthPages/AuthComplete';
import HomePage from './pages/HomePage/HomePage';

import { getLoggedInState } from './pages/AuthPages/slices/appSlice';

import './assets/stylesheets/base.scss';
import './assets/stylesheets/app.scss';

const App = () => {
  const dispatch = useDispatch();

  useEffect(() => {
    dispatch(getLoggedInState());
  }, []);

  return (
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
          <Route exact path="/logout" component={Logout} />
        </Switch>
      </div>
    </Router>
  );
};

export default App;
