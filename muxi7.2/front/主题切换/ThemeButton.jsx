import { useTheme } from './ThemeContext';

export function ThemeButton() {
    const { toggleTheme, theme } = useTheme();

    return (
        <button
            onClick={toggleTheme}
            style={{
                padding: '10px 20px',
                fontSize: '16px',
                backgroundColor: theme === 'light' ? '#333' : '#f5f5f5',
                color: theme === 'light' ? '#f5f5f5' : '#333',
                border: 'none',
                borderRadius: '4px',
                cursor: 'pointer'
            }}
        >
            🎨 切换主题
        </button>
    );
}