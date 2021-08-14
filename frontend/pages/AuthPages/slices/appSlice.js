import { createSlice } from '@reduxjs/toolkit';

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
