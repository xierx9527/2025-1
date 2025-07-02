import { useTheme } from './ThemeContext';

export function ThemeCard() {
    const { theme } = useTheme();

    const themeStyles = {
        light: {
            backgroundColor: '#f5f5f5',
            color: '#333',
            borderColor: '#ddd'
        },
        dark: {
            backgroundColor: '#333',
            color: '#f5f5f5',
            borderColor: '#555'
        }
    };

    return (
        <div style={{
            padding: '20px',
            margin: '20px 0',
            borderRadius: '8px',
            border: '1px solid',
            ...themeStyles[theme]
        }}>
            <h2>当前主题: {theme === 'light' ? '明亮模式' : '暗黑模式'}</h2>
            <p>这是一个{theme === 'light' ? '明亮' : '暗黑'}主题的卡片示例。</p>
        </div>
    );
}