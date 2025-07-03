import { useNavigate } from 'react-router-dom';

function NotFound() {
    const navigate = useNavigate();

    return (
        <div>
            <h2>404 - 页面未找到</h2>
            <p>您访问的页面不存在</p>
            <button onClick={() => navigate('/')}>返回首页</button>
        </div>
    );
}

export default NotFound;