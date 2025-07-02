import React, { useState, useReducer, useContext, useEffect, useMemo, useRef } from 'react';

// 1. 创建context
const TodoContext = React.createContext();

// 2. 定义action类型
const ACTIONS = {
  ADD_TODO: 'add-todo',
  TOGGLE_TODO: 'toggle-todo',
  DELETE_TODO: 'delete-todo'
};

// 3. reducer函数
function todoReducer(todos, action) {
  switch (action.type) {
    case ACTIONS.ADD_TODO:
      return [...todos, newTodo(action.payload.name)];
    case ACTIONS.TOGGLE_TODO:
      return todos.map(todo => {
        if (todo.id === action.payload.id) {
          return { ...todo, complete: !todo.complete };
        }
        return todo;
      });
    case ACTIONS.DELETE_TODO:
      return todos.filter(todo => todo.id !== action.payload.id);
    default:
      return todos;
  }
}

function newTodo(name) {
  return { id: Date.now(), name: name, complete: false };
}

// 4. Todo组件
function Todo({ todo }) {
  const dispatch = useContext(TodoContext);

  return (
    <div style={{ display: 'flex', alignItems: 'center', margin: '5px 0' }}>
      <input
        type="checkbox"
        checked={todo.complete}
        onChange={() => dispatch({ type: ACTIONS.TOGGLE_TODO, payload: { id: todo.id } })}
      />
      <span
        style={{
          margin: '0 10px',
          textDecoration: todo.complete ? 'line-through' : 'none',
          color: todo.complete ? '#888' : '#000'
        }}
      >
        {todo.name}
      </span>
      <button onClick={() => dispatch({ type: ACTIONS.DELETE_TODO, payload: { id: todo.id } })}>
        删除
      </button>
    </div>
  );
}

// 5. 主组件
function App() {
  const [todos, dispatch] = useReducer(todoReducer, []);
  const [name, setName] = useState('');
  const inputRef = useRef(null);

  // 使用useEffect在组件挂载时打印消息
  useEffect(() => {
    console.log('Todo List已加载');
  }, []);

  // 处理添加todo
  function handleSubmit(e) {
    e.preventDefault();
    if (!name.trim()) return;

    dispatch({ type: ACTIONS.ADD_TODO, payload: { name: name } });
    setName('');
    inputRef.current.focus();
  }

  // 使用useMemo优化列表渲染
  const todoList = useMemo(() => {
    return todos.map(todo => (
      <Todo key={todo.id} todo={todo} />
    ));
  }, [todos]);

  return (
    <TodoContext.Provider value={dispatch}>
      <div style={{ maxWidth: '400px', margin: '0 auto', padding: '20px' }}>
        <h1>Todo List</h1>

        <form onSubmit={handleSubmit} style={{ marginBottom: '20px' }}>
          {/* 受控组件示例 */}
          <input
            type="text"
            value={name}
            onChange={(e) => setName(e.target.value)}
            placeholder="输入任务内容"
            style={{ marginRight: '10px', padding: '5px' }}
          />
          <button type="submit">添加</button>
        </form>

        {/* 非受控组件示例 */}
        <form onSubmit={(e) => {
          e.preventDefault();
          const value = inputRef.current.value;
          if (!value.trim()) return;

          dispatch({ type: ACTIONS.ADD_TODO, payload: { name: value } });
          inputRef.current.value = '';
        }} style={{ marginBottom: '20px' }}>
          <input
            type="text"
            ref={inputRef}
            placeholder="输入任务内容(非受控)"
            style={{ marginRight: '10px', padding: '5px' }}
          />
          <button type="submit">添加(非受控)</button>
        </form>

        <div>
          {todoList}
        </div>
      </div>
    </TodoContext.Provider>
  );
}

export default App;