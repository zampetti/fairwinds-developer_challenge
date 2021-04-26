import { configureStore, getDefaultMiddleware } from '@reduxjs/toolkit';
import boardReducer from '../components/boardSlice';
// import createSagaMiddleware from "redux-saga";
// import saga from './sagas'

// let sagaMiddleware = createSagaMiddleware();
// const middleware = [...getDefaultMiddleware({ thunk: false }), sagaMiddleware];

export default configureStore({
  reducer: {
    board: boardReducer,
  },
//   middleware
});

// sagaMiddleware.run(saga)