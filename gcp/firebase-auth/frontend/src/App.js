import './App.css';
import SignInScreen from "./pages/SignInScreen";
import firebase from "firebase";
import {useEffect, useState} from "react";

function App() {
  const [isSignedIn, setIsSignedIn] = useState(false); // Local signed-in state.

  // Listen to the Firebase Auth state and set the local state.
  useEffect(() => {
    const unregisterAuthObserver = firebase.auth().onAuthStateChanged(user => {
      setIsSignedIn(!!user);
    });
    return () => unregisterAuthObserver(); // Make sure we un-register Firebase observers when the component unmounts.
  }, []);

  if (isSignedIn) {
    // ログインされているなら反応するはず
    firebase.auth().currentUser.getIdToken(/* forceRefresh */ true).then(function (idToken) {
      // Send token to your backend via HTTPS
      // ...
      console.log("id token の取得に成功しました！");
      console.log(idToken);
    }).catch(function (error) {
      console.log("id token の取得に失敗しました！");
      console.log(error);
    });
  }

  return (
    <div className="App">
      <h1>Firebase Auth デモアプリ</h1>
      <p>{isSignedIn ? "サインインしているっぽい" : "ログアウト中"}</p>
      <SignInScreen/>
    </div>
  );
}

export default App;
