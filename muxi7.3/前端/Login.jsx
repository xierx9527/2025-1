import { useLocation, useNavigate } from 'react-router-dom';

function Login() {
    const location = useLocation();
    const navigate = useNavigate();
    const from = location.state?.from?.pathname || '/';

    const handleLogin = () => {
        // 实际项目中这里会有登录逻辑
        // 登录成功后跳转回原页面
        navigate(from, { replace: true });
    };

    return (
        <div>
            <h2>登录页面</h2>
            <p>您需要登录才能访问 {from}</p>
            <button onClick={handleLogin}>模拟登录</button>
        </div>
    );
}

export default Login;