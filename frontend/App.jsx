import React from 'react';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import NavBar from './components/NavBar';
import SignIn from './components/SignIn';
import AuthComplete from './components/AuthComplete';

import './assets/stylesheets/base.scss';

const App = () => {
  return (
    <Router>
      <NavBar />
      <Switch>
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
    </Router>
  );
};

export default App;
