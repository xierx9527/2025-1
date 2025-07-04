import { Link } from 'react-router-dom';
import { getAuthToken, logout } from '../api/auth'; // 导入认证相关方法
import { useNavigate } from 'react-router-dom';

function Navigation() {
    const navigate = useNavigate();
    const isLoggedIn = !!getAuthToken(); // 检查登录状态

    const handleLogout = () => {
        logout(); // 清除token
        navigate('/login'); // 跳转到登录页
        window.location.reload(); // 刷新页面确保状态更新
    };

    return (
        <nav>
            <ul style={{ display: 'flex', gap: '20px', listStyle: 'none', alignItems: 'center' }}>
                <li><Link to="/">首页</Link></li>
                <li><Link to="/books">图书列表</Link></li>
                {isLoggedIn && (
                    <>
                        <li><Link to="/profile">用户中心</Link></li>
                        <li>
                            <button
                                onClick={handleLogout}
                                style={{
                                    background: 'none',
                                    border: 'none',
                                    color: 'inherit',
                                    cursor: 'pointer',
                                    padding: 0,
                                    font: 'inherit',
                                    textDecoration: 'underline'
                                }}
                            >
                                退出登录
                            </button>
                        </li>
                    </>
                )}
                {!isLoggedIn && <li><Link to="/login">登录</Link></li>}
            </ul>
        </nav>
    );
}

export default Navigation;