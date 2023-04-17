import React from "react"
import { useAuth } from "./reducers/Auth"
import Login from "./components/Login/Login"
import Landing from './components/Landing/Landing'
import { styled, createTheme, ThemeProvider } from "@mui/material/styles";
import { getDesignTokens } from "./assets/assets";

export const Container = () => {
    const authReducer = useAuth()

    React.useEffect(()=>{
        // check token 
        
    },[])

    return (
        <ThemeProvider theme={createTheme(getDesignTokens(authReducer.auth.theme))}>
       { authReducer.auth.accessToken !== "" ? <Landing authReducer={authReducer} /> : <Login authReducer={authReducer} />}
        </ThemeProvider>
    )
}