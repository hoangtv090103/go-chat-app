import React, { useState, useEffect } from "react";
import axios from "axios";

function RoomList({ onSelectRoom }) {
  const [rooms, setRooms] = useState([]);
  const [newRoomName, setNewRoomName] = useState("");

  useEffect(() => {
    // Fetch available rooms
    async function fetchRooms() {
      try {
        const response = await axios.get("http://localhost:3000/rooms");

        // Check the structure of the data
        console.log("Rooms data received:", response.data);

        setRooms(response.data);
      } catch (error) {
        console.error("Error fetching rooms", error);
        alert("Failed to fetch rooms");
      }
    }
    fetchRooms();
  }, []);

  // Handle the creation of a new room
  const createRoom = async (e) => {
    e.preventDefault();
    if (!newRoomName.trim()) {
      alert("Room name cannot be empty");
      return;
    }

    try {
      const response = await axios.post("http://localhost:3000/rooms", {
        name: newRoomName,
      });
      const createdRoom = response.data;

      // Add the newly created room to the room list
      setRooms([...rooms, createdRoom]);

      // Clear the input field
      setNewRoomName("");
    } catch (error) {
      console.error("Error creating room", error);
      alert("Failed to create room");
    }
  };

  return (
    <div className="room-list">
      <h3>Available Rooms</h3>
      <ul>
        {rooms.length === 0 ? (
          <li>No rooms available</li>
        ) : (
          rooms.map((room) => (
            <li key={room.id} onClick={() => onSelectRoom(room.id)}>
              {room.name ? room.name : "Unnamed Room"}
            </li>
          ))
        )}
      </ul>

      {/* Form to create a new room */}
      <div className="create-room">
        <h4>Create a new Room</h4>
        <form onSubmit={createRoom}>
          <input
            type="text"
            placeholder="Room Name"
            value={newRoomName}
            onChange={(e) => setNewRoomName(e.target.value)}
          />
          <button type="submit">Create Room</button>
        </form>
      </div>
    </div>
  );
}

export default RoomList;
