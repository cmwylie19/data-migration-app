import React, { useReducer, createContext, useContext, useState } from 'react'

let initialState = {
    accessToken: "",
    roles: [],
    given_name: "",
    family_name: "",
    email: "",
    theme: "light"
}

const AuthReducer = (state, action) => {
    switch (action.type) {
        case 'SET_AUTH':
            return {
                ...state,
                accessToken: action.payload.accessToken,
                roles: action.payload.roles,
                given_name: action.payload.given_name,
                family_name: action.payload.family_name,
                email: action.payload.email
            }
        case 'SET_THEME':
            return { ...state, theme: action.payload }
        case 'UNSET_AUTH':
            return {...initialState}
    }
}

export class AuthDispatcher {
    constructor(dispatcher) {
        this.dispatch = dispatcher
    }

    unsetAuth() {
        this.dispatch({
            type: "UNSET_AUTH"
        })
    }
    setAuth(accessToken, roles, given_name, family_name, email) {
        this.dispatch({
            type: "SET_AUTH", payload: {
                email,
                accessToken,
                roles,
                given_name,
                family_name
            }
        })
    }

    setTheme(theme) {
        this.dispatch({ type: "SET_THEME", payload: theme })
    }
}


export const AuthContext = createContext({})
export const AuthProvider = ({ children }) => {
    const [state, dispatch] = useReducer(AuthReducer, initialState)
    const [dispatcher] = useState(new AuthDispatcher(dispatch))

    return (
        <AuthContext.Provider
            value={{
                auth: state,
                dispatcher
            }}
        >
            {children}
        </AuthContext.Provider>
    )
}

export const AuthConsumer = AuthContext.Consumer
export const useAuth = () => useContext(AuthContext)