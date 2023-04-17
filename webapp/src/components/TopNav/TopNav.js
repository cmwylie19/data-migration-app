import * as React from 'react';
import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import Button from '@mui/material/Button';
import IconButton from '@mui/material/IconButton';
import Logo from "../../assets/logo.svg"
import Switch from '@mui/material/Switch';
import FormControlLabel from '@mui/material/FormControlLabel';
import FormGroup from '@mui/material/FormGroup';

export default function ButtonAppBar({given_name, authReducer}) {
  return (
    <Box sx={{ flexGrow: 1 }}>
      <AppBar position="static">
        <Toolbar>
          <IconButton
            size="large"
            edge="start"
            color="inherit"
            aria-label="menu"
            sx={{ mr: 2 }}
          >
           <img src={Logo} height="20"/>
          </IconButton>

          <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
            Data Migration
          </Typography>
          <FormGroup>
        <FormControlLabel
          control={
            <Switch
              checked={authReducer.auth.theme == "dark" ? true : false}
              onChange={()=>authReducer.dispatcher.setTheme(authReducer.auth.theme == "dark" ? "light" : "dark")}
              aria-label="login switch"
            />
          }
          label={authReducer.auth.theme == "dark" ? 'dark' : 'light'}
        />
      </FormGroup>
          <Button color="inherit" onClick={()=>authReducer.dispatcher.unsetAuth()}>Logout, {given_name}?</Button>
        </Toolbar>
      </AppBar>
    </Box>
  );
}