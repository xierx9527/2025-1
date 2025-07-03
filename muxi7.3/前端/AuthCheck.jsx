import { useEffect } from 'react';
import { useNavigate, useLocation } from 'react-router-dom';

function AuthCheck({ children }) {
    const navigate = useNavigate();
    const location = useLocation();
    const isAuthenticated = false; // 实际项目中应从状态管理获取

    useEffect(() => {
        if (!isAuthenticated) {
            navigate('/login', {
                state: { from: location },
                replace: true
            });
        }
    }, [navigate, location, isAuthenticated]);

    return isAuthenticated ? children : null;
}

export default AuthCheck;