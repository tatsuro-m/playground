import firebase from 'firebase/app';
import 'firebase/auth';
import client from '../apollo-client'; // If you need it

export const config = {
  apiKey: process.env.NEXT_PUBLIC_FIREBASE_API_KEY,
  authDomain: process.env.NEXT_PUBLIC_FIREBASE_AUTH_DOMAIN,
  projectId: process.env.NEXT_PUBLIC_FIREBASE_PROJECT_ID,
  storageBucket: process.env.NEXT_PUBLIC_FIREBASE_STORAGE_BUCKET,
  messagingSenderId: process.env.NEXT_PUBLIC_FIREBASE_MESSAGING_SENDER_ID,
  appId: process.env.NEXT_PUBLIC_FIREBASE_APP_ID,
};

!firebase.apps.length ? firebase.initializeApp(config) : firebase.app();

export const auth = firebase.auth();

// ログイン状態の検知
export const listenAuthState = (dispatch: any) => {
  return firebase.auth().onAuthStateChanged(function (user) {
    if (user) {
      // User is signed in.
      dispatch({
        type: 'login',
        payload: {
          user,
        },
      });
    } else {
      // User is signed out.
      // ...
      dispatch({
        type: 'logout',
      });
    }
  });
};

export const firebaseUser = (): firebase.User | null => {
  return firebase.auth().currentUser;
};

export const onLogout = (): void => {
  auth.signOut().then(() => {
    client.resetStore();
  });
};

// Configure FirebaseUI.
export const uiConfig = {
  // Popup signin flow rather than redirect flow.
  signInFlow: 'popup',
  signInOptions: [firebase.auth.GoogleAuthProvider.PROVIDER_ID],
  callbacks: {
    // Avoid redirects after sign-in.
    signInSuccessWithAuthResult: (): boolean => {
      client.resetStore();
      return false;
    },
  },
};
