import { Outlet, Link } from 'react-router-dom';

function Profile() {
    return (
        <div>
            <h2>用户中心</h2>
            <nav>
                <ul style={{ display: 'flex', gap: '20px', listStyle: 'none' }}>
                    <li><Link to="/profile">主页</Link></li>
                    <li><Link to="/profile/orders">订单</Link></li>
                    <li><Link to="/profile/settings">设置</Link></li>
                </ul>
            </nav>

            <div style={{ marginTop: '20px' }}>
                <Outlet />
            </div>
        </div>
    );
}

export default Profile;