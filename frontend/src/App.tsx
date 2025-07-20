import { useEffect, useState } from "react";

const API_BASE = import.meta.env.VITE_API_URL || "";

function App() {
  const [ping, setPing] = useState("...");

  useEffect(() => {
    console.log("API_BASE:", API_BASE);
    fetch(`${API_BASE}/api/ping`)
      .then((res) => res.text())
      .then(setPing)
      .catch((err) => setPing("Error: " + err.message));
  }, []);

  return (
    <div style={{ padding: "2rem", fontFamily: "sans-serif" }}>
      <h1>AskWise Frontend</h1>
      <p>
        Ping result from backend: <strong>{ping}</strong>
      </p>
    </div>
  );
}

export default App;
