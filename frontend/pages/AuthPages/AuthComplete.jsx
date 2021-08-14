import React, { useEffect, useState } from 'react';
import { useLocation, useHistory } from 'react-router-dom';
import { useDispatch } from 'react-redux';
import PropTypes from 'prop-types';
import api from '../../api';
import { updateAuth } from './slices/appSlice';

import './stylesheets/authcomplete.scss';

const AuthComplete = ({ provider }) => {
  const location = useLocation();
  const history = useHistory();
  const dispatch = useDispatch();

  const [loading, setLoading] = useState(true);

  const queryParams = new URLSearchParams(location.search);

  useEffect(() => {
    const code = queryParams.get('code');
    if (code) {
      api({
        method: 'POST',
        url: `/login/${provider}`,
        data: {
          code: queryParams.get('code'),
        },
      })
        .then((response) => {
          if (response.data.success) {
            dispatch(updateAuth({ isLoggedIn: true }));
            history.replace('/');
          } else {
            setLoading(false);
          }
        })
        .catch(() => {
          setLoading(false);
        });
    } else {
      setLoading(false);
    }
  }, []);

  return (
    <div className="authComplete">
      <h1 className="authComplete__heading">
        {loading
          ? `Authenticating with ${provider}...`
          : 'Unable to authenticate user, please try again.'}
      </h1>
    </div>
  );
};

AuthComplete.propTypes = {
  provider: PropTypes.oneOf(['github', 'gitlab']).isRequired,
};

export default AuthComplete;
