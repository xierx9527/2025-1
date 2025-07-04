import { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { getBookById } from '../api/books';

export default function BookDetail() {
    const { id } = useParams();
    const [book, setBook] = useState(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState('');

    useEffect(() => {
        const fetchBook = async () => {
            try {
                const data = await getBookById(id);
                setBook(data);
            } catch (err) {
                setError('获取图书信息失败');
            } finally {
                setLoading(false);
            }
        };
        fetchBook();
    }, [id]);

    if (loading) return <div>加载中...</div>;
    if (error) return <div>{error}</div>;
    if (!book) return <div>图书不存在</div>;

    return (
        <div>
            <h2>图书详情</h2>
            <h3>{book.Title || book.title}</h3>
            <p>作者: {book.Author || book.author}</p>
            <p>库存: {book.Stock || book.stock}</p>
            <p>ID: {book.ID || book.id}</p>
        </div>
    );
}