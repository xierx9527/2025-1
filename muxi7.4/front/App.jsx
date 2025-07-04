import { Routes, Route } from 'react-router-dom';
import Navigation from './components/Navigation';
import Home from './pages/Home';
import Books from './pages/Books';
import BookDetail from './pages/BookDetail';
import Profile from './pages/Profile/Profile';
import Orders from './pages/Profile/Orders';
import Settings from './pages/Profile/Settings';
import Login from './pages/Login';
import NotFound from './pages/NotFound';
import AuthCheck from './components/AuthCheck';
import BookForm from './pages/BookForm';

function App() {
    return (
        <div className="app">
            <Navigation />

            <Routes>
                <Route path="/" element={<Home />} />
                <Route path="/books" element={<Books />} />
                <Route path="/books/:id" element={<BookDetail />} />
                <Route path="/books/add" element={<BookForm mode="add" />} />
                <Route path="/books/edit/:id" element={<BookForm mode="edit" />} />

                {/* 受保护的路由 */}
                <Route path="/profile" element={<AuthCheck><Profile /></AuthCheck>}>
                    <Route index element={<div>个人中心主页</div>} />
                    <Route path="orders" element={<Orders />} />
                    <Route path="settings" element={<Settings />} />
                </Route>

                <Route path="/login" element={<Login />} />
                <Route path="*" element={<NotFound />} />
            </Routes>
        </div>
    );
}

export default App;