import axios from "axios";
import "./App.css";

function App() {
  const doSomething = async () => {
    const result = await axios.get("http://15.206.122.167:8081/dummy");
    console.log(result);
  };

  return (
    <>
      <button onClick={doSomething}>MAGIC</button>
    </>
  );
}

export default App;
