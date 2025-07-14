import { useEffect, useState } from "react";

function App() {
  const [ping, setPing] = useState("...");

  useEffect(() => {
    fetch("/api/ping")
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
