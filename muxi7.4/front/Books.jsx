import { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import { getAllBooks, deleteBook } from '../api/books';
import { getAuthToken } from '../api/auth';

export default function Books() {
    const [books, setBooks] = useState([]);
    const isAuthenticated = !!getAuthToken();

    useEffect(() => {
        const fetchBooks = async () => {
            const data = await getAllBooks();
            setBooks(data);
        };
        fetchBooks();
    }, []);

    const handleDelete = async (id) => {
        await deleteBook(id);
        setBooks(books.filter(book => book.id !== id));
    };

    return (
        <div>
            <h2>图书列表</h2>
            {isAuthenticated && <Link to="/books/add">添加新书</Link>}
            <ul>
                {books.map(book => (
                    <li key={book.ID || book.id}>
                        <Link to={`/books/${book.ID || book.id}`}>{book.Title || book.title}</Link>
                        {isAuthenticated && (
                            <>
                                <button onClick={() => handleDelete(book.ID || book.id)}>删除</button>
                                <Link to={`/books/edit/${book.ID || book.id}`}>编辑</Link>
                            </>
                        )}
                    </li>
                ))}
            </ul>
        </div>
    );
}