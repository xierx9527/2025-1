function Settings() {
    return (
        <div>
            <h3>账户设置</h3>
            <form>
                <div>
                    <label>用户名:</label>
                    <input type="text" />
                </div>
                <div>
                    <label>邮箱:</label>
                    <input type="email" />
                </div>
                <button type="submit">保存</button>
            </form>
        </div>
    );
}

export default Settings;