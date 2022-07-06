import React from "react";
import {
  AppBar,
  Box,
  Divider,
  List,
  ListItem,
  ListItemButton,
  Toolbar,
  Typography,
} from "@mui/material";
import { Link, Route, Routes } from "react-router-dom";

function App() {
  const pages = ["feed", "search", "post"];

  return (
    <Box>
      <AppBar position="static">
        <Toolbar>
          <List disablePadding sx={{ display: "flex", flexDirection: "row" }}>
            {pages.map((page) => (
              <ListItem sx={{ paddingLeft: "0px", paddingRight: "0px" }}>
                <ListItemButton component={Link} to={"/" + page}>
                  <Typography>{page}</Typography>
                </ListItemButton>
                <Divider orientation="vertical" />
              </ListItem>
            ))}
          </List>
        </Toolbar>
      </AppBar>
      <Routes>
        <Route path="/" element={<h1>login</h1>} />
        <Route path="/feed" element={<img src="/logo" />} />
        <Route path="/search" element={<h1>search</h1>} />
        <Route path="/post" element={<h1>post</h1>} />
      </Routes>
    </Box>
  );
}

export default App;
