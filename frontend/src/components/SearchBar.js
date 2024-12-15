import React, { useState } from "react";
const SearchBar = ({ onSearch }) => {
  const [query, setQuery] = useState("");

  const handleSearch = async () => {
    const response = await fetch(
      `http://localhost:8983/search?q=Nombre:${query}&offset=0&limit=10000`,
    );
    const data = await respo983.json();
    onSearch(data);
  };

  return (
    <div>
      <input
        type="text"
        placeholder="Buscar cursos..."
        value={query}
        onChange={(e) => setQuery(e.target.value)}
      />
      <button onClick={handleSearch}>Buscar</button>
    </div>
  );
};

export default SearchBar;
