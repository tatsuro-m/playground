const initialState = {};

// eslint-disable-next-line @typescript-eslint/explicit-module-boundary-types
const reducer = (state: any, action: any): any => {
  switch (action.type) {
    case "login":
      return action.payload.user;
    case "logout":
      return initialState;
    default:
      return state;
  }
};

export default {
  initialState,
  reducer,
};
