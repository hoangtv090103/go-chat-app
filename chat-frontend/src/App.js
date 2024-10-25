import React, { useEffect, useState } from "react";
import "./App.css";
import Login from "./components/Login";
import RoomList from "./components/RoomList";
import ChatRoom from "./components/ChatRoom";

function App() {
  const [token, setToken] = useState(null);
  const [roomID, setRoomID] = useState(null);
  useEffect(() => {
    setToken(localStorage.getItem("token"));
  }, []);
  if (!token) {
    return <Login onLogin={(token) => setToken(token)} />;
  }

  if (!roomID) {
    return <RoomList onSelectRoom={(roomID) => setRoomID(roomID)} />;
  }

  return <ChatRoom roomID={roomID} token={token} />;
}

export default App;
