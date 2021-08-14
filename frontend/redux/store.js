import { configureStore, getDefaultMiddleware } from '@reduxjs/toolkit';
import reducer from './reducer';
import constants from '../constants';

const { IS_PROD } = constants;

const middleware = getDefaultMiddleware({ serializableCheck: false });

if (!IS_PROD) {
  // push in logger as the last middleware
  // eslint-disable-next-line global-require
  const { createLogger } = require('redux-logger');
  middleware.push(createLogger({ diff: true, collapsed: true }));
}

const store = configureStore({
  devTools: !IS_PROD,
  middleware,
  reducer,
});

export default store;
