import { useState, useEffect } from 'react'; // 添加useState导入
import { useNavigate, useParams } from 'react-router-dom';
import { getBookById, addBook, updateBook } from '../api/books';

export default function BookForm({ mode }) {
    const { id } = useParams();
    const navigate = useNavigate();
    const [formData, setFormData] = useState({
        ID: '',
        Title: '',
        Author: '',
        Stock: 0
    });

    useEffect(() => {
        if (mode === 'edit' && id) {
            const loadBook = async () => {
                try {
                    const data = await getBookById(id);
                    setFormData({
                        ID: data.id || data.ID,
                        Title: data.title || data.Title,
                        Author: data.author || data.Author,
                        Stock: data.stock || data.Stock
                    });
                } catch (err) {
                    console.error('加载图书失败:', err);
                    navigate('/books');
                }
            };
            loadBook();
        }
    }, [id, mode, navigate]);

    const handleSubmit = async (e) => {
        e.preventDefault();
        const payload = {
            id: formData.ID,
            title: formData.Title,
            author: formData.Author,
            stock: Number(formData.Stock)
        };

        try {
            if (mode === 'edit') {
                await updateBook(id, payload);
            } else {
                await addBook(payload);
            }
            navigate('/books');
        } catch (err) {
            alert(`保存失败: ${err.response?.data?.error || err.message}`);
        }
    };

    return (
        <div className="book-form-container">
            <h2>{mode === 'edit' ? '编辑图书' : '添加新书'}</h2>
            <form onSubmit={handleSubmit}>
                <div className="form-group">
                    <label>图书ID：</label>
                    <input
                        type="text"
                        value={formData.ID}
                        onChange={(e) => setFormData({ ...formData, ID: e.target.value })}
                        required
                        disabled={mode === 'edit'} // 编辑时禁用ID修改
                    />
                </div>

                <div className="form-group">
                    <label>书名：</label>
                    <input
                        type="text"
                        value={formData.Title}
                        onChange={(e) => setFormData({ ...formData, Title: e.target.value })}
                        required
                    />
                </div>

                <div className="form-group">
                    <label>作者：</label>
                    <input
                        type="text"
                        value={formData.Author}
                        onChange={(e) => setFormData({ ...formData, Author: e.target.value })}
                    />
                </div>

                <div className="form-group">
                    <label>库存：</label>
                    <input
                        type="number"
                        min="0"
                        value={formData.Stock}
                        onChange={(e) => setFormData({ ...formData, Stock: e.target.value })}
                        required
                    />
                </div>

                <button type="submit" className="submit-btn">
                    {mode === 'edit' ? '更新图书' : '添加图书'}
                </button>
            </form>
        </div>
    );
}