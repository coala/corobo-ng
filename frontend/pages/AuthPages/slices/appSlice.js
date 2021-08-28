import { createSlice } from '@reduxjs/toolkit';
import api from '../../../api';

const { actions, reducer } = createSlice({
  name: 'auth',
  initialState: {
    isLoggedIn: false,
  },
  reducers: {
    updateAuth: (state, { payload }) => {
      const { isLoggedIn } = payload;
      state.isLoggedIn = isLoggedIn;
    },
  },
});

export default reducer;

export const { updateAuth } = actions;

export const getLoggedInState = () => (dispatch) => {
  api({
    method: 'GET',
    url: '/login',
  })
    .then((response) => {
      dispatch(updateAuth({ isLoggedIn: response.data.success }));
    })
    .catch(() => dispatch(updateAuth({ isLoggedIn: false })));
};
