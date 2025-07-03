import { Link } from 'react-router-dom';

function Navigation() {
    return (
        <nav>
            <ul style={{ display: 'flex', gap: '20px', listStyle: 'none' }}>
                <li><Link to="/">首页</Link></li>
                <li><Link to="/products">商品列表</Link></li>
                <li><Link to="/profile">用户中心</Link></li>
            </ul>
        </nav>
    );
}

export default Navigation;