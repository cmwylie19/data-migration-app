import { amber, red, blue, grey, green } from "@mui/material/colors";

export const getDesignTokens = (mode) => ({
    palette: {
      mode: mode,
      primary: {
        ...red,
        ...(mode === "light" && {
          main: red[800],
        }),
      },
      secondary: {
        ...red,
        ...(mode === "light" && {
          main: grey[900],
        }),
      },
      ...(mode === "dark" && {}),
      text: {
        ...(mode === "light"
          ? {
            primary: grey[900],
            secondary: grey[800],
          }
          : {
            primary: "#fff",
            secondary: grey[500],
          }),
      },
    },
  });