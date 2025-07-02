import { ThemeProvider } from './ThemeContext';
import { ThemeCard } from './ThemeCard';
import { ThemeButton } from './ThemeButton';

export default function ThemeApp() {
    return (
        <ThemeProvider>
            <div style={{ padding: '20px', maxWidth: '800px', margin: '0 auto' }}>
                <h1>主题切换应用</h1>
                <ThemeButton />
                <ThemeCard />
                <p>这个应用展示了如何使用React Context实现全局主题切换功能。</p>
            </div>
        </ThemeProvider>
    );
}