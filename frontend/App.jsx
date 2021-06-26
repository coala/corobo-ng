import React from 'react';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import NavBar from './components/NavBar';
import SignIn from './components/SignIn';

import './assets/stylesheets/base.scss';

const App = () => {
  return (
    <Router>
      <NavBar />
      <Switch>
        <Route exact path="/login">
          <SignIn />
        </Route>
      </Switch>
    </Router>
  );
};

export default App;
