import React from "react";
import {
  AppBar,
  Box,
  List,
  ListItem,
  ListItemButton,
  Toolbar,
  Typography,
} from "@mui/material";

function App() {
  const pages = ["Feed", "Search", "Post"];

  return (
    <Box>
      <AppBar>
        <Toolbar>
          <List sx={{ display: "flex", flexDirection: "row" }}>
            {pages.map((page) => (
              <ListItem>
                <ListItemButton>
                  <Typography>{page}</Typography>
                </ListItemButton>
              </ListItem>
            ))}
          </List>
        </Toolbar>
      </AppBar>
    </Box>
  );
}

export default App;
