import * as React from "react";
import {Container} from "./Container"
import { AuthProvider } from './reducers/Auth';

function App() {
  return (
    <AuthProvider>
        <Container />
    </AuthProvider>
  );
}

export default App;
