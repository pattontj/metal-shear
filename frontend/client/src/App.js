import logo from './logo.svg';
import './App.css';

import {BrowserRouter,
    Routes,    
    Route
} from 'react-router-dom';

import Home from './Home.js'
import Statsuki from './Statsuki';

function App() {
  return (
    <div className="App">
        <BrowserRouter>
            <Routes>

                <Route path="/" element={<Home/ >} />
                <Route path="/statsuki" element={<Statsuki/ >} />

            </Routes>
        </BrowserRouter>

    </div>
  );
}

export default App;
