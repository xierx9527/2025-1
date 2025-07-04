import { useNavigate } from 'react-router-dom';
import { logout } from '../api/auth';

export default function LogoutButton() {
    const navigate = useNavigate();

    const handleLogout = () => {
        logout();
        navigate('/login'); // 跳转到登录页
        window.location.reload(); // 确保状态更新
    };

    return (
        <button onClick={handleLogout} className="logout-btn">
            退出登录
        </button>
    );
}