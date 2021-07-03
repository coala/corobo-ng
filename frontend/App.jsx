import React from 'react';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import NavBar from './components/NavBar';
import SignIn from './components/SignIn';
import AuthComplete from './components/AuthComplete';
import HomePage from './components/HomePage';

import './assets/stylesheets/base.scss';
import './assets/stylesheets/app.scss';

const App = () => {
  return (
    <Router>
      <div className="app">
        <NavBar />
      </div>
      <div className="app__body">
        <Switch>
          <Route exact path="/" component={HomePage} />
          <Route exact path="/login" component={SignIn} />
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
  );
};

export default App;
