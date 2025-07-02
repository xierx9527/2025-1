import { createContext, useContext, useState, useEffect } from 'react';

const ThemeContext = createContext();

export function ThemeProvider({ children }) {
    const [theme, setTheme] = useState('light');

    const toggleTheme = () => {
        setTheme(prevTheme => (prevTheme === 'light' ? 'dark' : 'light'));
    };

    // 同步body背景色变化
    useEffect(() => {
        document.body.style.backgroundColor = theme === 'light' ? '#ffffff' : '#121212';
        document.body.style.color = theme === 'light' ? '#000000' : '#ffffff';
    }, [theme]);

    return (
        <ThemeContext.Provider value={{ theme, toggleTheme }}>
            {children}
        </ThemeContext.Provider>
    );
}

// 自定义Hook
export function useTheme() {
    return useContext(ThemeContext);
}