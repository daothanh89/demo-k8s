import React, { useState, useEffect } from "react";

function App() {
  const [quotes, setQuotes] = useState([]);

  useEffect(() => {
    fetch(process.env.REACT_APP_API_QUOTE)
      .then((res) => res.json())
      .then((q) => setQuotes(q))
      .catch(console.log);
  }, []);

  return (
    <div className="bg-gray-500">
      <div className="flex items-center justify-center pt-8 text-3xl">
        Quotes
      </div>
      <div className="py-10 mx-64">
        {quotes.map((quote, i) => (
          <div className="flex pt-6" key={"index" + i}>
            <blockquote className="flex flex-wrap flex-col bg-white text-indigo-700 border-l-8 italic border-gray-400 px-4 py-3">
              {quote.text}
              <span className="flex justify-end text-sm text-indigo-400 font-semibold pt-2 underline ">
                {quote.author}
              </span>
            </blockquote>
          </div>
        ))}
      </div>
    </div>
  );
}

export default App;
