import * as React from 'react';

import Button from '@mui/material/Button';
import CssBaseline from '@mui/material/CssBaseline';
import TextField from '@mui/material/TextField';
import Grid from '@mui/material/Grid';
import Box from '@mui/material/Box';
import Typography from '@mui/material/Typography';
import Container from '@mui/material/Container';
import Logo from "../../assets/logo.svg"
import {handleLogin } from "../../API"
export default function SignIn({authReducer}) {


  const [user, setUser] = React.useState("")
  const [pw, setPW] = React.useState("")
  const [errorMessage, setErrorMessage] = React.useState(false)

  const styles = {
    display: "flex",
    justifyContext: "center",
    alignItems: "center",
    height: "100%vh",
    width: "100%",
    position: "absolute",
    top: 0,
    left: 0,
    right: 0,
    bottom: 0
  }
  return (
    <Container component="main" maxWidth="xs" style={styles}>
      <CssBaseline />

      <Box
        sx={{
          marginTop: 12,
          display: 'flex',
          flexDirection: 'column',
          alignItems: 'center',
        }}
      >

        <img src={Logo} height="70" style={{ margin: "20px" }} />

        <Typography component="h1" variant="h5">
          Data Migration Sign in
        </Typography>
        <Box component="form" noValidate sx={{ mt: 1 }}>
          <TextField
            margin="normal"
            required
            fullWidth
            id="email"
            label="Email Address"
            name="email"
            autoComplete="email"
            onChange={e => setUser(e.target.value)}
            autoFocus
          />
          <TextField
            margin="normal"
            required
            fullWidth
            name="password"
            label="Password"
            type="password"
            id="password"
            onChange={e => setPW(e.target.value)}
            // autoComplete="current-password"
          />

          <Button
       
            fullWidth
            variant="contained"
            sx={{ mt: 3, mb: 2 }}
            onClick={()=>handleLogin(user, pw,authReducer.dispatcher,setErrorMessage)}
          >
            Sign In
          </Button>
         {errorMessage && <Typography align="center" color="error">Incorrect user or password</Typography>}
          <Grid container>

          </Grid>
        </Box>
      </Box>
    </Container>

  );
}