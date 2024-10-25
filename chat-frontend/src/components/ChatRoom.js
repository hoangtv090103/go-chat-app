import React, { useState, useEffect, useRef } from "react";

function ChatRoom({ token, roomID }) {
  const [messages, setMessages] = useState([]);
  const [input, setInput] = useState("");
  const socketRef = useRef(null);

  useEffect(() => {
    if (socketRef.current) {
      return; // Prevent creating a new socket if one already exists
    }

    const socket = new WebSocket(
      `ws://localhost:3001/ws?token=${token}&room-id=${roomID}`
    );

    socketRef.current = socket;

    socket.onopen = () => {
      console.log("WebSocket connection opened");
    };

    // Handle incoming messages (both message history and real-time messages)
    socket.onmessage = (event) => {
      const message = JSON.parse(event.data);

      setMessages((prevMessages) => [...prevMessages, message]);
    };

    socket.onerror = (error) => {
      console.error("WebSocket error:", error);
    };

    socket.onclose = (event) => {
      console.log("WebSocket connection closed:", event.code, event.reason);
    };

    // Cleanup WebSocket connection when component unmounts
    return () => {
      if (socketRef.current) {
        socketRef.current.close();
        socketRef.current = null; // Clean up the ref
      }
    };
  }, [token, roomID]);

  const sendMessage = () => {
    if (socketRef.current && socketRef.current.readyState === WebSocket.OPEN) {
      const message = { event: "message", message: input, roomID };
      socketRef.current.send(JSON.stringify(message)); // Send the message to the WebSocket server
      setInput(""); // Clear input after sending message
      // No need to update messages here! Let the WebSocket broadcast handle that
    }
  };

  return (
    <div className="chat-room">
      <h3>Room ID: {roomID}</h3>
      <div className="messages">
        {messages.map((msg) => (
          <div key={msg.id}>
            {/* Assuming each message has a unique id */}
            <strong>{msg.sender_id}:</strong> {msg.message}
          </div>
        ))}
      </div>
      <input
        type="text"
        value={input}
        onChange={(e) => setInput(e.target.value)}
        placeholder="Type a message"
      />
      <button onClick={sendMessage}>Send</button>
    </div>
  );
}

export default ChatRoom;
