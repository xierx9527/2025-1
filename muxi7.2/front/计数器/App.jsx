import { useCounter } from './useCounter';
import './App.css';

function App() {
    const { count, increment, decrement, reset } = useCounter(0);

    return (
        <div className="app">
            <h1>React计数器</h1>
            <div className="counter">{count}</div>

            <div className="controls">
                <button onClick={increment}>➕ 增加</button>
                <button onClick={decrement}>➖ 减少</button>
                <button onClick={reset}>🔄 重置</button>
            </div>
        </div>
    );
}

export default App;