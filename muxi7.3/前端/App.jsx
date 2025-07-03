import { Routes, Route } from 'react-router-dom';
import Navigation from './components/Navigation';
import Home from './pages/Home';
import Products from './pages/Products';
import ProductDetail from './pages/ProductDetail';
import Profile from './pages/Profile';
import Orders from './pages/Profile/Orders';
import Settings from './pages/Profile/Settings';
import Login from './pages/Login';
import NotFound from './pages/NotFound';
import AuthCheck from './components/AuthCheck';

function App() {
    return (
        <div className="app">
            <Navigation />

            <Routes>
                <Route path="/" element={<Home />} />
                <Route path="/products" element={<Products />} />
                <Route path="/products/:id" element={<ProductDetail />} />

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