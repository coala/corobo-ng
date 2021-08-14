import { combineReducers } from '@reduxjs/toolkit';
import authReducer from '../pages/AuthPages/slices/appSlice';

const app = combineReducers({
  auth: authReducer,
});

const reducer = combineReducers({ app });

export default reducer;
