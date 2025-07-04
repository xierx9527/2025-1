import { useEffect } from 'react';
import { useNavigate, useLocation } from 'react-router-dom';
import { getAuthToken } from '../api/auth';

export default function AuthCheck({ children }) {
    const navigate = useNavigate();
    const location = useLocation();

    useEffect(() => {
        const token = getAuthToken();
        if (!token) {
            navigate('/login', {
                state: { from: location },
                replace: true
            });
        }
    }, [navigate, location]);

    return getAuthToken() ? children : null;
}